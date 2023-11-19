# Log Indigester

## Setup

1. Run Docker compose command to setup the environment. (in the directory where the docker compose file is)
```
    docker-compose up -d .
```

Following services are running on your system now

1. Elastic Search -> port 9200
2. Kibana -> port 5601
3. Golang app -> port 3000
4. Grafana -> port 3001

## ENV VARIABLES

Following environment variables must be set before running the application. This should be saved in /api/.env

1. INDEX_NAME -> name of the index in the elastic search where the log data will be saved.
2. API_KEY -> api key of the elastic search.

**NOTE: The server can be hit with the post request at http://localhost:3000/api/v1/**

## Create an Index

You must create an index in elastic search. This index will hold all the documents representing the logs.


## Mappings 

```
{
  "properties": {
    "level": {"type": "text"},
	"message": {"type": "text"},
    "resourceId": {"type": "text"},
	"timestamp": {"type": "date"},
	"traceId": {"type":"keyword"},
    "spanId": {"type":"keyword"},
    "commit": {"type": "text"},
    "metadata": { "type": "object",
      "properties":
        {"parentResourceId": {"type":"keyword"}}
    }
  }
}
```


## Screenshots
![image](https://github.com/ArshpreetS/log-eater/assets/76895787/5b584de0-56d0-4808-b138-35cdd1544dd8)
![image](https://github.com/ArshpreetS/log-eater/assets/76895787/9844e46b-9e06-43dd-a32d-f6211a32d0a7)
![image](https://github.com/ArshpreetS/log-eater/assets/76895787/126aff45-15ee-4f3e-9838-f3242205f10b)

