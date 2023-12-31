package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"playbox/model"
)

func MakeSupaBaseConnectionDatabase(data *Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s",
		data.SupabaseUser, data.SupabasePassword, data.SupabaseHost, data.SupabasePort, data.SupabaseDbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.AquaFarmer{},
		&model.User{},
		&model.Tambak{},
		&model.Kolam{},
		&model.Siklus{},
		&model.SiklusHarian{},
		&model.Product{},
		&model.Paket{},
	); err != nil {
		return nil, err
	}
	return db, nil
}
