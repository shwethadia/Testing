package app

import "github.com/shwethadia/Testing/src/api/controllers"

func mapUrls(){

	router.GET("/locations/countries/:country_id",controllers.GetCountry)

}