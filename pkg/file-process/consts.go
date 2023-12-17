package fileprocess

const (
	DefaultQuality       = 80
	ParamErrorQuality    = "parameter error: quality can only be between 1-100"
	ParamErrorEncoding   = "parameter error: encoding can only be avif, heic or web"
	InternalServerErr    = "Internal Server Error"
	SupportedEncodings   = "avif, heic, webp"
	OriginalImagePathEnv = "ORIGINAL_IMAGE_PATH"
	EncodedImagePathEnv  = "ENCODED_IMAGE_PATH"
)
