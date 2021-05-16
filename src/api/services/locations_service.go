package services

import (
	"github.com/shwethadia/Testing/src/api/domain/locations"
	"github.com/shwethadia/Testing/src/api/providers/locations_provider"
	"github.com/shwethadia/Testing/src/api/utils/errors"
)

type locationService struct{}

type locationServiceInterface interface{

	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationService locationServiceInterface
)

func init(){

	LocationService = &locationService{}

}
func (s *locationService)GetCountry(countryId string) (*locations.Country, *errors.ApiError){

	return locations_provider.GetCountry(countryId)
}