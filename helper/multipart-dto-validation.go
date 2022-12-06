package helper

import (
	"mime/multipart"
	"reflect"

	"github.com/go-playground/validator"
)

func MultipartDTOValidation(data interface{}, validate validator.StructLevel) {
	val := reflect.ValueOf(data)
	for i := 0; i < val.NumField(); i++ {
		switch val.Field(i).Interface().(type) {
		case *multipart.FileHeader:
			datas := val.Field(i).Interface().(*multipart.FileHeader)
			errExt := ValidateExtension(datas.Filename)
			if !errExt {
				validate.ReportError("format file dari "+val.Type().Field(i).Name+"tidak di dukung", "FotoFreezerBawah", "FotoFreezerBawah", "FotoFreezerBawah", "")
			}
			errSize := ValidateSize(datas, val.Type().Field(i).Name)
			if errSize != nil {
				validate.ReportError("ukuran dari "+val.Type().Field(i).Name+" melebih 500KB", "FotoSelfie", "FotoSelfie", "FotoSelfie", "")
			}
		}

	}
}
