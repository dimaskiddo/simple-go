package configs

import (
	"os"

	"github.com/dimaskiddo/simple-go/api-simple/helpers"
)

var SvcIP string
var SvcPort string

var CORSAllowedHeaders []string
var CORSAllowedOrigins []string
var CORSAllowedMethods []string

// Initialize Configuration Variable
func Initialize() {
	// Get Service IP from Environment Variable
	SvcIP = os.Getenv("SERVICE_IP")
	if len(SvcIP) == 0 {
		// If Service IP Environment Variable Not Exist
		// Then Set Service IP Variable to Default Value
		SvcIP = "0.0.0.0"
	}

	// Get Service Port from Environment Variable
	SvcPort = os.Getenv("SERVICE_PORT")
	if len(SvcPort) == 0 {
		// If Service Port Environment Variable Not Exist
		// Then Set Service Port Variable to Default Value
		SvcPort = "3000"
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
}
