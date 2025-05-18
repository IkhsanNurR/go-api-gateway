package controllers

import (
	"api-gateway/response"
	"api-gateway/utils"
	"context"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	FileService *utils.FileService
}

type UploadRequest struct {
	Path string `form:"path" binding:"required" ` // wajib ada
}

func NewFileHandler(fs *utils.FileService) *FileHandler {
	return &FileHandler{
		FileService: fs,
	}
}

// Upload handler: menerima file dari form dan upload ke MinIO
func (h *FileHandler) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	objectPath := c.PostForm("path")
	if err := c.ShouldBind(&UploadRequest{Path: objectPath}); err != nil {
		errorsMap := utils.ParseValidationErrors(err)

		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse(
			errorsMap,
			"validation error",
			http.StatusBadRequest,
			"validation error on input fields",
		))
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse(
			nil,
			"file not provided",
			http.StatusBadRequest,
			err.Error(),
		))
		return
	}

	if objectPath == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse(
			nil,
			"path not provided",
			http.StatusBadRequest,
			nil,
		))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedResponse(
			nil,
			"failed to open file",
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}
	defer file.Close()

	contentType := fileHeader.Header.Get("Content-Type")
	fileSize := fileHeader.Size
	objectName := fileHeader.Filename

	url, err := h.FileService.UploadFile(context.Background(), path.Join(objectPath, objectName), file, fileSize, contentType)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse(
			nil,
			"upload failed",
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, response.SuccessOneResponse(gin.H{"url": url}, "file uploaded", http.StatusOK))
}

// Download handler: mengunduh file dari MinIO
func (h *FileHandler) Download(c *gin.Context) {
	objectName := c.Query("filename")
	if objectName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse(
			nil,
			"filename not provided",
			http.StatusBadRequest,
			"filename not provided",
		))
		return
	}
	reader, err := h.FileService.DownloadFile(context.Background(), objectName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, response.FailedResponse(
			nil,
			"file not found",
			http.StatusNotFound,
			err.Error(),
		))
		return
	}
	defer reader.Close()

	// Detect content type from file extension
	ext := filepath.Ext(objectName) // contoh: ".jpg"
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream" // fallback default
	}

	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(objectName))
	c.Header("Content-Type", contentType)
	if _, err := io.Copy(c.Writer, reader); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, response.FailedResponse(
			nil,
			"file not found",
			http.StatusNotFound,
			err.Error(),
		))
		return
	}
}

// Delete handler: menghapus file dari MinIO
func (h *FileHandler) Delete(c *gin.Context) {
	objectName := c.Query("filename")
	if objectName == "" {
		c.JSON(http.StatusBadRequest, response.FailedResponse(
			nil,
			"filename not provided",
			http.StatusBadRequest,
			"filename not provided",
		))
		return
	}
	err := h.FileService.DeleteFile(context.Background(), objectName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse(
			nil,
			"delete failed",
			http.StatusInternalServerError,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, response.SuccessOneResponse(nil, "file deleted", http.StatusOK))
}
