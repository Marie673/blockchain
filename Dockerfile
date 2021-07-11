FROM golang:buster
ADD ./app /app

# Set working directory: /app
WORKDIR /app

# Run the program
CMD ["go", "run", "main.go"]
