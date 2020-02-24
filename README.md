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


### TODO

1. Upload the images to another directory + show static files .
2. TLS. (https://github.com/denji/golang-tls  This example will be enough)
3. Crop better the image (accept all image formats).
4. Implement it in K8s.
