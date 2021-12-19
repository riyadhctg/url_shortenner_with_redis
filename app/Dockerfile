FROM golang:1.12-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/go-url-shortenner

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-url-shortenner .


# This container exposes port 4000 to the outside world
EXPOSE 4000

# Run the binary program produced by `go install`
CMD ["./out/go-url-shortenner"]