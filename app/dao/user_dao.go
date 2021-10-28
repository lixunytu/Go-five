package dao

import (
	"github.com/pkg/errors"
	"go_test/app/databases"
)

type User struct {
	Name string
	Id int
}

func GetUserById(id int)(User, error)  {
	var user01 User
	row := databases.Db.QueryRow("select id ,name from user2 where id = ?" ,id )
	err := row.Scan(&user01.Id,&user01.Name)
	if err != nil{
		return user01,errors.Wrap(err,"dao#GetUserById err")
	}

	return user01,nil
}