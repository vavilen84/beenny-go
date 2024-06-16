package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	awsClient "github.com/vavilen84/beenny-go/aws"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path/filepath"
)

func (c *UserController) UploadPreviewImage(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value("user").(*models.User)
	if !ok {
		err := errors.New("No logged in user")
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	db := store.GetDB()
	_, err := models.FindUserById(db, u.Id)
	if err != nil {
		helpers.LogError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("user with id %d not found", u.Id))
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExtension := filepath.Ext(r.FormValue("file"))
	s3Client, err := awsClient.GetS3Client()
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	bucketName := os.Getenv("S3_AWS_BUCKET")
	fileName := helpers.GenerateS3FilePath("preview-img", fileExtension)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	u.Photo = fileName
	err = models.SetUserPhoto(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	c.WriteSuccessResponse(w, bytes, http.StatusOK)
}
