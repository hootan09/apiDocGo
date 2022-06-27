# sudo apt install golang-golang-x-tools
doc:
	godoc -http=:8080 
test:
	go test -v
build:
	go build .
run:
	go run .
help:
	go run . -h
serve:
	go run . -p 8080

.PHONY: doc test build run