package configs

import (
	"os"
)

var SvcPort string

// Initialize Configuration Variable
func Initialize() {
	// Get Service Port from Environment Variable
	SvcPort = ":" + os.Getenv("SERVICE_PORT")
	if len(SvcPort) == 1 {
		// If Service Port Environment Variable Not Exist
		// Then Set Service Port Variable to Default Value
		SvcPort = ":3000"
	}
}
