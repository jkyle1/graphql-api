package postgres

import (
	"database/sql"
	"fmt"
	//postgres driver
	_"github.com/lib/pq"
)

//database struct for interacting with db
type Db struct {
	*sql.DB
}

//makes new conn string and returns it or returns error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName,
		)
}

type User struct {
	ID int
	Name string
	Age int
	Profession string
	Friendly bool
}

func(d *Db) GetUsersByName(name string) []User {
	//prep query, take name argument (to protect against SQL injection)
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preparation Error: ", err)
	}

	//make query passing in name arg
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	//create User struct for each row's data
	var r User
	//create slice of Users for our response
	users :=[]User{}
	//copy columns from row to values pointed at by r
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
			)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}
	return users
}