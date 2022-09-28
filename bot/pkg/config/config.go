package config

import (
	"localhost/isambot/bot/pkg/models"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// ReadConfig from config.yaml
func ReadConfig() *models.ISamConfig {

	config := &models.ISamConfig{}

	if err := godotenv.Load(".env_isam"); err != nil {
		log.Fatal().Err(err).Msg("iSam Config missing or incorrect.")
	}

	// Read Redis configs
	db, _ := strconv.Atoi(os.Getenv("redis_db"))
	config.RedisDB = db
	config.RedisHost = os.Getenv("redis_host")
	config.RedisPwd = os.Getenv("redis_pwd")

	// Read iSamBot configs
	debug, _ := strconv.ParseBool(os.Getenv("isam_debug"))
	tls, _ := strconv.ParseBool(os.Getenv("isam_tls"))
	config.ISamDebug = debug
  config.ISamTLS = tls
	config.ISamHookURL = os.Getenv("isam_hookurl")
	config.ISamToken = os.Getenv("isam_token")
  config.ISamCert = os.Getenv("isam_cert")

	return config
}

// var (
// 	ChatId         string = os.Getenv("CHAT_ID")
// 	ApiKey         string = os.Getenv("API_KEY")
// 	ApiUrl         string = "https://api.telegram.org"
// 	BotAddress     string = os.Getenv("BOT_ADDRESS")
// 	BotBindAddress string = os.Getenv("BOT_BIND_ADDRESS")
// 	BotId          string = os.Getenv("BOT_ID")
// 	BotKey         string = os.Getenv("BOT_KEY")
// 	BotPort        string = os.Getenv("BOT_PORT")
// 	CertPath       string = os.Getenv("BOT_CERT_PATH")
// 	KeyPath        string = os.Getenv("BOT_KEY_PATH")
// 	LogPath        string = os.Getenv("LOG_PATH")
// 	RedisAddress   string = os.Getenv("REDIS_ADDRESS")
// 	MongoAddress   string = os.Getenv("MONGO_ADDRESS")
// 	MongoUser      string = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
// 	MongoPassword  string = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
// 	MongoDatabase  string = os.Getenv("MONGO_DATABASE")
// )
//
