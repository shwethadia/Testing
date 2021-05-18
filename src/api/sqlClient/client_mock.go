package sqlClient

import "database/sql"


var (

	isMocked bool
)


func StartMockupServer() {

	isMocked = true
}

func stopMockupServer(){

	isMocked = false
}


type clientMock struct {

	mocks map[string]Mock	
}

type Mock struct {

	Query string
	Args []interface{}
	Error error
	Columns []string
	Rows  [][]interface{}
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {

	mock, exists := c.mocks[query]
	if !exists {
		return nil, errors.New("No Mock available")
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	rows := rowsMock{

		Columns : mock.Columns,
		Rows : mock.Rows,

	}
	return &rows,nil
}

func AddMock(mock Mock ){

	client , okType := dbClient.(*clientMock)

	if !okType {
		return 
	}

	if client.mocks == nil {

		client.mocks = make(map[string]Mock, 0)
	}
	client.mocks[mocks.Query] = mock

}
