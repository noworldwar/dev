package main

import (
	app "github.com/PGITAb/bc-service-entity-playerloginlog-go/internal/app/starter"
)

func main() {
	app.InitConfig("config", "./config", "../../config")
	app.InitMongo()
	app.GoMicroServerRun()
}
