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
sm.HandleFunc("/persons", getPersons).Methods("GET") | http://localhost:9090/persons | GET |  | 
sm.HandleFunc("/talks/person/{personId}", getTalks).Methods("GET") | http://localhost:9090//talks/person/1 | GET |  | 
sm.HandleFunc("/topics", getTopics).Methods("GET") | http://localhost:9090//topics | GET |  | 

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
