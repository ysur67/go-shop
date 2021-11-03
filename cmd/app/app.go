package main

import (
	server "shop/server"
)

func main() {
	srv := server.NewApp()
	if err := srv.Run("4444"); err != nil {
		panic(err)
	}
}
