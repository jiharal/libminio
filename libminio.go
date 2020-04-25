package libminio

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v6"
)

// Client is a struct
type Client struct {
	Host      string
	AccessKey string
	SecretKey string
	Bucket    string
	Region    string
	SSL       bool
	LogLevel  int
	Logger    *log.Logger
}

// NewClient is used to initial client...
func NewClient() Client {
	return Client{
		LogLevel: 2,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

// Call is used to initial minio client
func (opt *Client) Call() *minio.Client {
	logger := opt.Logger
	err := StringValidate(opt.Host, opt.AccessKey, opt.SecretKey, opt.Region, opt.Bucket)
	if err != nil {
		logger.Fatalf("Initial %v", err.Error())
	}
	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.NewWithRegion(opt.Host, opt.AccessKey, opt.SecretKey, opt.SSL, opt.Region)
	if err != nil {
		logger.Fatalf("Could not initialize S3 %v", err.Error())

	}
	return client
}

// Upload is used to ...
func (opt *Client) Upload(tempFileName string, buffer []byte, size int64, contentType string) (string, error) {
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err := opt.Call().PutObject(
		opt.Bucket,
		strings.TrimPrefix(tempFileName, "/"),
		bytes.NewReader(buffer),
		size,
		minio.PutObjectOptions{
			ContentType:  contentType,
			UserMetadata: userMetaData,
		})
	if err != nil {
		return "", err
	}

	return tempFileName, nil
}
