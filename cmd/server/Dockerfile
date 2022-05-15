FROM golang:1.18-bullseye AS build
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o /buildsys /app/cmd/server/main.go

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=build /buildsys /buildsys
EXPOSE 8080
CMD ["/buildsys"]
