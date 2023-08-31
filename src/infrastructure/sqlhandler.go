package infrastructure

import (
	"fmt"
	"os"

	"github.com/Egas88/AvitoTest/src/interfaces/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.db = db
	return SqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}

func (handler *SqlHandler) Find(obj interface{}, conds ...interface{}) (tx *gorm.DB) {
	return handler.db.Find(obj, conds...)
}

func (handler *SqlHandler) Preload(query string) (tx *gorm.DB) {
	return handler.db.Preload(query)
}

func (handler *SqlHandler) Where(cond interface{}, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Where(cond, args...)
}

func (handler *SqlHandler) Association(column string) (tx *gorm.Association) {
	return handler.db.Association(column)
}

func (handler *SqlHandler) Model(obj interface{}) (tx *gorm.DB) {
	return handler.db.Model(obj)
}

func (handler *SqlHandler) First(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return handler.db.First(dest, conds...)
}
