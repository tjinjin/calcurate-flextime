NAME := calcurate-flextime
GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)

.PHONY: build

build:
	glide install

release: clean build
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/$(NAME)
	(cd dist;zip -m $(NAME)-linux-amd64.zip $(NAME))
	GOOS=darwin GOARCH=amd64 go build -o dist/$(NAME)
	(cd dist;zip -m $(NAME)-darwin-amd64.zip $(NAME))

install: build
	go build

clean:
	rm -f dist/$(NAME)*
