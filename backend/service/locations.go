package service

import (
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Location struct {
	gorm.Model
	Location_id   int    `json:"id"`
	Location_name string `json:"location_name"`
}

func getLocations(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Location{})

	// Create
	db.Create(&Location{Location_id: 1, Location_name: "Smederevo"})

	// Read
	var location Location
	db.First(&location, 1) // find product with id 1

	json.NewEncoder(w).Encode(location)


}
