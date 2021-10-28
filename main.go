package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go_test/app/dao"
	"go_test/app/databases"
	"log"
)

func QueryService(id int) ( dao.User,error) {
	//查询
	u, err := dao.GetUserById(id)
	if err!=nil {
		//处理error
		return u,errors.Wrap(err,"service#QueryService err")
	}
	return u,nil
}

func main() {
	err := databases.Mysql_Init()
	if err!=nil {
		//处理error
		log.Println(err)
		panic(err)
	}

	id:=1
	u, err := QueryService(id)

	if err!=nil {
		//处理error
		if errors.Is(err,sql.ErrNoRows) {
			log.Println("user is not found")
		}
		fmt.Println(err)
		return
	}

	fmt.Println(u)
}
