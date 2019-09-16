FROM golang as build

# cache modules layer
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# build app layer
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o level

# executable layer
FROM scratch
COPY --from=build /app/level level

# executable layer
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["./level"]