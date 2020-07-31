# message-api
Simple API with Golang




# how to post

```
curl --location --request POST 'http://localhost:8080/message' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message" : "hello John"
}'
```