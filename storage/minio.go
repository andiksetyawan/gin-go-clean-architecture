package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

type minIO struct {
	MinioClient *minio.Client
}

func NewMinIO() Storage {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY_ID")
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}
	return &minIO{MinioClient: client}
}

func (m minIO) Upload(fileID string, f io.Reader, bucket string) error {
	//buff := bytes.NewBuffer([]byte{})
	//size, err := io.Copy(buff, f)
	//if err != nil {
	//	return err
	//}

	res, err := m.MinioClient.PutObject(context.TODO(), bucket, fileID, f, -1, minio.PutObjectOptions{})

	if err != nil {
		log.Println("minio error", err)
		return err
	}
	log.Println(res)
	return nil
}
