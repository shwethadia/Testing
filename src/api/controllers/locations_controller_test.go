package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"

	"github.com/shwethadia/Testing/src/api/domain/locations"
	"github.com/shwethadia/Testing/src/api/services"
	"github.com/shwethadia/Testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc  func(countryId string) (*locations.Country, *errors.ApiError)
)
func TestMain(m *testing.M){

	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationServiceMock struct {

}

func(*locationServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError){

	return getCountryFunc(countryId)

}


func TestGetCountryNotFound(t *testing.T){

	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError){

		return nil, &errors.ApiError{
			Status: http.StatusNotFound,
			Message: "Country not found",
		}
	}

	services.LocationService = &locationServiceMock{}
	response := httptest.NewRecorder()
	c , _:= gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet,"",nil)
	c.Params =  gin.Params{
		{Key : "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var apiErr errors.ApiError

	err := json.Unmarshal(response.Body.Bytes(),&apiErr)
	assert.Nil(t,err)
	assert.EqualValues(t,http.StatusNotFound,apiErr.Status)
	assert.EqualValues(t,"Country Not found",apiErr.Message)
	
}

func TestGetCountryNotError(t *testing.T){

	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError){

		return &locations.Country{Id:"AR",Name:"Argentina"}, nil
	}

	services.LocationService = &locationServiceMock{}
	response := httptest.NewRecorder()
	c , _:= gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet,"",nil)
	c.Params =  gin.Params{
		{Key : "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country locations.Country

	err := json.Unmarshal(response.Body.Bytes(),&country)
	assert.Nil(t,err)
	assert.EqualValues(t,"AR",country.Id)
	assert.EqualValues(t,"Argentina",country.Name)
	
}