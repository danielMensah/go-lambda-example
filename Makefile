.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/customer/get_customers/api api/customer/get_customers/*.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/order/get_orders/api api/order/get_orders/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --stage=$(stage) --verbose

deploy-function: clean build
	serverless deploy function --stage=$(stage) -f $(function) --verbose # --aws-s3-accelerate can be used to speed it up