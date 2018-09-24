GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
NAME=computor
NAME_UNIX=$(NAME)_unix

all : build

build:
	$(GOBUILD) -o $(NAME) -v

test: 
		$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(NAME)
	rm -f $(NAME_UNIX)

run: 
	$(GOBUILD) -o $(NAME) -v ./...
	./$(NAME)

