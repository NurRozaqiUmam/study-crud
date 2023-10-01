package main

import "echo_crud/cmd"

// @title SWAGGER API
// @version 1.0
// @description This is swagger api
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.RunServer() // memanggil fungsi RunServer dari package cmd untuk menjalankan server
}
