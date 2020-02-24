# GO Crop HTTP Server + Redis 

## Description.

This is an Golang HTTP Server that is having 2 endpoints for:

1. Check the status (/status)
2. Receive a image, crop it and upload the url of the modified image to a redis db (/upload) .

## How to run it

First you need to download the repository:

```
git clone https://github.com/Chobito/golang-test.git
```

### Docker redis + go in local ( TestMode )

#### Run the Redis Docker
```
docker run --name redis-test-instance -p 6379:6379 -d redis
```

#### It will install dependencies for the Go

```
go get -d -v ./
```

#### Run the Go script

```
go run main.go
```

### Docker-compose ( "Prod" Mode )
#### Build and run the docker-compose
```bash
docker-compose build
docker-compose up
```

## How to test it.

1. Status endpoint.
  
  ```
  curl localhost:8080/status
  ```
It will return you the status of the server.

2. Upload image.

  ```bash
  curl -X POST --form "file=@maxdefault.jpg"   localhost:8080/upload
  ```
  
It will post a image, crop it and upload the url of the new image to a redis.


## Why I choose the libraries, the workflow, etc.

### Libraries
* log: Log on the server side the exceptions
* net/http: This library is used to manage all the HTTP requests + start the HTTP Server.
* go-redis/redis: This library is used on this program to connect to the redis server and upload the uri.
* gorilla/mux: This library is to router all the Endpoints easily.
* os: This library is needed to launch OS methods in this case open the file.
* io: Library to access to I/O primitives (printing output of the upload and copying) . 
* desintegration/imaging: This library is what I used to crop the image
* image + image/color: This libraries is needed to use desintegration cropping tool.
* math/rand: I used it to generate random numbers
* strconv: This library I used for convert numbers in text to concatenate them. 

### Workflow

1. I write the complete program on a paper.
2. I started managing only the Endpoints + the upload image, after that the cropping was easier.
3. I spent more time than I expected trying to manage the HTTP header and seems that this libraries doesn't expect any header.
4. After that, I mounted the redis docker in order to try simple communication through the Go program and the redis server.
5. Before dockerize the app , I cleaned a bit the code (I am not an Go expert)
6. I dockerized the Go app and managed it with docker-compose (I wanted to put it on k8s but it's crazy for this simple app.
7. I enjoyed a lot learning Go.

### TODO

1. Upload the images to another directory + show static files .
2. TLS. (https://github.com/denji/golang-tls  This example will be enough)
3. Crop better the image (accept all image formats).
4. Implement it in K8s.
