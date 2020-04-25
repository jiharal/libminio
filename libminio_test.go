package libminio

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLibMinio(t *testing.T) {
	client := NewClient()
	client.Host = "-"
	client.AccessKey = "-"
	client.SecretKey = "-+"
	client.Bucket = ""
	client.Region = "-"
	client.SSL = true

	file, _ := os.Open("/Users/zzz/Downloads/prod.png")
	defer file.Close()

	fileImg, _, err := image.Decode(file)
	require.NoError(t, err)
	fileName := "content/square/marketplace/3bed03b8-76db-4526-8fb4-18f581dd958b.jpg"
	var newBuf bytes.Buffer

	err = jpeg.Encode(&newBuf, fileImg, nil)
	require.NoError(t, err)
	contentType := http.DetectContentType(newBuf.Bytes())
	t.Log(contentType)
	t.Log(fileName)
	url, err := client.Upload(fileName, newBuf.Bytes(), int64(newBuf.Len()), http.DetectContentType(newBuf.Bytes()))
	require.NoError(t, err)
	log.Println(url)
}
