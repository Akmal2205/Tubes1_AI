# Use official Golang image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire project
COPY . .

RUN chmod 777 run

# Keeps the container running in an interactive shell
ENTRYPOINT [ "tail", "-f", "/dev/null" ]