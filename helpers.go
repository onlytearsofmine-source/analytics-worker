package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AssetID returns a short id for the provided asset
func AssetID(asset string) string {
	return fmt.Sprintf("%x", Hash(asset))
}

// Hash returns a SHA-256 hash of the provided string
func Hash(str string) []byte {
	h := sha256.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

// CreateS3Session creates a new AWS S3 session
func CreateS3Session() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return s3.New(sess)
}

// CreateMongoSession creates a new MongoDB session
func CreateMongoSession() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(os.Getenv("MONGO_NAME"))
}

// ProcessFiles processes a list of files in a goroutine
func ProcessFiles(wg *sync.WaitGroup, ch chan *fileData, done chan struct{}) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for fd := range ch {
			processFile(fd)
		}
		close(done)
	}()
}

// Filename returns the filename from the provided file path
func Filename(file string) string {
	return filepath.Base(file)
}

// Extension returns the file extension from the provided file name
func Extension(file string) string {
	return filepath.Ext(file)
}

// GetFileMetadata returns the metadata of the file
func GetFileMetadata(file string) (ftype string, size int64, err error) {
	f, err := os.Open(file)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		return "", 0, err
	}
	return "", st.Size(), nil
}

// processFile processes a single file
func processFile(fd *fileData) {
	// Implementation of processFile
}

type fileData struct {
	Path   string
	AssetID string
	Upload bool
}