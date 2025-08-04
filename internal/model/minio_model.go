package model

type UploadFileInput struct {
	Bucket      string
	ObjectKey   string
	Content     []byte
	ContentType string
}

type PresignedURLInput struct {
	Bucket    string
	ObjectKey string
	Expiry    int64
}

type UploadFileRequest struct {
	Bucket string `form:"bucket" validate:"required"`
}

type DeleteFileRequest struct {
	Bucket    string `json:"bucket" validate:"required"`
	ObjectKey string `json:"object_key" validate:"required"`
}
