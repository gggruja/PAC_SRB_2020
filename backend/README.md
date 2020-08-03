# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

docker build . -t backend-go

# Available REST API's:

## EVENTS

### POJO

```
type Event struct {
	gorm.Model
	EventName string  `json:"event_name"`
	StartDate time.Time `json:"StartDate"`
	EndDate time.Time `json:"EndtDate"`
	LocationId uint   `json:"-"`
}
```

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

#### GET ALL locations and there event
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-03T22:50:31Z",
        "UpdatedAt": "2020-08-03T22:50:32Z",
        "DeletedAt": null,
        "location_name": "Beograd",
        "events": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-03T22:50:09Z",
                "UpdatedAt": "2020-08-03T22:50:10Z",
                "DeletedAt": null,
                "event_name": "PAC",
                "StartDate": "0000-00-00",
                "EndtDate": "0000-00-00"
            },
            {
                "ID": 2,
                "CreatedAt": "2020-08-03T22:56:18Z",
                "UpdatedAt": "2020-08-03T22:56:21Z",
                "DeletedAt": null,
                "event_name": "Kubernetes",
                "StartDate": "0000-00-00",
                "EndtDate": "0000-00-00"
            }
        ]
    }
]
```

#### GET ONE LOCATION AMD ALL EVENTS
```
{
    "ID": 1,
    "CreatedAt": "2020-08-03T22:50:31Z",
    "UpdatedAt": "2020-08-03T22:50:32Z",
    "DeletedAt": null,
    "location_name": "Beograd",
    "events": [
        {
            "ID": 1,
            "CreatedAt": "2020-08-03T22:50:09Z",
            "UpdatedAt": "2020-08-03T22:50:10Z",
            "DeletedAt": null,
            "event_name": "PAC",
            "StartDate": "2020-08-03T22:58:13Z",
            "EndtDate": "2020-08-03T22:58:16Z"
        },
        {
            "ID": 2,
            "CreatedAt": "2020-08-03T22:56:18Z",
            "UpdatedAt": "2020-08-03T22:56:21Z",
            "DeletedAt": null,
            "event_name": "Kubernetes",
            "StartDate": "2020-08-03T22:58:17Z",
            "EndtDate": "2020-08-03T22:58:18Z"
        }
    ]
}
```
