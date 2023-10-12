.PHONY: build clean deploy

build:
	go mod tidy
	GO111MOULE=on
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/sendotp cmd/sendotp/main.go
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/verifyotp cmd/verifyotp/main.go

clean:
	rm -rf ./bin

deploy-dev: clean build
	sls deploy --stage development --verbose

deploy-prod: clean build
	sls deploy --stage production --verbose

deploy-local: build
	sls offline --useDocker
