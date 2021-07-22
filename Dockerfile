FROM golang:buster

# Import repository content
ADD ./src /app

# Set working directory
WORKDIR /app

# Make directory for data
RUN mkdir -p /app/Blockchain

# Install dependencies
RUN go get github.com/labstack/echo

# Build
RUN go build

# Run the program
CMD ["./blockchain"]