build: palodbd

palodbd:
	go build -o bin/palodbd ./cmd/palodbd

test:
	go test -v github.com/leopoldhoudin/palodb/core/lang
