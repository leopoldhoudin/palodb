build: palodbd palodb

palodbd:
	go build -o bin/palodbd ./cmd/palodbd

palodb:
	go build -o bin/palodb ./cmd/palodb

test:
	go test -v github.com/leopoldhoudin/palodb/core/lang
