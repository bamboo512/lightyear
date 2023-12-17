package schema

type StorageConfig struct {
	OriginalImagePath   string `yaml:"original_image_path"`
	ThumbnailImagePath  string `yaml:"thumbnail_image_path"`
	TranscodedImagePath string `yaml:"transcoded_image_path"`
}
