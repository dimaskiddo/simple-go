package configs

import (
	"os"
)

var SvcPort string
var DBConfig databases

// Initialize Configuration Variable
func Initialize() {
	// Get Service Port from Environment Variable
	SvcPort = ":" + os.Getenv("SERVICE_PORT")
	if len(SvcPort) == 1 {
		// If Service Port Environment Variable Not Exist
		// Then Set Service Port Variable to Default Value
		SvcPort = ":3000"
	}

	// Get Database Host from Environment Variable
	DBConfig.Host = os.Getenv("DB_HOST")
	if len(DBConfig.Host) == 0 {
		// If Database Host Environment Variable Not Exist
		// Then Set Database Host Variable to Default Value
		DBConfig.Host = "127.0.0.1"
	}

	// Get Database Port from Environment Variable
	DBConfig.Port = os.Getenv("DB_PORT")
	if len(DBConfig.Port) == 0 {
		// If Database Port Environment Variable Not Exist
		// Then Set Database Port Variable to Default Value
		DBConfig.Port = "3306"
	}

	// Get Database User from Environment Variable
	DBConfig.User = os.Getenv("DB_USER")
	if len(DBConfig.User) == 0 {
		// If Database User Environment Variable Not Exist
		// Then Set Database User Variable to Default Value
		DBConfig.User = "user"
	}

	// Get Database Password from Environment Variable
	DBConfig.Password = os.Getenv("DB_PASSWORD")
	if len(DBConfig.Password) == 0 {
		// If Database Password Environment Variable Not Exist
		// Then Set Database Password Variable to Default Value
		DBConfig.Password = "password"
	}

	// Get Database Name from Environment Variable
	DBConfig.Name = os.Getenv("DB_NAME")
	if len(DBConfig.Name) == 0 {
		// If Database Name Environment Variable Not Exist
		// Then Set Database Name Variable to Default Value
		DBConfig.Name = "dbs"
	}

}
