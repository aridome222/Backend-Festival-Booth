package migration

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure"
)

func Migrate() {
	db := infrastructure.ConnectDB()

	if err := db.AutoMigrate(
		&domain.Product{},
		&domain.Profile{},
		&domain.Comment{},
	); err != nil {
		panic("failed to migrate")
	}
}
