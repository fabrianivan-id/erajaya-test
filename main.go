package main

import (
	"github.com/fabrianivan21/erajaya-test/api/router"
	"github.com/fabrianivan21/erajaya-test/config"
)

func main() {
	db := config.InitDB()
	e := router.New(db)

	e.Logger.Fatal(e.Start(":8000"))
}
