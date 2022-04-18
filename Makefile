doc:
	godoc -hhtp=:8080
build:
	go build .

.PHONY: doc build