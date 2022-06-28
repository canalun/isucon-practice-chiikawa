REVISION=$(shell git rev-parse --short HEAD)

build:
	GOARCH="amd64" GOOS="linux" go build -trimpath -ldflags "-X main.revision=$(REVISION)"

mac-build:
	go build -v -trimpath -ldflags "-X main.revision=$(REVISION)"

deploy: build
	ssh isucon@13.112.9.14 sudo systemctl stop isucholar.go.service
	scp -P 22 isucholar isucon@13.112.9.14:/home/isucon/webapp/go
	ssh isucon@13.112.9.14 sudo systemctl start isucholar.go.service
	ssh isucon@13.112.9.14 sudo systemctl enable isucholar.go.service
