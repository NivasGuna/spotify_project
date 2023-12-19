package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func ConnectDb(){
	var err error
	
    data:=os.Getenv("DB_URL")
	DB, err=gorm.Open(postgres.Open(data), &gorm.Config{})
	

	if err !=nil{
		log.Fatal("Error loading during connection D")
	}
}