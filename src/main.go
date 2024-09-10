package main

import (
	"github.com/aridome222/Backend-Festival-Booth/api/router"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/migration"
)

func main() {
	db := infrastructure.ConnectDB()
	// migrationしたい場合はコメントアウトを外す
	migration.Migrate()
	router.SetupRouter(db)
}
