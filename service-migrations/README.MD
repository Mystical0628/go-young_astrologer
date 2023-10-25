Create migration
```
migrate create -ext sql -dir migrations {name}
```

Up migrations
```
migrate -path ./migrations -database postgres://astrologer:12345678@localhost:9432/postgres?sslmode=disable up
```