package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=ep-withered-math-a8repc9c-pooler.eastus2.azure.neon.tech " +
		"user=neondb_owner " +
		"password=npg_yQD3UXTeZRP1 " +
		"dbname=neondb " +
		"port=5432 " +
		"sslmode=require " +
		"timezone=America/Bogota"

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, 
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error al obtener instancia de base de datos: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Conexión exitosa a NeonDB PostgreSQL")
}
