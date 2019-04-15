# vasiliyaltunin/dbre (gocraft/dbr wrapper)


gocraft/dbr wrapper that allow you do simple sql query - select, insert, update and delete just in one string


```
$ go get -u github.com/vasiliyaltunin/dbre
```

```go
import "github.com/vasiliyaltunin/dbre"
```

## Examples

```go
//User - a site user
type User struct {
	UID       string `db:"uid" form:"uid"`
	UserID    string `db:"user_id" form:"user_id"`
	UserGroup string `db:"user_grp" form:"user_grp"`
	Email     string `db:"email" form:"email"`
}

//UserList - a list of site user
type UserList []User
```

### Select

```go
selectVal := dbre.DbrExpress(Session).Select("users", "*", "uid>5", user)
vals := (*selectVal.(*UserList))
```

### Insert 

```go
dbre.DbrExpress(Session).Insert("users", []string{"user_id", "email"}, User{UserID: "111112221111", Email: "mail@mail.com"})
```
### Update

```go
dbre.DbrExpress(Session).Update("users", []string{"UserID", "Email"}, User{UserID: "999112221111", Email: "mail@mail.com"}, "`uid` = ?", "5")
```
### Delete

```go
dbre.DbrExpress(Session).Delete("users", "`uid` = ?", "6")
```

# License
MIT
