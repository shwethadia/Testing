package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/shwethadia/Testing/src/api/app"
)



func TestMain(m *testing.M){
	rest.StartMockupServer()
	fmt.Println("About to start the application")
	app.StartApp()
	fmt.Println("Application started about to start the test cases")
	os.Exit(m.Run())
}