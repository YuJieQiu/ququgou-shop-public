package resource

import "errors"

var (
	//
	ErrFileEmpty = errors.New("file empty")
	//
	ErrFileExceedSize = errors.New("file exceed size")
	//
	ErrFileTypeError = errors.New("file type error")
	//
	ErrByteNullError = errors.New("file byte error")

	ErrInvalidStorageType = errors.New("invalid storage type")
)

type UploadError struct {
	Msg string
	Err error
}

func (e *UploadError) Error() string {
	if e.Msg == "" {
		e.Msg = "文件上传错误"
	}
	return e.Msg + ":" + e.Err.Error()
}

func (e *UploadError) Unwrap() error {
	return e.Err
}
