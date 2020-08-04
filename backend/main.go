package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gggruja/PAC_SRB_2020/backend/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var db *gorm.DB

func healthChecking(rw http.ResponseWriter, r *http.Request) {
	log.Println("Healthy App - Running!")
	log.Println("Backend App is up and running - Goran Grujic")

	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Backend App - Goran Grujic!"))
}

func main() {

	// log instance
	l := hclog.Default()

	// load the configuration
	cnf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	l.Info("Loaded config: ", hclog.Fmt("%+v", cnf))

	fmt.Println("Go with MariaDB")

	db, err = gorm.Open("mysql", cnf.BindDatabase)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	DbInit()

	sm := mux.NewRouter()

	sm.HandleFunc("/", healthChecking)
	sm.Handle("/metrics", promhttp.Handler())

	//LOCATIONS API's
	sm.HandleFunc("/locations", createLocations).Methods("POST")
	sm.HandleFunc("/locations", getLocations).Methods("GET")
	sm.HandleFunc("/locations/{locationId}", getLocation).Methods("GET")
	sm.HandleFunc("/locations/{locationId}", updateLocation).Methods("PUT")
	sm.HandleFunc("/locations/{locationId}", deleteLocation).Methods("DELETE")
	sm.HandleFunc("/events", getAllEvents).Methods("GET")
	sm.HandleFunc("/events/{eventId}", getOneEvent).Methods("GET")

	// create Server
	s := http.Server{
		Addr:         cnf.BindAddress,
		Handler:      sm,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 120,
	}

	go func() {
		l.Info("Starting server on " + s.Addr)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received.
	sig := <-c
	l.Info("Received signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	l.Info("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

func DbInit() {

	// Drop
	db.DropTableIfExists(&Location{}, &Event{}, &Organization{}, &Person{}, &Room{},
		&Language{}, &Talk{}, &Topic{}, &Child{}, &Parent{})

	// Create new one
	db.AutoMigrate(&Location{}, &Event{}, &Organization{}, &Person{}, &Room{},
		&Language{}, &Talk{}, &Topic{}, &Child{}, &Parent{})


	// Create records
	db.Create(&Location{LocationName: "Beograd"})
	db.Create(&Location{LocationName: "Smederevo"})

	db.Create(&Event{EventName: "Heapcon Belgrade", StartDate: time.Now(), EndDate: time.Now(), LocationId: 1})

	db.Create(&Organization{OrganizationName: "PRODYNA"})

	db.Create(&Person{PersonName: "Goran Grujic", OrganizationId: 1, TalkId: 1})

	db.Create(&Room{RoomName: "Hawaii", OrganizationId: 1})

	db.Create(&Language{LanguageName: "Serbian"})
	db.Create(&Language{LanguageName: "English"})
	db.Create(&Language{LanguageName: "German"})

	db.Create(&Talk{TitleName: "CKAD - Kubernetes Development", StartDate: time.Now(), EndDate: time.Now(), LanguageId: 1, Level: "Beginner"})
	db.Create(&Talk{TitleName: "Weed - Rolling papers", StartDate: time.Now(), EndDate: time.Now(), LanguageId: 1, Level: "Architect"})

	var childs = []Child {
		Child {
			TopicName: "Kubernetes Child 1",
		},
		Child {
			TopicName: "Kubernetes Child 2",
		},
	}

	var parents = []Parent {
		Parent {
			TopicName: "Kubernetes Parent 1",
		},
		Parent {
			TopicName: "Kubernetes Parent 2",
		},
	}

	db.Create(&Topic{TopicName: "Kubernetes", TalkId: 1, Childs:childs, Parents:parents})
	db.Create(&Topic{TopicName: "Exam", TalkId: 1})
	db.Create(&Topic{TopicName: "Rolling Pappers", TalkId: 2})
	db.Create(&Topic{TopicName: "Weed", TalkId: 2})




}

type Location struct {
	gorm.Model
	LocationName string  `json:"location_name"`
	Events       []Event `json:"events" gorm:"foreignkey:LocationId"`
}

type Event struct {
	gorm.Model
	EventName  string    `json:"event_name"`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndtDate"`
	LocationId uint      `json:"-"`
}

type Organization struct {
	gorm.Model
	OrganizationName string   `json:"Organization_name"`
	Persons          []Person `json:"persons" gorm:"foreignkey:OrganizationId"`
	Rooms            []Room   `json:"rooms" gorm:"foreignkey:OrganizationId"`
}

type Person struct {
	gorm.Model
	PersonName     string `json:"person_name"`
	OrganizationId uint   `json:"-"`
	TalkId         uint   `json:"-"`
}

type Room struct {
	gorm.Model
	RoomName       string `json:"room_name"`
	OrganizationId uint   `json:"-"`
}

type Language struct {
	gorm.Model
	LanguageName string `json:"room_name"`
	Talks        []Talk `json:"talks" gorm:"foreignkey:LanguageId"`
}

type Talk struct {
	gorm.Model
	TitleName  string    `json:"title_name"`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndtDate"`
	LanguageId uint      `json:"-"`
	Persons    []Person  `json:"persons" gorm:"foreignkey:TalkId"`
	Level      string    `json:"level"`
	Topics     []Topic   `json:"topics" gorm:"foreignkey:TalkId"`
}

type Topic struct {
	gorm.Model
	TopicName string   `json:"topic_name"`
	TalkId    uint     `json:"-"`
	Childs    []Child  `gorm:"many2many:topic_childs;"`
	Parents   []Parent `gorm:"many2many:topic_parents;"`
}

type Child struct {
	gorm.Model
	TopicName string `json:"topic_name"`
}

type Parent struct {
	gorm.Model
	TopicName string `json:"topic_name"`
}

func deleteLocation(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	inputLocationId := params["locationId"]
	// Convert `locationId` string param to uint64
	id64, _ := strconv.ParseUint(inputLocationId, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Location{})
	w.WriteHeader(http.StatusNoContent)
}

func updateLocation(w http.ResponseWriter, r *http.Request) {

	var updatedLocation Location
	json.NewDecoder(r.Body).Decode(&updatedLocation)
	db.Save(&updatedLocation)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedLocation)

}

func getLocation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputLocationId := params["locationId"]

	var location Location
	db.Preload("Events").First(&location, inputLocationId)
	json.NewEncoder(w).Encode(location)

}

func getLocations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var locations []Location
	db.Preload("Events").Find(&locations)
	json.NewEncoder(w).Encode(locations)

}

func createLocations(w http.ResponseWriter, r *http.Request) {

	var location Location
	json.NewDecoder(r.Body).Decode(&location)

	db.Create(&location)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)

}

func getAllEvents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var events []Event
	db.Find(&events)
	json.NewEncoder(w).Encode(events)

}

func getOneEvent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputEventId := params["eventId"]

	var event Event
	db.First(&event, inputEventId)
	json.NewEncoder(w).Encode(event)

}
