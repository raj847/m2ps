package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver = GetEnv("DB_DRIVER", "postgres")
	DBName   = GetEnv("DB_NAME", "local")
	DBHost   = GetEnv("DB_HOST", "localhost")
	DBPort   = GetEnv("DB_PORT", "5432")
	DBUser   = GetEnv("DB_USER", "root")
	DBPass   = GetEnv("DB_PASS", "")
	SSLMode  = GetEnv("SSL_MODE", "disable")

	SALT_KEY = GetEnv("SALT_KEY")

	REDISHost = GetEnv("REDIS_HOST")
	REDISPass = GetEnv("REDIS_PASS")
	REDISPort = GetEnv("REDIS_PORT")

	REDISHostLocal = GetEnv("REDIS_HOST_LOCAL")
	REDISPassLocal = GetEnv("REDIS_PASS_LOCAL")
	REDISPortLocal = GetEnv("REDIS_PORT_LOCAL")

	MONGOHost = GetEnv("MONGO_HOST")
	MONGOPort = GetEnv("MONGO_PORT")
	MONGODB   = GetEnv("MONGO_DB")

	MONGOHostPg = GetEnv("MONGO_HOST_PG")
	MONGOPortPg = GetEnv("MONGO_PORT_PG")
	MONGODBPg   = GetEnv("MONGO_DB_PG")

	AMQPServerUrl   = GetEnv("AMQP_SERVER_URL")
	DYNAMOServerUrl = GetEnv("DYNAMO_SERVER_URL")
	DYNAMORegion    = GetEnv("DYNAMO_REGION")
	DYNAMOProfile   = GetEnv("DYNAMO_PROFILE")

	QUEUEName           = GetEnv("QUEUE_NAME")
	DURATIONCallbackMls = GetEnv("DURATION_CALLBACK_MLS")

	BASEApiUrl         = GetEnv("BASE_API_URL")
	SUFFIXApiUrl       = GetEnv("SUFFIX_API_TRX")
	CALLBACKPaymentUrl = GetEnv("CALLBACK_PAYMENT_URL")
	BASEPaymentUrl     = GetEnv("BASE_PAYMENT_URL")
	RSYNCConsumerType  = GetEnv("RSYNC_CONSUMER_TYPE")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}
}
