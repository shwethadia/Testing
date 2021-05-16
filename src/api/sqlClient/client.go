package sqlClient

import (
	"errors"
)
type client struct {

	db *sql.DB
}


type row struct {}


type SqlClient interface{

	Query(query string, args ...interface{}) (*row, error)
}

func Open(driverName,dataSourceName string) (SqlClient, error){

	if driverName == ""{

		return nil, errors.New("invalid driver name")
	}
	db , err := sql.Open(driverName,dataSourceName)
	if err != nil {
		return nil, err 
	}
	client := &client {

		db : db,

	}

	return client, nil
}

func (c *client) Query(query string, args ...interface{}) (*row, error){

	return nil,nil

}