package helper

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

func ValidateSize(file *multipart.FileHeader, newFileName string) error {
	if file.Size > 500000 {
		return fmt.Errorf("ukuran file %s ini melebihi 500KB", newFileName)
	}
	return nil
}

func RetrieveImage(file *multipart.FileHeader, newFileName string) interface{} {

	read, err := file.Open()
	log.Println(file.Size)
	// if file.Size > 500000 {
	// 	return nil
	// }
	if err != nil {
		return nil
	}

	folderName := os.Getenv("OUTPUT_IMAGE")
	// folderName := "/home/drewjya/ayam/"
	currFileName := (strings.ReplaceAll(strings.ReplaceAll((folderName+newFileName), " ", ""), "-", ""))

	out, err := os.Create(currFileName)
	if err != nil {

		return nil
	}

	defer out.Close()

	_, err = io.Copy(out, read)
	if err != nil {

		return nil
	}
	return "/images/" + strings.ReplaceAll(strings.ReplaceAll(newFileName, " ", ""), "-", "")
}
