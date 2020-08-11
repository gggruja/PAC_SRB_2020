# PAC Serbia 2020

This is a public repository, for PAC 2019 Task 1.1

docker build . -t backend-go

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
	TalkId         uint   `json:"TalkId"`
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
	People     []Person  `json:"People" gorm:"foreignkey:TalkId"`
	Level      string    `json:"Level"`
	Topics     []Topic   `json:"Topics" gorm:"foreignkey:TalkId"`
	RoomId     uint      `json:"RoomId"`
}

type Topic struct {
	gorm.Model
	TopicName string   `json:"TopicName"`
	TalkId    uint     `json:"-"`
	Children  []Child  `gorm:"many2many:topic_children;"`
	Parents   []Parent `gorm:"many2many:topic_parents;"`
}

type Child struct {
	gorm.Model
	TopicName string `json:"TopicName"`
}

type Parent struct {
	gorm.Model
	TopicName string `json:"TopicName"`
}
```

### SETTINGS API's
Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/", healthChecking) | http://localhost:9090/ | GET | Health checking for Kubernetes probes | 
sm.Handle("/metrics", promhttp.Handler()) | http://localhost:9090/metrics | GET | Metrics scraper API, for prometheus | 
sm.HandleFunc("/init", DbInit).Methods("GET") | http://localhost:9090/init | GET | Database initialization, run by kubernetes job | 

### VIEW API's
Code | REST API | METHOD | COMMENT |
--- | --- | --- | --- |
sm.HandleFunc("/api/events", getListOfAllEvents).Methods("GET") | http://localhost:9090/api/events | GET | Event view | 

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

### CRUD REST API's
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
        "CreatedAt": "2020-08-04T14:13:13Z",
        "UpdatedAt": "2020-08-04T14:13:13Z",
        "DeletedAt": null,
        "LocationName": "Beograd",
        "Events": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T14:13:13Z",
                "UpdatedAt": "2020-08-04T14:13:13Z",
                "DeletedAt": null,
                "EventName": "Heapcon Belgrade",
                "StartDate": "2020-08-04T14:13:13Z",
                "EndDate": "2020-08-04T14:13:13Z"
            }
        ]
    },
    {
        "ID": 2,
        "CreatedAt": "2020-08-04T14:13:13Z",
        "UpdatedAt": "2020-08-04T14:13:13Z",
        "DeletedAt": null,
        "LocationName": "Smederevo",
        "Events": []
    }
]
```

#### GET ONE LOCATION AND ALL EVENTS
```
{
    "ID": 1,
    "CreatedAt": "2020-08-04T14:13:13Z",
    "UpdatedAt": "2020-08-04T14:13:13Z",
    "DeletedAt": null,
    "LocationName": "Beograd",
    "Events": [
        {
            "ID": 1,
            "CreatedAt": "2020-08-04T14:13:13Z",
            "UpdatedAt": "2020-08-04T14:13:13Z",
            "DeletedAt": null,
            "EventName": "Heapcon Belgrade",
            "StartDate": "2020-08-04T14:13:13Z",
            "EndDate": "2020-08-04T14:13:13Z"
        }
    ]
}
```

#### GET ALL PERSONS
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T14:13:13Z",
        "UpdatedAt": "2020-08-04T14:13:13Z",
        "DeletedAt": null,
        "PersonName": "Goran Grujic",
        "OrganizationId": 1,
        "TalkId": 1
    }
]

```
#### GET TOPICS PER PERSON
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T14:13:14Z",
        "UpdatedAt": "2020-08-04T14:13:14Z",
        "DeletedAt": null,
        "TitleName": "CKAD - Kubernetes Development",
        "StartDate": "2020-08-04T14:13:14Z",
        "EndDate": "2020-08-04T14:13:14Z",
        "LanguageId": 1,
        "People": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T14:13:13Z",
                "UpdatedAt": "2020-08-04T14:13:13Z",
                "DeletedAt": null,
                "PersonName": "Goran Grujic",
                "OrganizationId": 1,
                "TalkId": 1
            }
        ],
        "Level": "Beginner",
        "Topics": null
    }
]
```
#### GET TOPICS AND TALKS AND PERSONS
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T14:13:14Z",
        "UpdatedAt": "2020-08-04T14:13:14Z",
        "DeletedAt": null,
        "TitleName": "CKAD - Kubernetes Development",
        "StartDate": "2020-08-04T14:13:14Z",
        "EndDate": "2020-08-04T14:13:14Z",
        "LanguageId": 1,
        "People": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T14:13:13Z",
                "UpdatedAt": "2020-08-04T14:13:13Z",
                "DeletedAt": null,
                "PersonName": "Goran Grujic",
                "OrganizationId": 1,
                "TalkId": 1
            }
        ],
        "Level": "Beginner",
        "Topics": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T14:13:14Z",
                "UpdatedAt": "2020-08-04T14:13:14Z",
                "DeletedAt": null,
                "TopicName": "Kubernetes",
                "TalkId": 1,
                "Children": null,
                "Parents": null
            },
            {
                "ID": 2,
                "CreatedAt": "2020-08-04T14:13:14Z",
                "UpdatedAt": "2020-08-04T14:13:14Z",
                "DeletedAt": null,
                "TopicName": "Exam",
                "TalkId": 1,
                "Children": null,
                "Parents": null
            }
        ]
    },
    {
        "ID": 2,
        "CreatedAt": "2020-08-04T14:13:14Z",
        "UpdatedAt": "2020-08-04T14:13:14Z",
        "DeletedAt": null,
        "TitleName": "Weed - Rolling papers",
        "StartDate": "2020-08-04T14:13:14Z",
        "EndDate": "2020-08-04T14:13:14Z",
        "LanguageId": 1,
        "People": [],
        "Level": "Architect",
        "Topics": [
            {
                "ID": 3,
                "CreatedAt": "2020-08-04T14:13:14Z",
                "UpdatedAt": "2020-08-04T14:13:14Z",
                "DeletedAt": null,
                "TopicName": "Rolling Pappers",
                "TalkId": 2,
                "Children": null,
                "Parents": null
            },
            {
                "ID": 4,
                "CreatedAt": "2020-08-04T14:13:14Z",
                "UpdatedAt": "2020-08-04T14:13:14Z",
                "DeletedAt": null,
                "TopicName": "Weed",
                "TalkId": 2,
                "Children": null,
                "Parents": null
            }
        ]
    }
]
```
