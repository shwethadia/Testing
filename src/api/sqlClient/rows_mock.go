package sqlClient


import (

	"errors"
)

type rowsMock struct {

	Columns []string
	Rows   [][]interface{}

	currentIndex int
}

func (m *rowsMock) HasNext() bool {


	return m.currentIndex < len(m.Rows)
	
}

func (m *rowsMock) Close() error{

	return nil
}

func  (m *rowsMock) Scan(destinations  ...interface{}) error {

	mockedRow := m.Rows[m.currentIndex]
	if len(currentRow) != len(destinations) {

		return errors.New("Invalid Destination len")
	}
	for index, value := mockedRow {

		destinations[index] = value
	}

	return nil

}

