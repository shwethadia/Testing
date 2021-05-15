package locations_provider

import (
	"github.com/shwethadia/Testing/src/api/domain/locations"
	"github.com/shwethadia/Testing/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"fmt"
	"net/http"
	"encoding/json"
)


const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError){

	res := rest.Get(fmt.Sprintf(urlGetCountry,countryId))
	if res == nil || res.Response == nil {
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s",countryId),
		}
	}

	if res.StatusCode > 299 {

		var apiErr errors.ApiError

		if err := json.Unmarshal(res.Bytes(),&apiErr); err != nil {

			return nil, &errors.ApiError{
				Status: http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s",countryId),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country

	if err := json.Unmarshal(res.Bytes(),&result) ; err != nil {
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s",countryId),
		}
	}
	return &result,nil
}