TARGET_NAME=windspiel

all: release

clean:
	rm -rf ./build/

build-jobs: build-x86-linux build-x64-linux build-armv7-linux build-x64-windows build-x86-windows build-darwin

prepare:
	mkdir -p ./build
	cp config.yml ./build/


build-x86-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o build/$(TARGET_NAME)-x86-linux 

build-x64-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/$(TARGET_NAME)-x64-linux 

# raspberry pi
build-armv7-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-s -w" -o build/$(TARGET_NAME)-armv7-linux 

build-x64-windows: 
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/$(TARGET_NAME)-x64-windows.exe 

build-x86-windows: 
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o build/$(TARGET_NAME)-x86-windows.exe 

build-darwin: 
	CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s -w" -o build/$(TARGET_NAME)-macos 

release: clean prepare build-jobs