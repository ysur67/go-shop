package main

import (
	"fmt"
	server "shop/server"
)

func main() {
	srv := server.NewApp()
	if err := srv.Run("4444"); err != nil {
		fmt.Println("asdfsf")
		panic(err)
	}
}
