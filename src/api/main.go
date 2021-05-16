package main

import(
	"github.com/shwethadia/Testing/src/api/app"
	"database/sql"
)

var (
	dbClient *sql.DB
)
const (

	queryGetUser = "SELECT id,email FROM users WHERE id=%d;"
)
func init(){

	var err error
	dbClient , err := sql.Open("mysql","this is the connection string")
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
	app.StartApp()
}

func GetUser(id int64) (*User , error ){


	rows, err := dbClient.Query(fmt.Sprintf(queryGetUser,id))
	if err != nil {
		return nil,err
	}

	var user User
	for rows.Next(){
		
		if err:= rows.Scan(&user.Id, &user.Email); err != nil {
		
			return nil,err
		}
		return &user,nil
	}

	return nil, errors.New("User not found")
}