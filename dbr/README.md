# [gocraft/dbr](https://github.com/gocraft/dbr)
```go
//NULL許可のカラムはdbr.NULLStringのようにしないとだめ
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

//接続
conn, err := dbr.Open("mysql", "username:password@tcp(host:port)/dbname", nil)
if err != nil {
  panic(err)
}
sess := conn.NewSession(nil)

//select
var users []User
count, err := sess.Select("*").From("user_table").Load(&users)
if err != nil {
  fmt.Println(err)
}
fmt.Println(count)
fmt.Println(users[1].Password.String)
//[{1 watanabe y.watanabe@gandh.jp {{test true}} {{ts true}} 0 0 0 0} {2 test  {{ false}} {{ false}} 0 0 0 0}]
```
**型をdbr.NULLStringのようにしたものは構造体で返ってくる**

## NULLStringなどについて
```golang
type NullString struct {
  String string
  Valid bool // Valid is true if String is not NULL
}

type NullFloat64 struct {
  Float64 float64
  Valid bool // Valid is true if Float64 is not NULL
}

type NullInt64 struct {
  Int64 int64
  Valid bool // Valid is true if Int64 is not NULL
}

type NullBool struct {
  Bool bool
  Valid bool // Valid is true if Bool is not NULL
}

//dbr独自のNULL許容
type NullTime struct {
  Time  time.Time
  Valid bool // Valid is true if Time is not NULL
}
```