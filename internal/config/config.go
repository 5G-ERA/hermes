package config

type HermesConfig struct {
	AwsAccessKeyId     string
	AwsAccessSecretKey string
	AwsRegion          string
	AwsBucket          string
	FetchDir           string
	PostDir            string
}

type HermesOption func(*HermesConfig)

func WithRegion(region string) HermesOption {
	return func(c *HermesConfig) {
		c.AwsRegion = region
	}
}

func WithBucket(bucket string) HermesOption {
	return func(c *HermesConfig) {
		c.AwsBucket = bucket
	}
}
func WithFetchDir(fetchDir string) HermesOption {
	return func(c *HermesConfig) {
		c.FetchDir = fetchDir
	}
}
func WithPostDir(postDir string) HermesOption {
	return func(c *HermesConfig) {
		c.PostDir = postDir
	}
}

func NewConfig(awsAccessKeyId, awsAccessSecretKey string, opts ...HermesOption) *HermesConfig {
	if awsAccessKeyId == "" || awsAccessSecretKey == "" {
		panic("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}
	cfg := &HermesConfig{
		AwsAccessKeyId:     awsAccessKeyId,
		AwsAccessSecretKey: awsAccessSecretKey,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
