package main

import (
  "fmt"
  "image"
  "net/http"
  _ "image/gif"
  _ "image/jpeg"
  _ "image/png"
  "com_github_muesli_smartcrop"
  "time"
)

type HttpResponse struct {
  index    int
	url      string
	response *http.Response
	err      error
}

func GetImages (urls []string) []image.Image {
  ch := make(chan *HttpResponse, len(urls))

  tr := &http.Transport{
    MaxIdleConns:       10,
    IdleConnTimeout:    30 * time.Second,
    DisableCompression: true,
  }
  client := &http.Client{Transport: tr}

  for index, url := range urls {
		go func(url string, index int) {
      fmt.Printf("Fetching %d\n", index)
			resp, err := client.Get(url)
			ch <- &HttpResponse{index, url, resp, err}
		}(url, index)
	}

  successCount := 0
  images := make([]image.Image, len(urls))
  for {
		select {
		case r := <-ch:
      fmt.Printf("Fetched %d\n", r.index)
      successCount++
      images[r.index] = nil
      if r.err == nil && r.response.StatusCode == 200 {
        img, _, err := image.Decode(r.response.Body)
        r.response.Body.Close()
        if err == nil {
          images[r.index] = cropImage(img)
        }
      }
			if successCount == len(urls) {
				return images
			}
    }
	}

  return images;
}

func cropImage(img image.Image) image.Image {
  analyzer := smartcrop.NewAnalyzer()
  topCrop, _ := analyzer.FindBestCrop(img, 250, 250)
  type SubImager interface {
      SubImage(r image.Rectangle) image.Image
  }
  return img.(SubImager).SubImage(topCrop)
}
