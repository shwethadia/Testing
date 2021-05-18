package sqlClient

import (
	"errors"
)

const (

	goEnvironment = "GO_ENVIRONMENT"
	production = "production"
)

var(

	
	client SqlClient
)

type client struct {

	db *sql.DB
}

type SqlClient interface{

	Query(query string, args ...interface{}) (rows, error)
}


type isProduction()bool {
	return os.Getenv(goEnvironment) == production 
}

func Open(driverName,dataSourceName string) (SqlClient, error){

	if !isProduction() || isMocked {
		
		dbClient = &clientMock{}
		return &dbClient, nil 

	}
	if driverName == ""{

		return nil, errors.New("invalid driver name")
	}
	db , err := sql.Open(driverName,dataSourceName)
	if err != nil {
		return nil, err 
	}
	dbClient := &client {
		db : db,
	}
	return dbClient, nil
}

func (c *client) Query(query string, args ...interface{}) (rows, error){

	returnedRows , err := c.db.Query(query, args...)
	if err!= nil {
		return nil, err
	}

	result := sqlRows{

		rows: returnedRows,
	}

	return &result,nil

}