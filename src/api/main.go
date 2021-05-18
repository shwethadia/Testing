package main

import (
	"errors"
	"fmt"

	"github.com/shwethadia/Testing/src/api/app"
	"github.com/shwethadia/Testing/src/api/sqlClient"
)

var (
	dbClient sqlClient.SqlClient
)
const (

	queryGetUser = "SELECT id,email FROM users WHERE id=%d;"
)
func init(){


	var err error
	dbClient , err =  sqlClient.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	"root","root","127.0.0.1:3306","users_db"))
	if err != nil {
		panic(err)
	}
}

type User struct {

	Id int64
	Email string
}

func main(){

	
	user , err := GetUser(123)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
	fmt.Println(user.Id)
	app.StartApp()
}

func GetUser(id int64) (*User , error ){

	sqlClient.AddMock(sqlClient.Mock{

		Query : "SELECT id,email FROM users WHERE id=?;",
		Args : []interface{}{1},
		Error: errors.New("Error creating query"),
		Columns : []string {"id","email"},
		Rows : [][]interface{}{
			{1 , "email1"},
			{2,  "email2"},
		},

	})
	rows, err := dbClient.Query(fmt.Sprintf(queryGetUser,id))
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var user User
	for rows.HasNext(){	
		if err:= rows.Scan(&user.Id, &user.Email); err != nil {
			return nil,err
		}
		return &user,nil
	}
	return nil, errors.New("User not found")
}