build: 
	export GO111MODULE=on
	go build  -o bin/parseTickers parseTickers/main.go
	go build  -o bin/watchReddit watchReddit/main.go

watchReddit: build
	./bin/watchReddit