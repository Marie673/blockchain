FROM ubuntu:20.04

# Install Go
RUN apt update; apt install golang

# Import repository content
ADD ./src /app

# Set working directory: /app
WORKDIR /app

# Run the program
CMD ["go", "run", "*.go"]
