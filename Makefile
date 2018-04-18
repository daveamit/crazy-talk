 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST= looper
GOGET=$(GOCMD) get -u -v
PROTOC	= protoc

.PHONY: all  
all: test build

.PHONY: build
build:
	@for dir in .proto/*; do \
		echo "Compiling $$(basename $$dir)"; \
		mkdir -p $$(basename $$dir); \
		${PROTOC} -I=$$dir --go_out=plugins=grpc:$$(basename $$dir)/ $$dir/*.proto; \
	done

.PHONY: test
test:
	$(GOTEST) --debug

.PHONY: clean
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: run
run: test
	# $(GOBUILD) -o $(BINARY_NAME) -v
	# ./$(BINARY_NAME)

.PHONY: deps
deps:
	$(GOGET) google.golang.org/grpc
	$(GOGET) github.com/golang/protobuf/protoc-gen-go
	$(GOGET) github.com/nathany/looper
	$(GOGET) github.com/stretchr/testify