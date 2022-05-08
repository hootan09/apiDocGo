doc:
	godoc -http=:8080
build:
	go build .

.PHONY: doc build