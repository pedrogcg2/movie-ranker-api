package main

import (
	"fmt"
	"log"
	"movie-api/api/routes"
	"movie-api/domain/entities"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := connectDatabase()

	sqlDb, _ := db.DB()

	routes.InitRoutes(*db)
	s := &http.Server{
		Addr: ":8000",

		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}

	defer sqlDb.Close()
	log.Fatal(s.ListenAndServe())
}

// TODO: add config to .env
// Add proper error handling
// maybe create a package for this
func connectDatabase() *gorm.DB {
	host := "postgres"
	user := "postgres"
	password := "pedro"
	dbname := "MovieRanker"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Fortaleza",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(dsn)
	}

	db.AutoMigrate(&entities.User{}, &entities.Movie{}, &entities.MovieAvaliation{})

	return db
}
