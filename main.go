//usage:
// Status:     curl localhost:8080/status
// Upload:     curl -i -X POST -F file=@Felis_silvestris_silvestris_small_gradual_decrease_of_quality.png   localhost:8080/upload
//accepting files: jpeg, png ...    JPEG is failing (I suppose per encoding)
package main

import (
	"github.com/disintegration/imaging"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"image"
	"image/color"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var count int = 0
var count_to_redis string = ""

// handler for status (GET)
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Status OK"}`))
}

//handler for the post (not needed)
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

//handler for the notFound
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

// function to upload + resize file + upload key to redis
func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, err := imaging.Open(handler.Filename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Crop the original image to 300x300px size using the center anchor.
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	img1 := imaging.Invert(src)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(200, 400, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))

	//generating random file name for the image
	var str1 int
	str1 = rand.Int()

	var str2 string
	str2 = "_modified.jpg"
	concatenated := strconv.Itoa(str1) + str2

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, concatenated)

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	//using a counter for the redis key.
	count = count + 1
	count_to_redis = strconv.Itoa(count)

	//TODO: The images are not accesible from outside, I tried serving static files but it's not working
	_, _ = io.WriteString(w, "File Uploaded localhost:8080/"+concatenated+" Stored in Redis with key: "+count_to_redis+"\n")
	_, _ = io.Copy(f, file)
	//connection to redis to upload the key.
	client := redis.NewClient(&redis.Options{Addr: "redis:6379", DB: 0})
	err = client.Set(count_to_redis, concatenated, 0).Err()
}

//Main func + server.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/status", get).Methods(http.MethodGet)
	r.HandleFunc("/upload", UploadFile).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
