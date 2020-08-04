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

#### GET ONE LOCATION AND ALL EVENTS
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

#### GET ALL PERSONS
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T12:13:44Z",
        "UpdatedAt": "2020-08-04T12:13:44Z",
        "DeletedAt": null,
        "person_name": "Goran Grujic",
        "organizationId": 1,
        "talkId": 1
    }
]

```
#### GET TOPICS PER PERSON
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T12:33:49Z",
        "UpdatedAt": "2020-08-04T12:33:49Z",
        "DeletedAt": null,
        "title_name": "CKAD - Kubernetes Development",
        "StartDate": "2020-08-04T12:33:49Z",
        "EndtDate": "2020-08-04T12:33:49Z",
        "persons": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T12:33:49Z",
                "UpdatedAt": "2020-08-04T12:33:49Z",
                "DeletedAt": null,
                "person_name": "Goran Grujic",
                "organizationId": 1,
                "talkId": 1
            }
        ],
        "level": "Beginner",
        "topics": null
    }
]
```
#### GET TOPICS AND TALKS AND PERSONS
```
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-04T13:06:26Z",
        "UpdatedAt": "2020-08-04T13:06:26Z",
        "DeletedAt": null,
        "title_name": "CKAD - Kubernetes Development",
        "StartDate": "2020-08-04T13:06:26Z",
        "EndtDate": "2020-08-04T13:06:26Z",
        "persons": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T13:06:26Z",
                "UpdatedAt": "2020-08-04T13:06:26Z",
                "DeletedAt": null,
                "person_name": "Goran Grujic",
                "organizationId": 1,
                "talkId": 1
            }
        ],
        "level": "Beginner",
        "topics": [
            {
                "ID": 1,
                "CreatedAt": "2020-08-04T13:06:26Z",
                "UpdatedAt": "2020-08-04T13:06:26Z",
                "DeletedAt": null,
                "topic_name": "Kubernetes",
                "TalkId": 1,
                "Childs": null,
                "Parents": null
            },
            {
                "ID": 2,
                "CreatedAt": "2020-08-04T13:06:27Z",
                "UpdatedAt": "2020-08-04T13:06:27Z",
                "DeletedAt": null,
                "topic_name": "Exam",
                "TalkId": 1,
                "Childs": null,
                "Parents": null
            }
        ]
    },
    {
        "ID": 2,
        "CreatedAt": "2020-08-04T13:06:26Z",
        "UpdatedAt": "2020-08-04T13:06:26Z",
        "DeletedAt": null,
        "title_name": "Weed - Rolling papers",
        "StartDate": "2020-08-04T13:06:26Z",
        "EndtDate": "2020-08-04T13:06:26Z",
        "persons": [],
        "level": "Architect",
        "topics": [
            {
                "ID": 3,
                "CreatedAt": "2020-08-04T13:06:27Z",
                "UpdatedAt": "2020-08-04T13:06:27Z",
                "DeletedAt": null,
                "topic_name": "Rolling Pappers",
                "TalkId": 2,
                "Childs": null,
                "Parents": null
            },
            {
                "ID": 4,
                "CreatedAt": "2020-08-04T13:06:27Z",
                "UpdatedAt": "2020-08-04T13:06:27Z",
                "DeletedAt": null,
                "topic_name": "Weed",
                "TalkId": 2,
                "Childs": null,
                "Parents": null
            }
        ]
    }
]
```
