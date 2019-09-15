FROM golang

# cache modules layer
WORKDIR /app
COPY go.mod go.sum ./
RUN go get -d -v ./...

# build app layer
COPY . .
RUN go install -v ./...

# executable layer
ENV PORT 8080
EXPOSE 8080
CMD ["level-api"]