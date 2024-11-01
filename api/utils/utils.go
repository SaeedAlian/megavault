package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ParseJSONFromRequest(r *http.Request, payload any) error {
	body := r.Body

	if body == nil {
		return fmt.Errorf("Request body is not found")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSONInResponse(
	w http.ResponseWriter,
	status int,
	payload any,
	headers *map[string]string,
) error {
	if headers != nil {
		for k, v := range *headers {
			w.Header().Add(k, v)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

func WriteErrorInResponse(w http.ResponseWriter, status int, message string) error {
	return WriteJSONInResponse(w, status, map[string]string{"message": message}, nil)
}

func CreateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			return r
		}
		return -1
	}, slug)

	return slug
}

func FileUploadHandler(
	field string,
	maxSizeInMB int64,
	mimeTypes []string,
	directory string,
) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		maxSizeInBytes := maxSizeInMB * 1024 * 1024

		r.Body = http.MaxBytesReader(w, r.Body, maxSizeInBytes)
		if err := r.ParseMultipartForm(maxSizeInBytes); err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf(
					"The uploaded file is too big. Please choose an file that's less than %dMB in size",
					maxSizeInMB,
				),
			)
			return
		}

		file, handler, err := r.FormFile(field)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Cannot retrieve the file: %v", err),
			)
			return
		}
		defer file.Close()

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error in file uploading: %v", err),
			)
			return
		}

		mimeType := mimetype.Detect(buf)

		for i := range mimeTypes {
			m := mimeTypes[i]

			if mimeType.String() != m {
				allowedMimeTypesString := strings.Join(mimeTypes, " , ")
				WriteErrorInResponse(
					w,
					http.StatusBadRequest,
					fmt.Sprintf(
						"Cannot upload %s file, please upload only %s files",
						mimeType,
						allowedMimeTypesString,
					),
				)
				return
			}
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error in file uploading: %v", err),
			)
			return
		}

		err = os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error in file uploading: %v", err),
			)
			return
		}

		filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(handler.Filename))
		fullpath := fmt.Sprintf("%s/%s", directory, filename)

		dest, err := os.Create(fullpath)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error in file uploading: %v", err),
			)
			return
		}
		defer dest.Close()

		_, err = io.Copy(dest, file)
		if err != nil {
			WriteErrorInResponse(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error in file uploading: %v", err),
			)
			return
		}

		WriteJSONInResponse(w, http.StatusOK, map[string]string{
			"message": fmt.Sprintf(
				"File with the name '%s' has been uploaded successfully",
				filename,
			),
		}, nil)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
