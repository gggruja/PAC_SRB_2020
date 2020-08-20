# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

docker build . -t backend

# Available REST API's:

## POJO's

```
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
	LocationId uint   `json:"LocationId"`
}

type Event struct {
	gorm.Model
	EventName  string    `json:"EventName"`
	StartDate  time.Time `json:"StartDate"`
	EndDate    time.Time `json:"EndDate"`
	LocationId uint      `json:"LocationId"`
}

type Organization struct {
	gorm.Model
	OrganizationName string   `json:"OrganizationName"`
	People           []Person `json:"People" gorm:"foreignkey:OrganizationId"`
}

type Person struct {
	gorm.Model
	PersonName     string `json:"PersonName"`
	OrganizationId uint   `json:"OrganizationId"`
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
	LanguageId uint      `json:"LanguageId"`
	People     []Person  `gorm:"many2many:talks_persons;"`
	Level      string    `json:"Level"`
	Topics     []Topic   `json:"Topics" gorm:"foreignkey:TalkId"`
	RoomId     uint      `json:"RoomId"`
}

type Topic struct {
	gorm.Model
	TopicName string  `json:"TopicName"`
	TalkId    uint    `json:"TalkId"`
	Children  []Child `gorm:"many2many:topic_children;"`
}

type Child struct {
	gorm.Model
	TopicName string `json:"TopicName"`
}

```

Note: If you run it in minikube BASE_URL is http://conference.backend/*********

### SETTINGS API's
Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/", healthChecking) | http://localhost:9090/ | GET | Health checking for Kubernetes probes | 
sm.Handle("/metrics", promhttp.Handler()) | http://localhost:9090/metrics | GET | Metrics scraper API, for prometheus | 
sm.HandleFunc("/init", DbInit).Methods("GET") | http://localhost:9090/init | GET | Database initialization, run by kubernetes job | 

### CRUD REST API's

Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/locations", createLocations).Methods("POST") | http://localhost:9090/locations | POST |  | 
sm.HandleFunc("/locations", getLocations).Methods("GET") | http://localhost:9090/locations | GET |  | 
sm.HandleFunc("/locations/{locationId:[0-9]+}", getLocation).Methods("GET") | http://localhost:9090/locations/1 | GET |  | 
sm.HandleFunc("/locations/{locationId:[0-9]+}", updateLocation).Methods("PUT") | http://localhost:9090/locations/1 | PUT |  | 
sm.HandleFunc("/locations/{locationId:[0-9]+}", deleteLocation).Methods("DELETE") | http://localhost:9090/locations/1 | DELETE |  | 
sm.HandleFunc("/events", getEvents).Methods("GET") | http://localhost:9090/events | GET |  | 
sm.HandleFunc("/events/{eventId:[0-9]+}", getEvent).Methods("GET") | http://localhost:9090/events/1 | GET |  | 
sm.HandleFunc("/events", createEvent).Methods("POST") | http://localhost:9090/events | POST |  | 
sm.HandleFunc("/events/{eventId:[0-9]+}", updateEvent).Methods("PUT") | http://localhost:9090/events/1 | PUT |  | 
sm.HandleFunc("/events/{eventId:[0-9]+}", deleteEvent).Methods("DELETE") | http://localhost:9090/events/1 | DELETE |  | 
sm.HandleFunc("/organizations", getOrganizations).Methods("GET") | http://localhost:9090/organizations | GET |  | 
sm.HandleFunc("/organizations/{organizationId:[0-9]+}", getOrganization).Methods("GET") | http://localhost:9090/organizations/1 | GET |  | 
sm.HandleFunc("/organizations", createOrganization).Methods("POST") | http://localhost:9090/organizations | POST |  | 
sm.HandleFunc("/organizations/{organizationId:[0-9]+}", updateOrganization).Methods("PUT") | http://localhost:9090/organizations/1 | PUT |  | 
sm.HandleFunc("/organizations/{organizationId:[0-9]+}", deleteOrganization).Methods("DELETE") | http://localhost:9090/organizations/1 | DELETE |  | 
sm.HandleFunc("/persons", getPersons).Methods("GET") | http://localhost:9090/persons | GET |  | 
sm.HandleFunc("/persons/{personId:[0-9]+}", getPerson).Methods("GET") | http://localhost:9090/persons/1 | GET |  | 
sm.HandleFunc("/persons", createPerson).Methods("POST") | http://localhost:9090/persons | POST |  | 
sm.HandleFunc("/persons/{personId:[0-9]+}", updatePerson).Methods("PUT") | http://localhost:9090/persons/1 | PUT |  | 
sm.HandleFunc("/persons/{personId:[0-9]+}", deletePerson).Methods("DELETE") | http://localhost:9090/persons/1 | DELETE |  | 
sm.HandleFunc("/rooms", getRooms).Methods("GET") | http://localhost:9090/rooms | GET |  | 
sm.HandleFunc("/rooms/{roomId:[0-9]+}", getRoom).Methods("GET") | http://localhost:9090/rooms/1 | GET |  | 
sm.HandleFunc("/rooms", createRoom).Methods("POST") | http://localhost:9090/rooms | POST |  | 
sm.HandleFunc("/rooms/{roomId:[0-9]+}", updateRoom).Methods("PUT") | http://localhost:9090/rooms/1 | PUT |  | 
sm.HandleFunc("/rooms/{roomId:[0-9]+}", deleteRoom).Methods("DELETE") | http://localhost:9090/rooms/1 | DELETE |  | 
sm.HandleFunc("/topics", getTopics).Methods("GET") | http://localhost:9090/topics | GET |  | 
sm.HandleFunc("/topics/{topicId:[0-9]+}", getTopic).Methods("GET") | http://localhost:9090/topics/1 | GET |  | 
sm.HandleFunc("/topics", createTopic).Methods("POST") | http://localhost:9090/topics | POST |  | 
sm.HandleFunc("/topics/{topicId:[0-9]+}", updateTopic).Methods("PUT") | http://localhost:9090/topics/1 | PUT |  | 
sm.HandleFunc("/topics/{topicId:[0-9]+}", deleteTopic).Methods("DELETE") | http://localhost:9090/topics/1 | DELETE |  | 
sm.HandleFunc("/talks", getTalks).Methods("GET") | http://localhost:9090/talks | GET |  | 
sm.HandleFunc("/talks/{talkId:[0-9]+}", getTalk).Methods("GET") | http://localhost:9090/talks/1 | GET |  | 
sm.HandleFunc("/talks", createTalk).Methods("POST") | http://localhost:9090/talks | POST |  | 
sm.HandleFunc("/talks/{talkId:[0-9]+}", updateTalk).Methods("PUT") | http://localhost:9090/talks/1 | PUT |  | 
sm.HandleFunc("/talks/{talkId:[0-9]+}", deleteTalk).Methods("DELETE") | http://localhost:9090/talks/1 | DELETE |  | 


### VIEW API's
Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/api/events", getListOfAllEvents).Methods("GET") | http://localhost:9090/api/events | GET | Event view | 
sm.HandleFunc("/api/persons", getPersons).Methods("GET") | http://localhost:9090/api/persons | GET | Person view | 
sm.HandleFunc("/api/persons/{personId:[0-9]+}/talks", getAllTalksForOnePerson).Methods("GET") | http://localhost:9090/api/persons/6/talks | GET | Person talks detail view |
sm.HandleFunc("/api/talks", getTalks).Methods("GET") | http://localhost:9090/api/talks | GET | Get all talks view |
sm.HandleFunc("/api/events/select-box", getEvents).Methods("GET") | http://localhost:9090/api/events/select-box | GET | Get all events for selecting |
sm.HandleFunc("/api/locations/{locationId:[0-9]+}/rooms", getAllRoomsAtLocation).Methods("GET") | http://localhost:9090/api/locations/1/rooms | GET | Get rooms per location |
sm.HandleFunc("/api/rooms/{roomId:[0-9]+}/talks", getAllTalksInARoom).Methods("GET") | http://localhost:9090/api/rooms/1/talks | GET | Get all talks per room |


#### Event view
```
[
    {
        "EventName": "Event in Belgrade",
        "StartDate": "2020-08-11T14:16:10+02:00",
        "EndDate": "2020-08-11T14:16:10+02:00",
        "LocationName": "Beograd",
        "RoomName": "Hawaii",
        "TitleName": "CKAD - Kubernetes Development",
        "TopicName": "Kubernetes"
    }
]
```

#### Person view
```
[
    {
        "ID": 5,
        "CreatedAt": "2020-08-18T12:32:41Z",
        "UpdatedAt": "2020-08-18T12:32:41Z",
        "DeletedAt": null,
        "PersonName": "Goran Grujic"
    }
]
```

#### Person talks view
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-18T12:32:41Z",
        "UpdatedAt": "2020-08-18T12:32:41Z",
        "DeletedAt": null,
        "TitleName": "CKAD - Kubernetes Development",
        "StartDate": "2021-02-12T12:00:00Z",
        "EndDate": "2021-02-12T13:00:00Z",
        "People": [
            {
                "ID": 5,
                "CreatedAt": "2020-08-18T12:32:41Z",
                "UpdatedAt": "2020-08-18T12:32:41Z",
                "DeletedAt": null,
                "PersonName": "Goran Grujic"
            },
            {
                "ID": 6,
                "CreatedAt": "2020-08-18T12:32:41Z",
                "UpdatedAt": "2020-08-18T12:32:41Z",
                "DeletedAt": null,
                "PersonName": "Zoran Zuric"
            }
        ],
        "Level": "Junior",
        "Topics": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-18T12:32:41Z",
                "UpdatedAt": "2020-08-18T12:32:41Z",
                "DeletedAt": null,
                "TopicName": "Kubernetes",
                "Children": null
            },
            {
                "ID": 2,
                "CreatedAt": "2020-08-18T12:32:41Z",
                "UpdatedAt": "2020-08-18T12:32:41Z",
                "DeletedAt": null,
                "TopicName": "Exam",
                "Children": null
            }
        ],
        "RoomId": 1
    }
]
```

#### Talks view
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-18T14:11:21Z",
        "UpdatedAt": "2020-08-18T14:11:21Z",
        "DeletedAt": null,
        "TitleName": "CKAD - Kubernetes Development",
        "StartDate": "2021-02-12T12:00:00Z",
        "EndDate": "2021-02-12T13:00:00Z",
        "People": [
            {
                "ID": 5,
                "CreatedAt": "2020-08-18T14:11:21Z",
                "UpdatedAt": "2020-08-18T14:11:21Z",
                "DeletedAt": null,
                "PersonName": "Goran Grujic"
            },
            {
                "ID": 6,
                "CreatedAt": "2020-08-18T14:11:21Z",
                "UpdatedAt": "2020-08-18T14:11:21Z",
                "DeletedAt": null,
                "PersonName": "Zoran Zuric"
            }
        ],
        "Level": "Junior",
        "Topics": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-18T14:11:21Z",
                "UpdatedAt": "2020-08-18T14:11:21Z",
                "DeletedAt": null,
                "TopicName": "Kubernetes",
                "Children": null
            },
            {
                "ID": 2,
                "CreatedAt": "2020-08-18T14:11:21Z",
                "UpdatedAt": "2020-08-18T14:11:21Z",
                "DeletedAt": null,
                "TopicName": "Exam",
                "Children": null
            }
        ],
        "RoomId": 1
    }
]
```

#### Events where talk is held
```
[
    {
        "EventName": "Event in Belgrade",
        "StartDate": "2021-02-12T00:00:00Z",
        "EndDate": "2021-02-15T00:00:00Z",
        "LocationName": "Beograd",
        "RoomName": "Hawaii",
        "TitleName": "CKAD - Kubernetes Development",
        "TopicName": ""
    }
]
```

#### Get all events for selecting
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-18T16:02:35Z",
        "UpdatedAt": "2020-08-18T16:02:35Z",
        "DeletedAt": null,
        "EventName": "Event in Belgrade",
        "StartDate": "2021-02-12T00:00:00Z",
        "EndDate": "2021-02-15T00:00:00Z",
        "LocationId": 1
    }
]
```

#### Get all rooms at location
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-18T16:04:12Z",
        "UpdatedAt": "2020-08-18T16:04:12Z",
        "DeletedAt": null,
        "RoomName": "Hawaii",
        "Talk": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "TitleName": "",
            "StartDate": "0001-01-01T00:00:00Z",
            "EndDate": "0001-01-01T00:00:00Z",
            "LanguageId": 0,
            "People": null,
            "Level": "",
            "Topics": null,
            "RoomId": 0
        },
        "LocationId": 1
    }
]
```

#### Get talks per room
```
[
    {
        "EventName": "Event in Belgrade",
        "StartDate": "2021-02-12T00:00:00Z",
        "EndDate": "2021-02-15T00:00:00Z",
        "LocationName": "Beograd",
        "RoomName": "Hawaii",
        "TitleName": "CKAD - Kubernetes Development",
        "TopicName": ""
    }
]
