# syntax=docker/dockerfile:1

##
## BUILD
##

FROM golang:1.19-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /ping-service

##
## TEST
##

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

##
## DEPLOY
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /ping-service /ping-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ping-service"]

