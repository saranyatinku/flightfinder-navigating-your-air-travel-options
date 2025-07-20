APISERVER = ./cmd/finder_web/app/apiserver/go/routers.go
WEBSERVER = ./server

default: ${APISERVER}

docker: 
	docker build . -t flight-finder:latest
	
rundocker: 
	docker run --rm --name=flight-finder -p 8080:80 -v $HOME/.aws:/root/.aws flight-finder:latest
	
pushdocker:
	docker login -u mateuszmidor -p <pass>
	docker build . -t mateuszmidor/flight-finder:latest
	docker push mateuszmidor/flight-finder:latest

runcli: ${APISERVER}
	go run cmd/finder_cli/main.go -flights_data=./assets

runweb: ${APISERVER}
	GIN_MODE=release  go run cmd/finder_web/main.go -port=8080 -flights_data=./assets -web_data=./web -aws_region=us-east-1 -aws_xray=false -redis_addr=localhost:6379 -redis_pass=CACHE

buildweb: ${WEBSERVER}
${WEBSERVER}: ${APISERVER}
	@echo "Building web server"
	CGO_ENABLED=0 go build -o ${WEBSERVER} cmd/finder_web/main.go

test: ${APISERVER}
	go vet ./...
	go test ./...

# Generate Go server stub from OpenAPI3 specification
# see: https://github.com/OpenAPITools/openapi-generator
# see: https://openapi-generator.tech/docs/generators/go-gin-server
# see: https://openapi-generator.tech/docs/generators/go
${APISERVER}: ./api/openapi3.yaml
	@echo "Building OpenAPI3 server"
	docker run \
		--rm \
		-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
		-v "$(shell pwd):/build" \
		--entrypoint=/bin/bash \
		openapitools/openapi-generator-cli:latest -c "cd /build && scripts/build_openapi3.sh"

.PHONY: default docker rundocker runcli runweb  buildweb test