FROM golang:1.20

WORKDIR /usr/src/app

COPY *.go ./

COPY go.mod go.sum ./

COPY src/ ./src

# Install dependencies
RUN go mod download

# Build the app
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 5000

# Run the application when the container starts
CMD ["./main"]