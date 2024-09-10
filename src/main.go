package main

import (
	"github.com/aridome222/Backend-Festival-Booth/api/router"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure"
)

func main() {
	db := infrastructure.ConnectDB()

	router.SetupRouter(db)
}
