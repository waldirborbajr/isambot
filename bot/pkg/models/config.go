package models

type ISamConfig struct {
	RedisHost string `default:"127.0.0.1"`
	RedisPwd  string
	RedisDB   int `default:"0"`

	ISamToken   string `required:"true"`
  ISamHookURL string `required:"true"`
	ISamDebug   bool `default:"true"`
	ISamTLS     bool `default:"false"`
  ISamCert    string
}
