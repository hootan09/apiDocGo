# sudo apt install golang-golang-x-tools
doc:
	godoc -http=:8080 
test:
	go test -v
build:
	go build .
run:
	go run .

.PHONY: doc test build run