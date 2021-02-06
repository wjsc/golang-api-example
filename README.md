go build && ./app-v1

https://github.com/gin-gonic/gin

```
curl -XGET http://localhost:8080/guitar | jq .
curl -XPOST -H'content-type: application/json' http://localhost:8080/guitar -d'{"id":1, "year":1965, "model":"Telecaster"}' | jq .
curl -XPOST -H'content-type: application/json' http://localhost:8080/guitar -d'{"id":2, "year":1975, "model":"Stratocaster"}' | jq .
curl -XPOST -H'content-type: application/json' http://localhost:8080/guitar -d'{"id":3, "year":1985, "model":"Les XXXX"}' | jq .
curl -XPUT -H'content-type: application/json' http://localhost:8080/guitar/3 -d'{"id":3, "year":1985, "model":"Les Paul"}' | jq .
curl -XGET http://localhost:8080/guitar | jq .
curl -XGET http://localhost:8080/guitar/3 | jq .
curl -XDELETE http://localhost:8080/guitar/1 | jq .
```