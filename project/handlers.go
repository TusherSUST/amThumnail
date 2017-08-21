package main

import (
    "fmt"
    "net/http"
  	"image/jpeg"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func GenerateThumbnail(w http.ResponseWriter, r *http.Request) {
    urls := r.URL.Query()["url"]

    images := GetImages(urls)
    hasImage := false
    for i,image := range(images) {
      if (image != nil) {
        hasImage = true
        jpeg.Encode(w, images[i], nil)
        break
      }
    }

    if !hasImage {
      fmt.Fprintln(w, "Not Available");
    }
}
