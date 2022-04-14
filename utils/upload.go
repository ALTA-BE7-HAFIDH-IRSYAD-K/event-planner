package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func UploadFile(fileParam *multipart.FileHeader, uploadFileDir string) (string, error) {
	var url string
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	awsRegion := os.Getenv("AWS_S3_REGION")
	awsBucket := os.Getenv("AWS_S3_BUCKET")
	awsKey := os.Getenv("AWS_S3_KEY")
	awsAccess := os.Getenv("AWS_S3_ACCESS")

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKey, awsAccess, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	upFile, err := fileParam.Open()
	if err != nil {
		return url, err
	}
	defer upFile.Close()

	var fileSize int64 = fileParam.Size
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	contentType := http.DetectContentType(fileBuffer)
	if contentType != "image/png" && contentType != "image/jpeg" {
		return url, errors.New("the file image not in jpg / png format")
	}
	uploadFileDir = strings.ReplaceAll(fmt.Sprint(uploadFileDir+"."+contentType[6:]), " ", "")

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(awsBucket),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	url = "https://eventabbi.s3.ap-southeast-1.amazonaws.com/" + uploadFileDir
	return url, err
}
