package pgsql

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	db              *gorm.DB
	ID              string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name            string
	Sender          string
	ThreadId        string
	UnsubscribeLink string
	Unsubscribed    bool
}

func GetConnection() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=pgsql user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Service{})

	return db
}

func (*Service) Create(service Service) {
	db := GetConnection()

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("got error while closing db connection")
	}

	defer sqlDB.Close()
	db.Create(service)
}

func (*Service) SearchByNameAndSender(name string, sender string) []Service {
	db := GetConnection()
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("got error while closing db connection")
	}
	defer sqlDB.Close()

	var services []Service

	db.Where(&Service{Name: name, Sender: sender}).Find(&services)

	return services
}

func (*Service) Get() []Service {
	db := GetConnection()
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("got error while closing db connection")
	}
	defer sqlDB.Close()

	var services []Service

	db.Model(&Service{}).Where("unsubscribed = ?", false).Find(&services)

	return services
}

func (*Service) Unsubscribe(id string) {
	db := GetConnection()
	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("got error while closing db connection")
	}
	defer sqlDB.Close()

	db.Model(&Service{}).Where("id = ?", id).Update("unsubscribed", true)
}
