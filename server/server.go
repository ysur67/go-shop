package server

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/account"
	accountHttp "shop/internal/account/http"
	accountRepo "shop/internal/account/repository/postgres"
	accountUseCase "shop/internal/account/usecase"
	"shop/internal/category"
	categoryHttp "shop/internal/category/http"
	categoryRepo "shop/internal/category/repository/postgres"
	categoryUseCase "shop/internal/category/usecase"
	"shop/internal/product"
	productHttp "shop/internal/product/http"
	productRepo "shop/internal/product/repository/postgres"
	productUseCase "shop/internal/product/usecase"
	"shop/internal/utils"
	"time"
)

type App struct {
	server          *http.Server
	categoryUseCase category.UseCase
	productUseCase  product.UseCase
	accountUseCase  account.UseCase
}

func NewApp() *App {
	if err := readConfig(); err != nil {
		panic(err)
	}
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	catRepo := categoryRepo.NewRepository(db)
	if err := catRepo.AutoMigrate(); err != nil {
		panic(err)
	}
	prodRepo := productRepo.NewRepository(db)
	if err := prodRepo.AutoMigrate(); err != nil {
		panic(err)
	}
	accRepo := accountRepo.NewRepository(db)
	if err := accRepo.AutoMigrate(); err != nil {
		panic(err)
	}
	shouldUploadFixtures := parseCommandLine()
	if shouldUploadFixtures {
		sqlDb, err := db.DB()
		if err != nil {
			panic(err)
		}
		if err := uploadFixtures(sqlDb); err != nil {
			panic(err)
		}
	}
	fmt.Println(viper.GetDuration("token_ttl"))
	return &App{
		categoryUseCase: categoryUseCase.NewUseCase(catRepo),
		productUseCase:  productUseCase.NewUseCase(prodRepo, catRepo),
		accountUseCase: accountUseCase.NewAccountUseCase(
			accRepo,
			viper.GetString("hash_salt"),
			[]byte(viper.GetString("signing_key")),
			viper.GetDuration("token_ttl"),
		),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	router.LoadHTMLGlob("templates/*/*.html")
	router.Static("/static", "./static")
	accountHttp.RegisterApiEndpoints(router, app.accountUseCase)
	accountMiddleware := accountHttp.NewAuthMiddleware(app.accountUseCase)
	api := router.Group("/api", accountMiddleware)
	categoryHttp.RegisterApiEndpoints(api, app.categoryUseCase)
	httpRoute := router.Group("")
	categoryHttp.RegisterHttpEndpoints(httpRoute, app.categoryUseCase)
	productHttp.RegisterHttpEndpoints(httpRoute, app.productUseCase)
	accountRoute := router.Group("/account", accountMiddleware)
	accountHttp.RegisterHttpEndpoints(accountRoute, app.accountUseCase)

	app.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return app.server.Shutdown(ctx)
}

func readConfig() error {
	viper.SetConfigType("env")
	viper.SetConfigFile("./config/.env")
	return viper.ReadInConfig()
}

func initDB() (*gorm.DB, error) {
	host := viper.GetString("db_host")
	port := viper.GetInt("db_port")
	user := viper.GetString("db_user")
	password := viper.GetString("db_password")
	dbname := viper.GetString("db_name")
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s "+
			"dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}

func parseCommandLine() bool {
	shouldUpload := flag.Bool("upload", false, "Should the fixtures be loaded")
	flag.Parse()
	return *shouldUpload
}

func uploadFixtures(db *sql.DB) error {
	fixturesUploadCommand := utils.NewLoadFixturesCommand(
		db,
		"upload-fixtures",
	)
	return fixturesUploadCommand.Execute()
}
