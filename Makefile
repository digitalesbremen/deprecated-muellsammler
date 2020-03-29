# Go parameters

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.
# -o gpio-test-x  -> force to build an executable gpio-test-x file (instead of default https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

BINARY_NAME=muellsammler
BINARY_AMD64=$(BINARY_NAME)_amd64
BINARY_ARM=$(BINARY_NAME)_arm
RASPI_SCP_FOLDER=pi@pi4-rack-0.local:/home/pi/test

all: deps clean test build
build:
	go build -a -o $(BINARY_NAME) -v
test:
	go test -v ./...
	rm -f test.json
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_AMD64)
	rm -f $(BINARY_ARM)
	rm -f test.json
deps:
	go get

# Cross compilation
build-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $(BINARY_AMD64) -v
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -o $(BINARY_ARM) -v

# Deploy
copy-raspberry:
	scp $(BINARY_ARM) $(RASPI_SCP_FOLDER)

