package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"technopartner/config"
	"technopartner/db/seed"
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
	DB.Migrator().DropTable()
	DB.AutoMigrate()
	//DB.Migrator().HasConstraint(&ue.User{}, "UserDetail")
	//DB.Migrator().HasConstraint(&re.Role{}, "Users")
	//DB.Migrator().HasConstraint(&pe.Product{}, "TransactionDetail")
}
