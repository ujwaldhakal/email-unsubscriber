package pgsql

import (
	_ "fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/driver/postgres"
"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	ID uuid.UUID  `gorm:"type:char(36);primary_key"`
	Name  string
	UnsubscribeLink string
	Unsubscribed bool

}
func GetConnection() {
// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=pgsql user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
	panic(err)
}
	db.AutoMigrate(&Service{})
}