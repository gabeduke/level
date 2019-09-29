FROM golang:alpine as build

RUN apk --no-cache add ca-certificates

# cache modules layer
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# build app layer
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o level -ldflags "-X main.version=$(cat .release)"

# executable layer
FROM scratch
COPY --from=build /app/level level
COPY --from=build /etc/ssl/certs /etc/ssl/certs

# executable layer
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["./level"]