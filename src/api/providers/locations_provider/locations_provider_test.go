package locations_provider

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)


func TestMain(m *testing.M){

	rest.StartMockupServer()
	os.Exit(m.Run())
}


func TestGetCountryRestClient(t *testing.T){

	//Init

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: 0,

	})
	//Execution
	country, err := GetCountry("AR")
	//Validation
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.Status)
	assert.EqualValues(t,"invalid restclient response error when getting coutry AR",err.Message)

}

func TestGetCountryNotFound(t *testing.T){

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody : `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})

	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusNotFound,err.Status)
	assert.EqualValues(t,"Country not found",err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T){

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody : `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.Status)
	assert.EqualValues(t,"Invalid error interface when getting country AR",err.Message)

}

func TestGetCountryInvalidJSONResponse(t *testing.T){

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody : `{"id":123,"name":"Argentina","time_zone":"GMT-03:00"}`,
	})
	country, err := GetCountry("AR")
	assert.Nil(t,country)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.Status)
	assert.EqualValues(t,"error when trying to unmarshal country AR",err.Message)

}


func TestGetCoutnryNoError(t *testing.T){


	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody : `{"id":123,"name":"Argentina","time_zone":"GMT-03:00"}`,
	})
	country, err := GetCountry("AR")
	assert.Nil(t,err)
	assert.NotNil(t,country)
	assert.EqualValues(t,"AR",country.Id)
	assert.EqualValues(t,"Argentina",country.Name)
	assert.EqualValues(t,"GMT-03:00",country.TimeZone)
	assert.EqualValues(t,24,len(country.States))
}