REVISION=$(shell git rev-parse --short HEAD)

build:
	GOARCH="amd64" GOOS="linux" go build -trimpath -ldflags "-X main.revision=$(REVISION)"  -o isucholar

mac-build:
	go build -v -trimpath -ldflags "-X main.revision=$(REVISION)" -o isucholar

deploy: build
	ssh ubuntu@13.112.9.14 -i ~/.ssh/aws-tocomi.pem sudo systemctl stop isucholar.go.service
	scp -i ~/.ssh/aws-tocomi.pem -P 22 isucholar ubuntu@13.112.9.14:/home/ubuntu
	ssh ubuntu@13.112.9.14 -i  ~/.ssh/aws-tocomi.pem sudo mv /home/ubuntu/isucholar /home/isucon/webapp/go/
	ssh ubuntu@13.112.9.14 -i  ~/.ssh/aws-tocomi.pem sudo systemctl start isucholar.go.service
	ssh ubuntu@13.112.9.14 -i  ~/.ssh/aws-tocomi.pem sudo systemctl enable isucholar.go.service
