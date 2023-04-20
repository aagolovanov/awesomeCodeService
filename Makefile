run:
	go run ./cmd -TTL 300


build: clean
	go build -o build/runner ./...


test:
	go test -cover ./...


clean:
	rm -rf build


docker:
	docker build -t supermantelecomworker/awesomecodeservice:latest .

compose:
	docker-compose -f docker-compose.yaml -p awesomecodeservice up -d
