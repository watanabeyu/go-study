package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	//"github.com/gocraft/dbr/dialect"

	"fmt"
)

type User struct {
	Uid         int            `db:"uid"`
	Username    string         `db:"username"`
	Email       string         `db:"email"`
	Password    dbr.NullString `db:"password"`
	Facebook_id dbr.NullString `db:"facebook_id"`
	Last_login  int            `db:"last_login"`
	Delete_flg  int            `db:"delete_flg"`
	Created     int            `db:"created"`
	Modified    int            `db:"modified"`
}

func main() {
	//connect
	conn, err := dbr.Open("mysql", "username:password@tcp(host:port)/dbname", nil)
	if err != nil {
		panic(err)
	}
	sess := conn.NewSession(nil)

	//select all user
	var users []User
	count_users, err := sess.Select("*").From("user_table").Load(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count_users)
	fmt.Println(users[0].Password.String)

	//select user uid=1
	var user User
	count_user, err := sess.Select("*").From("user_table").Where("uid = ?", 1).Load(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count_user)
	fmt.Println(user.Password.String)
}
