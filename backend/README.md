# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

docker build . -t backend-go

# Available REST API's:

## EVETS

### POJO

```
type Event struct {
	gorm.Model
	EventName string  `json:"event_name"`
	StartDate time.Time `json:"StartDate"`
	EndDate time.Time `json:"EndtDate"`
	LocationId uint   `json:"-"`
}

## LOCATIONS

### POJO

```
type Location struct {
	gorm.Model
	LocationName string `json:"location_name"`
	Events []Event    `json:"events" gorm:"foreignkey:LocationId"`
}
```

### REST
Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/locations", createLocations).Methods("POST") | http://localhost:9090/locations | POST |  | 
sm.HandleFunc("/locations", getLocations).Methods("GET") | http://localhost:9090/locations | GET |  | 
sm.HandleFunc("/locations/{locationId}", getLocation).Methods("GET") | http://localhost:9090/locations/1 | GET |  | 
sm.HandleFunc("/locations/{locationId}", updateLocation).Methods("PUT") | http://localhost:9090/locations/1 | PUT |  | 
sm.HandleFunc("/locations/{locationId}", deleteLocation).Methods("DELETE") | http://localhost:9090/locations/1 | DELETE |  | 
