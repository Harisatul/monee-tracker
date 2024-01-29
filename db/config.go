package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"technopartner/config"
	"technopartner/db/seed"
	"technopartner/internal/entity"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}
func InitDB() {

	db := config.GetConfig()

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		db.DB_USERNAME,
		db.DB_PASSWORD,
		db.DB_NAME,
		db.DB_HOST,
		db.DB_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	//&entity.Role{}
	DB.Migrator().DropTable(&entity.User{})
	DB.AutoMigrate(entity.User{}, entity.Role{}, entity.Category{})
	DB.Migrator().HasConstraint(&entity.User{}, "role")
	//DB.Migrator().HasConstraint(&re.Role{}, "Users")
	//DB.Migrator().HasConstraint(&pe.Product{}, "TransactionDetail")
}
