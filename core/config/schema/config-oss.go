package schema

type OssConfig struct {
	Enabled         bool   `yaml:"enabled"`
	AccountId       string `yaml:"account_id"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
}
