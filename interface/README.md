## interface + echo + application/json post
このjsonをpostされた場合に処理をしたい
```json
{
  "email":"y.watanabe@gandh.jp",
  "password":"password",
  "hoge":["test1", "test2"],
  "foo":{
    "bar":"test3",
    "barbar":"test4"
  }
}
```

### structはこうしよう
FormValueで受け取ろうとするとうまくいかないので、structを使うべき  
-> **送信側のfetchがおかしかったため、問題なし**
-> **ただどちらにしろstruct使うと簡単(その場合はformタグを使う)**
```golang
type LoginForm struct {
  Email    string                 `json:"email"`
  Password string                 `json:"password"`
  Foo      map[string]interface{} `json:"foo"`
  Hoge     []string               `json:"hoge"`

  //下記のように単純なinterface型にするとform.Foo["bar"]としてとれない
  //Foo    interface{}            `json:"foo"`
}

//Fooの中が正しくわかっているのであれば下記が望ましい
type LoginForm struct {
  Email    string                 `json:"email"`
  Password string                 `json:"password"`
  Foo      map[string]string      `json:"foo"`
  Hoge     []string               `json:"hoge"`
}
```

### 処理
```golang
package handler

import (
  "github.com/labstack/echo"

  "fmt"
)

type LoginForm struct {
  Email    string                 `json:"email" form:"email"`
  Password string                 `json:"password" form:"password"`
  Foo      map[string]string      `json:"foo" form:"foo"`
  Hoge     []string               `json:"hoge" form:"hoge"`
}

func Login(c echo.Context) error {
  var form LoginForm
  err := c.Bind(&form)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(form)
    fmt.Println(form.Email)
    fmt.Println(form.Password)
    fmt.Println(form.Foo)
    fmt.Println(form.Foo["bar"])
    //fmt.Println(form.Foo["barbar"]) //単純なinterface型とするとmap的にとれない
    fmt.Println(form.Hoge)
    fmt.Println(form.Hoge[0])
    fmt.Println(form.Hoge[1])
  }

  email := c.FormValue("email")
  password := c.FormValue("password")
}
```

補足として、fileをアップする場合はbase64エンコードしてjsonに入れてpostでOK  
interfaceを処理するときはそれがどんな型なのか確認をして、アサーションする必要がある