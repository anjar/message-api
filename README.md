# simple-message-api
Simple API with Golang


## How to install
1. Clone this repo.
2. Create .env from .env.example
3. Create new database, and run ```go run migrate.go```
4. Run service with ```go run main.go```
5. Open http://localhost:8080

## Run with Docker
1. go to repo dir ```cd /path/to/repo/simpleapi```
2. Run ```docker-compose up```
3. Open http://localhost/client or try below endpoint

## How to Post

Using curl: 
```
curl --location --request POST 'http://localhost:8080/message' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message" : "hello John"
}'
```

## How to get data
```
curl --location --request GET 'http://localhost:8080/message/list'
```

## Websocket

Open http://localhost:8080/client in multiple browser, fill message and hit send button, in other browser, you should retrieve the message in realtime.
