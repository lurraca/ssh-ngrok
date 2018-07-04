BINARY_NAME=ssh-ngrok

all: deps test build

build:
	go build -o $(BINARY_NAME) -v

test:
	ginkgo -p -r

clean:
	go clean
	rm -rf $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega/...
	go get github.com/ogier/pflag
