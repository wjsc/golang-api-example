```
curl -XGET http://localhost:8080/guitar/3
```

```
curl -XPOST -H'content-type: application/json' http://localhost:8080/guitar -d'{"id":3, "year":1995, "model":"Les paul"}'
```