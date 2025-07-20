# Generate API code from OpenAPI spec
FROM openapitools/openapi-generator-cli:latest as openapi_builder
WORKDIR /builder
COPY ./api/ ./api/
COPY ./cmd/ ./cmd/
COPY ./scripts/ ./scripts/
RUN ["/bin/bash", "scripts/build_openapi3.sh"]

# Build server binary
FROM golang as builder
WORKDIR /builder
COPY ./go.mod ./go.mod 
COPY ./go.sum ./go.sum 
COPY ./api/ ./api/
COPY ./pkg/ ./pkg/
COPY --from=openapi_builder /builder/cmd/ ./cmd/
RUN CGO_ENABLED=0 go build -o server cmd/finder_web/main.go 

# Build final docker image
FROM alpine
WORKDIR /opt/
COPY ./assets/*.gz ./assets/
COPY ./web/ ./web/ 
EXPOSE 80
COPY --from=builder /builder/server .
ENTRYPOINT ["/opt/server", "-port=80", "-flights_data=./assets", "-web_data=./web", "-aws_region=us-east-1", "-aws_xray=false", "-redis_addr=localhost:6379", "-redis_pass=CACHE"]
