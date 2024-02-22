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
	"strconv"
)

func (c *DropController) UploadPreviewImage(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user").(*models.User)
	if !ok {
		err := errors.New("No logged in user")
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	dropIdStr := r.PostFormValue("dropId")
	dropId, err := strconv.Atoi(dropIdStr)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusBadRequest)
		return
	}

	db := store.GetDB()
	m, err := models.FindDropById(db, dropId)
	if err != nil {
		helpers.LogError(err)
		if err == gorm.ErrRecordNotFound {
			err = errors.New(fmt.Sprintf("drop with id %d not found", dropId))
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

	m.PreviewImg = fileName
	err = models.UpdateDrop(db, m)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	c.WriteSuccessResponse(w, bytes, http.StatusOK)
}
