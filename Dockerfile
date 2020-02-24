# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git 

# Copy the source from the current directory to the Working Directory inside the container
COPY main.go .
COPY Dockerfile .
COPY docker-compose.yml .

# Build the Go app
RUN go get -d -v ./
# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT [ "go","run","main.go" ] 
