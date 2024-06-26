package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine/v2"
	"google.golang.org/appengine/v2/blobstore"
	"google.golang.org/appengine/v2/image"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const jsonMimeType = "application/json"

type jsonResponse struct {
	ImageURL string `json:"image_url,omitempty"`
	Error    string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response(
			w,
			jsonResponse{Error: "Invalid request. Use `/image-url` endpoint."},
			http.StatusBadRequest,
		)
	})
	http.HandleFunc("/image-url", getImageHandler)

	appengine.Main()
}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	bucketName := r.URL.Query().Get("bucket")
	imagePath := r.URL.Query().Get("image")

	if bucketName == "" {
		response(
			w,
			jsonResponse{Error: "Missing `bucket` parameter."},
			http.StatusUnprocessableEntity,
		)

		return
	}

	if imagePath == "" {
		response(
			w,
			jsonResponse{Error: "Missing `image` parameter."},
			http.StatusUnprocessableEntity,
		)

		return
	}

	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		response(
			w,
			jsonResponse{Error: "Error creating storage client."},
			http.StatusInternalServerError,
		)

		return
	}
	defer client.Close()

	blobKey, err := blobstore.BlobKeyForFile(ctx, fmt.Sprintf("/gs/%s/%s", bucketName, imagePath))
	if err != nil {
		response(
			w,
			jsonResponse{Error: fmt.Sprintf("Error getting blob key: %v", err)},
			http.StatusInternalServerError,
		)

		return
	}

	imageServingUrl, err := image.ServingURL(ctx, blobKey, nil)
	if err != nil {
		response(
			w,
			jsonResponse{Error: fmt.Sprintf("Error getting serving URL: %v", err)},
			http.StatusInternalServerError,
		)

		return
	}

	response(
		w,
		jsonResponse{ImageURL: imageServingUrl.String()},
		http.StatusOK,
	)
}

func response(w http.ResponseWriter, data jsonResponse, status int) {
	w.Header().Set("Content-Type", jsonMimeType)
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
