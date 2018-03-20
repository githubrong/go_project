
export GOPATH=$(shell pwd)/:$(shell pwd)/vendor

RELEASE_TIME=$(shell date +%Y%m%d-%H%M)

all:iot_learn

codecheck:
    #find src -name *.go | xargs -n 1 go vet

iot_learn:codecheck   
	go build -o $(@F) '-X main.module_name=$(@F)' router
	mv $(@F) bin/$(@F)
clean:
	rm -rf release
