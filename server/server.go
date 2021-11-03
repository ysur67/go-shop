package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop/internal/category"
	categoryHttp "shop/internal/category/http"
	categoryRepo "shop/internal/category/repository/postgres"
	categoryUseCase "shop/internal/category/usecase"
	"time"
)

type App struct {
	server          *http.Server
	categoryUseCase category.UseCase
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
	return &App{
		categoryUseCase: categoryUseCase.NewUseCase(catRepo),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")
	categoryHttp.RegisterHttpEndpoints(api, app.categoryUseCase)

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
	viper.SetConfigName(".env")
	return viper.ReadInConfig()
}

func initDB() (*gorm.DB, error) {
	host := viper.GetString("db-host")
	port := viper.GetString("db-port")
	user := viper.GetString("db-user")
	password := viper.GetString("db-user-password")
	dbname := viper.GetString("db-name")
	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s "+
			"port=%s user=%s passsword=%s",
		host, port, dbname, host, user, password,
	)
	fmt.Print(connectionString)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
