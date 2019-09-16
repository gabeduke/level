FROM golang as build

# cache modules layer
WORKDIR /app
COPY go.mod go.sum ./
RUN go get -d -v ./...

# build app layer
COPY . .
RUN go build -v .

# executable layer
FROM scratch
COPY --from=build level /

# executable layer
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/level"]