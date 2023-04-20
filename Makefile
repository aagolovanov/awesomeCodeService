run:
	go run ./cmd


build: clean
	go build -o build/runner ./...


test:
	go test -cover ./...


clean:
	rm -rf build