FROM golang:buster

# Import repository content
ADD ./src /app

# Set working directory
WORKDIR /app

RUN go get github.com/labstack/echo

# Run the program
CMD ["go", "run", "./"]
