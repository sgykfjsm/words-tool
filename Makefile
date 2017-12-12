APP_NAME := words-tool
VERSION :=
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

init:
	dep init

deps:
	dep ensure
	dep ensure --update

build: clean
	mkdir -pv bin
	go build $(LDFLAGS) -o bin/$(APP_NAME) main.go

clean:
	$(RM) -rv bin

run: build
	$(shell cat env.sh) ./bin/$(APP_NAME) $(ARG)
