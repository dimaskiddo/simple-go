package configs

import (
	"os"

	"github.com/dimaskiddo/simple-go/api-mysql/helpers"
)

var CORSAllowedHeaders []string
var CORSAllowedOrigins []string
var CORSAllowedMethods []string

// Initialize Configuration Variable
func Initialize() {
	// Get Service IP from Environment Variable
	helpers.ServerConfig.IP = os.Getenv("SERVICE_IP")
	if len(helpers.ServerConfig.IP) == 0 {
		// If Service IP Environment Variable Not Exist
		// Then Set Service IP Variable to Default Value
		helpers.ServerConfig.IP = "0.0.0.0"
	}

	// Get Service Port from Environment Variable
	helpers.ServerConfig.Port = os.Getenv("SERVICE_PORT")
	if len(helpers.ServerConfig.Port) == 0 {
		// If Service Port Environment Variable Not Exist
		// Then Set Service Port Variable to Default Value
		helpers.ServerConfig.Port = "3000"
	}

	// Get CORS Allowed Headers from Environment Variable
	CORSAllowedHeaders = []string{os.Getenv("CORS_ALLOWED_HEADERS")}
	if len(CORSAllowedHeaders) == 0 {
		// If CORS Allowed Headers Environment Variable Not Exist
		// Then Set CORS Allowed Headers Variable to Default Value
		CORSAllowedHeaders = []string{"X-Requested-With"}
	}

	// Get CORS Allowed Origins from Environment Variable
	CORSAllowedOrigins = []string{os.Getenv("CORS_ALLOWED_ORIGINS")}
	if len(CORSAllowedOrigins) == 0 {
		// If CORS Allowed Origins Environment Variable Not Exist
		// Then Set CORS Allowed Origins Variable to Default Value
		CORSAllowedOrigins = []string{"*"}
	}

	// Get CORS Allowed Methods from Environment Variable
	CORSAllowedMethods = []string{os.Getenv("CORS_ALLOWED_METHODS")}
	if len(CORSAllowedMethods) == 0 {
		// If CORS Allowed Methods Environment Variable Not Exist
		// Then Set CORS Allowed Methods Variable to Default Value
		CORSAllowedMethods = []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	}

	// Get JWT Signing Key from Environment Variable
	helpers.JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")
	if len(helpers.JWTSigningKey) == 1 {
		// If JWT Signing Key Environment Variable Not Exist
		// Then Set JWT Signing Key Variable to Default Value
		helpers.JWTSigningKey = "signingkey"
	}

	// Get Database Host from Environment Variable
	helpers.MySQLConfig.Host = os.Getenv("DB_HOST")
	if len(helpers.MySQLConfig.Host) == 0 {
		// If Database Host Environment Variable Not Exist
		// Then Set Database Host Variable to Default Value
		helpers.MySQLConfig.Host = "127.0.0.1"
	}

	// Get Database Port from Environment Variable
	helpers.MySQLConfig.Port = os.Getenv("DB_PORT")
	if len(helpers.MySQLConfig.Port) == 0 {
		// If Database Port Environment Variable Not Exist
		// Then Set Database Port Variable to Default Value
		helpers.MySQLConfig.Port = "3306"
	}

	// Get Database User from Environment Variable
	helpers.MySQLConfig.User = os.Getenv("DB_USER")
	if len(helpers.MySQLConfig.User) == 0 {
		// If Database User Environment Variable Not Exist
		// Then Set Database User Variable to Default Value
		helpers.MySQLConfig.User = "user"
	}

	// Get Database Password from Environment Variable
	helpers.MySQLConfig.Password = os.Getenv("DB_PASSWORD")
	if len(helpers.MySQLConfig.Password) == 0 {
		// If Database Password Environment Variable Not Exist
		// Then Set Database Password Variable to Default Value
		helpers.MySQLConfig.Password = "password"
	}

	// Get Database Name from Environment Variable
	helpers.MySQLConfig.Name = os.Getenv("DB_NAME")
	if len(helpers.MySQLConfig.Name) == 0 {
		// If Database Name Environment Variable Not Exist
		// Then Set Database Name Variable to Default Value
		helpers.MySQLConfig.Name = "dbs"
	}
}
