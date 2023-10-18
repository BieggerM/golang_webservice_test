package storage

import (
	"log"
	"os"
	"strconv"
	"example.com/httprepeater/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	DB *gorm.DB
}

var db DBConnection
// SetupDBConnection contains connection details
func SetupDBConnection(stage string) {
	db.DB = NewDBConnection()
	if stage == "development" {
		MigrateDB(db.DB)
	}
}

func NewDBConnection() *gorm.DB {
	// connect to postgresql database
	log.Default().Println("Connecting to Postgres...")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("Successfully connected to Postgres!")
	return db
}

func MigrateDB (db *gorm.DB) {
	// migrate database
	log.Default().Println("Migrating database...")
	err := db.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatalln("Failed to migrate database!")
	}
}

// NewPerson create new person in database
func NewPerson(person models.Person) {
	log.Default().Println("Creating new person...")
	db.DB.Create(&person)
	log.Default().Println("Successfully created new person!")
}

func GetPersonByID(id uint64) models.Person {
	log.Default().Println("Getting person by ID...")
	var person models.Person
	db.DB.First(&person, id)
	log.Default().Println("Successfully got person by ID!")
	return person
}

func GetPersonByIDstr(stringID string) models.Person {
	log.Default().Println("Getting person by ID...")
	id, err := strconv.ParseUint(stringID, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var person models.Person
	db.DB.First(&person, id)
	log.Default().Println("Successfully got person by ID!")
	return person
}

func UpdatePerson(id uint64, person models.Person) {
	log.Default().Println("Updating person...")
	// update entry with id in database
	db.DB.Model(&models.Person{}).Where("id = ?", id).Updates(person)
	log.Default().Println("Successfully updated person!")
}

func DeletePersonById(id uint64) {
	log.Default().Println("Deleting person...")
	// delete entry with id in database
	db.DB.Delete(&models.Person{}, id)
	log.Default().Println("Successfully deleted person!")
}