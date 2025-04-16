package request

type FileUploadRequest struct {
	FileName   string `json:"file_name"`
	IsUploaded bool   `json:"is_uploaded"`
}

type FileProcessRequest struct {
	FileName    string `json:"file_name"`
	IsProcessed bool   `json:"is_processed"`
}