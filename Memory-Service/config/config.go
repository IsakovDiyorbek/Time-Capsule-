package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	DefaultOffset string
	DefaultLimit  string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found", err)
	}

	config := Config{}

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7777"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "memory"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "20005"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "memory_exam"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}




// SELECT 
//     borrower_id, 
//     user_id, 
//     book_id, 
//     borrow_date, 
//     return_date, 
//     created_at, 
//     updated_at, 
//     deleted_at
// FROM 
//     borrow_records
// WHERE 
//     return_date IS NULL
//     AND borrow_date < NOW() - INTERVAL '30 days' -- Assuming a 30-day borrowing period
//     AND deleted_at IS NULL;
