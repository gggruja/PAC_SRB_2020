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

	var dbUrl string

	fmt.Println(dbUrl)
	dbUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cnf.DbUser, cnf.DbPassword, cnf.DbHost, cnf.DbPort, cnf.DbName)
	fmt.Println(dbUrl)

	db, err = gorm.Open(cnf.DbDriver, dbUrl)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	sm := mux.NewRouter()

	sm.HandleFunc("/", healthChecking)
	sm.Handle("/metrics", promhttp.Handler())
	sm.HandleFunc("/init", DbInit).Methods("GET")

	// CRUD API's
	// Locations
	sm.HandleFunc("/locations", createLocations).Methods("POST")
	sm.HandleFunc("/locations", getLocations).Methods("GET")
	sm.HandleFunc("/locations/{locationId:[0-9]+}", getLocation).Methods("GET")
	sm.HandleFunc("/locations/{locationId:[0-9]+}", updateLocation).Methods("PUT")
	sm.HandleFunc("/locations/{locationId:[0-9]+}", deleteLocation).Methods("DELETE")

	// Events
	sm.HandleFunc("/events", getEvents).Methods("GET")
	sm.HandleFunc("/events/{eventId:[0-9]+}", getEvent).Methods("GET")
	sm.HandleFunc("/events", createEvent).Methods("POST")
	sm.HandleFunc("/events/{eventId:[0-9]+}", updateEvent).Methods("PUT")
	sm.HandleFunc("/events/{eventId:[0-9]+}", deleteEvent).Methods("DELETE")

	// Organizations
	sm.HandleFunc("/organizations", getOrganizations).Methods("GET")
	sm.HandleFunc("/organizations/{organizationId:[0-9]+}", getOrganization).Methods("GET")
	sm.HandleFunc("/organizations", createOrganization).Methods("POST")
	sm.HandleFunc("/organizations/{organizationId:[0-9]+}", updateOrganization).Methods("PUT")
	sm.HandleFunc("/organizations/{organizationId:[0-9]+}", deleteOrganization).Methods("DELETE")

	// Persons
	sm.HandleFunc("/persons", getPersons).Methods("GET")
	sm.HandleFunc("/persons/{personId:[0-9]+}", getPerson).Methods("GET")
	sm.HandleFunc("/persons", createPerson).Methods("POST")
	sm.HandleFunc("/persons/{personId:[0-9]+}", updatePerson).Methods("PUT")
	sm.HandleFunc("/persons/{personId:[0-9]+}", deletePerson).Methods("DELETE")

	// Rooms
	sm.HandleFunc("/rooms", getRooms).Methods("GET")
	sm.HandleFunc("/rooms/{roomId:[0-9]+}", getRoom).Methods("GET")
	sm.HandleFunc("/rooms", createRoom).Methods("POST")
	sm.HandleFunc("/rooms/{roomId:[0-9]+}", updateRoom).Methods("PUT")
	sm.HandleFunc("/rooms/{roomId:[0-9]+}", deleteRoom).Methods("DELETE")

	// Topics
	sm.HandleFunc("/topics", getTopics).Methods("GET")
	sm.HandleFunc("/topics/{topicId:[0-9]+}", getTopic).Methods("GET")
	sm.HandleFunc("/topics", createTopic).Methods("POST")
	sm.HandleFunc("/topics/{topicId:[0-9]+}", updateTopic).Methods("PUT")
	sm.HandleFunc("/topics/{topicId:[0-9]+}", deleteTopic).Methods("DELETE")

	// Talks
	sm.HandleFunc("/talks", getTalks).Methods("GET")
	sm.HandleFunc("/talks/{talkId:[0-9]+}", getTalk).Methods("GET")
	sm.HandleFunc("/talks", createTalk).Methods("POST")
	sm.HandleFunc("/talks/{talkId:[0-9]+}", updateTalk).Methods("PUT")
	sm.HandleFunc("/talks/{talkId:[0-9]+}", deleteTalk).Methods("DELETE")

	// VIEW API's
	sm.HandleFunc("/api/events", getListOfAllEvents).Methods("GET")
	sm.HandleFunc("/api/persons", getPersons).Methods("GET")
	sm.HandleFunc("/api/persons/{personId:[0-9]+}/talks", getAllTalksForOnePerson).Methods("GET")

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

func DbInit(rw http.ResponseWriter, r *http.Request) {

	log.Println("Init DB!")

	rw.Header().Set("Content-Type", "text/plain")

	// Drop
	db.DropTable(&Location{}, &Event{}, &Organization{}, &Person{}, &Room{},
		&Language{}, &Talk{}, &Topic{}, &Child{})

	// Create new one
	db.AutoMigrate(&Location{}, &Event{}, &Organization{}, &Person{}, &Room{},
		&Language{}, &Talk{}, &Topic{}, &Child{})

	// Create records
	db.Create(&Location{LocationName: "Beograd"})
	db.Create(&Location{LocationName: "Smederevo"})

	db.Create(&Event{EventName: "Event in Belgrade", StartDate: time.Date(2021, time.Month(2), 12, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, time.Month(2), 15, 0, 0, 0, 0, time.UTC), LocationId: 1})
	db.Create(&Event{EventName: "Event in Smederevo", StartDate: time.Date(2021, time.Month(3), 12, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2021, time.Month(3), 15, 0, 0, 0, 0, time.UTC), LocationId: 2})

	db.Create(&Room{RoomName: "Hawaii", LocationId: 1})
	db.Create(&Room{RoomName: "Bora Bora", LocationId: 2})

	db.Create(&Organization{OrganizationName: "PRODYNA d.o.o"})
	db.Create(&Organization{OrganizationName: "PRODYNA SE"})

	personGG := Person{
		PersonName:     "Goran Grujic",
		OrganizationId: 1,
	}

	personZZ := Person{
		PersonName:     "Zoran Zuric",
		OrganizationId: 1,
	}

	personTB := Person{
		PersonName:     "Torben Bock",
		OrganizationId: 2,
	}

	personDK := Person{
		PersonName:     "Darko Krizic",
		OrganizationId: 2,
	}

	db.Create(personGG)
	db.Create(personZZ)
	db.Create(personTB)
	db.Create(personDK)

	db.Create(&Language{LanguageName: "Serbian"})
	db.Create(&Language{LanguageName: "English"})
	db.Create(&Language{LanguageName: "German"})

	db.Create(&Talk{TitleName: "CKAD - Kubernetes Development", StartDate: time.Date(2021, time.Month(2), 12, 12, 0, 0, 0, time.UTC), EndDate: time.Date(2021, time.Month(2), 12, 13, 0, 0, 0, time.UTC), LanguageId: 1, Level: "Junior", RoomId: 1, People:[]Person{personGG, personZZ}})
	db.Create(&Talk{TitleName: "Event-driven microservices: what can go wrong?", StartDate: time.Date(2021, time.Month(2), 12, 13, 0, 0, 0, time.UTC), EndDate: time.Date(2021, time.Month(2), 12, 14, 0, 0, 0, time.UTC), LanguageId: 2, Level: "Junior", RoomId: 2, People:[]Person{personDK, personTB}})

	var childs = []Child{
		Child{
			TopicName: "Kubernetes Child 1",
		},
		Child{
			TopicName: "Kubernetes Child 2",
		},
	}

	db.Create(&Topic{TopicName: "Kubernetes", TalkId: 1, Children: childs})
	db.Create(&Topic{TopicName: "Exam", TalkId: 1})
	db.Create(&Topic{TopicName: "Event-driven microservices", TalkId: 2})
	db.Create(&Topic{TopicName: "Microservices", TalkId: 2})
	db.Create(&Topic{TopicName: "Event-driven", TalkId: 2})

	log.Println("Init DB, done!")

	rw.WriteHeader(http.StatusOK)

}

type Location struct {
	gorm.Model
	LocationName string  `json:"LocationName"`
	Events       []Event `json:"Events" gorm:"foreignkey:LocationId"`
	Rooms        []Room  `json:"Rooms" gorm:"foreignkey:LocationId"`
}

type Room struct {
	gorm.Model
	RoomName   string `json:"RoomName"`
	Talk       Talk   `gorm:"foreignkey:TalkId"`
	LocationId uint   `json:"-"`
}

type Event struct {
	gorm.Model
	EventName  string    `json:"EventName"`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndDate"`
	LocationId uint      `json:"-"`
}

type Organization struct {
	gorm.Model
	OrganizationName string   `json:"OrganizationName"`
	People           []Person `json:"People" gorm:"foreignkey:OrganizationId"`
}

type Person struct {
	gorm.Model
	PersonName     string `json:"PersonName"`
	OrganizationId uint   `json:"-"`
}

type Language struct {
	gorm.Model
	LanguageName string `json:"LanguageName"`
	Talks        []Talk `json:"Talks" gorm:"foreignkey:LanguageId"`
}

type Talk struct {
	gorm.Model
	TitleName  string    `json:"TitleName"`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndDate"`
	LanguageId uint      `json:"-"`
	People     []Person  `gorm:"many2many:talks_persons;"`
	Level      string    `json:"Level"`
	Topics     []Topic   `json:"Topics" gorm:"foreignkey:TalkId"`
	RoomId     uint      `json:"RoomId"`
}

type Topic struct {
	gorm.Model
	TopicName string  `json:"TopicName"`
	TalkId    uint    `json:"-"`
	Children  []Child `gorm:"many2many:topic_children;"`
}

type Child struct {
	gorm.Model
	TopicName string `json:"TopicName"`
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(updatedLocation)

}

func getLocation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputLocationId := params["locationId"]

	var location Location
	db.Preload("Events").Preload("Rooms").First(&location, inputLocationId)
	json.NewEncoder(w).Encode(location)

}

func getLocations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var locations []Location
	db.Preload("Events").Preload("Rooms").Find(&locations)
	json.NewEncoder(w).Encode(locations)

}

func createLocations(w http.ResponseWriter, r *http.Request) {

	var location Location
	json.NewDecoder(r.Body).Decode(&location)

	db.Create(&location)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(location)

}

func getEvents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var events []Event
	db.Find(&events)
	json.NewEncoder(w).Encode(events)

}

func getEvent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	inputEventId := params["eventId"]

	var event Event
	db.First(&event, inputEventId)
	json.NewEncoder(w).Encode(event)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputEventId := params["eventId"]

	id64, _ := strconv.ParseUint(inputEventId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Event{})

	w.WriteHeader(http.StatusNoContent)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

	var updatedEvent Event

	json.NewDecoder(r.Body).Decode(&updatedEvent)
	db.Save(&updatedEvent)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedEvent)

}

func createEvent(w http.ResponseWriter, r *http.Request) {

	var event Event
	json.NewDecoder(r.Body).Decode(&event)

	db.Create(&event)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(event)
}

func deleteTalk(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputTalkId := params["talkId"]

	id64, _ := strconv.ParseUint(inputTalkId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Talk{})

	w.WriteHeader(http.StatusNoContent)

}

func updateTalk(w http.ResponseWriter, r *http.Request) {

	var updatedTalk Talk

	json.NewDecoder(r.Body).Decode(&updatedTalk)
	db.Save(&updatedTalk)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedTalk)

}

func createTalk(w http.ResponseWriter, r *http.Request) {

	var talk Talk
	json.NewDecoder(r.Body).Decode(&talk)

	db.Create(&talk)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(talk)

}

func getTalk(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputTalkId := params["talkId"]

	var talk Talk
	db.Preload("People").Preload("Topics").First(&talk, inputTalkId)
	json.NewEncoder(w).Encode(talk)

}

func getTalks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var talks []Talk
	db.Preload("People").Preload("Topics").Find(&talks)
	json.NewEncoder(w).Encode(talks)

}

func deleteTopic(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputTopicId := params["topicId"]

	id64, _ := strconv.ParseUint(inputTopicId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Topic{})

	w.WriteHeader(http.StatusNoContent)
}

func updateTopic(w http.ResponseWriter, r *http.Request) {

	var updatedTopic Topic

	json.NewDecoder(r.Body).Decode(&updatedTopic)
	db.Save(&updatedTopic)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedTopic)
}

func createTopic(w http.ResponseWriter, r *http.Request) {

	var topic Topic
	json.NewDecoder(r.Body).Decode(&topic)

	db.Create(&topic)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(topic)

}

func getTopic(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputTopicId := params["topicId"]

	var topic Topic
	db.Preload("Children").First(&topic, inputTopicId)
	json.NewEncoder(w).Encode(topic)

}

func getTopics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var topics []Topic
	db.Preload("Children").Find(&topics)
	json.NewEncoder(w).Encode(topics)

}

func deleteRoom(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputRoomId := params["roomId"]

	id64, _ := strconv.ParseUint(inputRoomId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Room{})

	w.WriteHeader(http.StatusNoContent)
}

func updateRoom(w http.ResponseWriter, r *http.Request) {

	var updatedRoom Room

	json.NewDecoder(r.Body).Decode(&updatedRoom)
	db.Save(&updatedRoom)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedRoom)

}

func createRoom(w http.ResponseWriter, r *http.Request) {

	var room Room
	json.NewDecoder(r.Body).Decode(&room)

	db.Create(&room)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(room)

}

func getRoom(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputRoomId := params["roomId"]

	var room Room
	db.Preload("Talk").First(&room, inputRoomId)
	json.NewEncoder(w).Encode(room)

}

func getRooms(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rooms []Room
	db.Preload("Talk").Find(&rooms)
	json.NewEncoder(w).Encode(rooms)

}

func deletePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputPersonId := params["personId"]

	id64, _ := strconv.ParseUint(inputPersonId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Person{})

	w.WriteHeader(http.StatusNoContent)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {

	var updatedPeston Person

	json.NewDecoder(r.Body).Decode(&updatedPeston)
	db.Save(&updatedPeston)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedPeston)

}

func createPerson(w http.ResponseWriter, r *http.Request) {

	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	db.Create(&person)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(person)

}

func getPerson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputPersonId := params["personId"]

	var person Person
	db.First(&person, inputPersonId)
	json.NewEncoder(w).Encode(person)

}

func getPersons(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var people []Person
	db.Find(&people)
	json.NewEncoder(w).Encode(people)

}

func deleteOrganization(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	inputOrganizationId := params["organizationId"]

	id64, _ := strconv.ParseUint(inputOrganizationId, 10, 64)
	idToDelete := uint(id64)

	db.Where("id = ?", idToDelete).Delete(&Organization{})

	w.WriteHeader(http.StatusNoContent)

}

func updateOrganization(w http.ResponseWriter, r *http.Request) {

	var updatedOrganization Organization

	json.NewDecoder(r.Body).Decode(&updatedOrganization)
	db.Save(&updatedOrganization)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(updatedOrganization)

}

func createOrganization(w http.ResponseWriter, r *http.Request) {

	var organization Organization
	json.NewDecoder(r.Body).Decode(&organization)

	db.Create(&organization)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(organization)

}

func getOrganization(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	inputOrganizationId := params["organizationId"]

	var organization Organization
	db.Preload("People").First(&organization, inputOrganizationId)
	json.NewEncoder(w).Encode(organization)

}

func getOrganizations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var organization []Organization
	db.Preload("People").Find(&organization)
	json.NewEncoder(w).Encode(organization)

}

type EventResult struct {
	EventName    string
	StartDate    time.Time
	EndDate      time.Time
	LocationName string
	RoomName     string
	TitleName    string
	TopicName    string
}

func getListOfAllEvents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var events []EventResult

	db.Table("locations").
		Select("events.event_name, events.start_date, events.end_date, location_name, rooms.room_name, talks.title_name, topics.topic_name").
		Joins("JOIN rooms on locations.id = rooms.location_id").
		Joins("JOIN events on locations.id = events.location_id").
		Joins("JOIN talks on rooms.id = talks.room_id").
		Joins("JOIN topics on talks.id = topics.talk_id").Scan(&events)

	json.NewEncoder(w).Encode(events)

}

func getAllTalksForOnePerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["personId"], 10, 32)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var talks []*Talk

	db.Preload("People").
		Preload("Topics").
		Where("id IN (?)", db.Table("talks_persons").Select("talk_id").Where("person_id = ?", id).SubQuery()).
		Find(&talks)

	json.NewEncoder(w).Encode(talks)
}
