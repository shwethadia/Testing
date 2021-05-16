package test

import "testing"


func TestCountriesNotFound(t *testing.T){

	fmt.Printf("About to functional test get countries")

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL : "https://api.mercadolibre.com/countries/AR",
		HTTPMethod: http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody : `{"status":404,"error":"not_found","message":"no country with id AR"}`,
	})
	response,err := http.Get("http://localhost:8080/locations/countries/AR")
	assert.Nil(t,err)
	assert.NotNil(t,response)
	bytes , _ := ioutil.ReadAll(response.Body)
	fmt.Println(bytes)

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.NotNil(t,err)

	assert.EqualValues(t, http.StatusNotFound,apiErr.Status)
	assert.EqualValues(t,"not_found",apiErr.Error)
	assert.EqualValues(t,"no country with id AR",apiErr.Message)
	
}
