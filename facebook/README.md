# [huandu/facebook](https://github.com/huandu/facebook)
```go
//NULL許可のカラムはsql.NULLStringのようにしないとだめ
type fbResponse struct {
  Id       string `json:"id"`
  Name     string `json:"name"`
  Email    string `json:"email"`
  Birthday string `json:"birthday"`
  Gender   string `json:"gender"`
  Location loc    `json:"location"`
}

type loc struct {
  Id   sql.NullString `json:"id"`
  Name sql.NullString `json:"name"`
}

//接続
var globalApp = fb.New("AppId", "AppSecret")
session, _ := globalApp.SessionFromSignedRequest("facebook signed request")

err := session.Validate()

if err != nil {
  panic(err)
}

response, _ := session.Get("/me", fb.Params{
  "fields": "id,name,email,birthday,gender,location",
  "locale": "ja_JP",
})

var fbresponse fbResponse
util.MapToStruct(response, &fbresponse)
```
**responseを構造体に入れることで、取れなかったものをnilで返せる**

## mapToStruct
```golang
func MapToStruct(m map[string]interface{}, val interface{}) error {
  tmp, err := json.Marshal(m)
  if err != nil {
    return err
  }
  err = json.Unmarshal(tmp, val)
  if err != nil {
    return err
  }
  return nil
}
```