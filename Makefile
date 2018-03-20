
export GOPATH=$(shell pwd)/:$(shell pwd)/vendor

RELEASE_TIME=$(shell date +%Y%m%d-%H%M)

all:iot_learn

codecheck:
    #find src -name *.go | xargs -n 1 go vet

iot_learn:codecheck   
	go build -o $(@F)  -ldflags  '-extldflags "-static" -X main.module_name=$(@F)' router
clean:
	rm -rf release
