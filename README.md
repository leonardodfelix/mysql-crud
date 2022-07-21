# Mysql CRUD GO Application

## Commands

```bash
$ go mod init github.com/username/mysql-crud
$ go get "github.com/jinzhu/gorm"
$ go get "github.com/jinzhu/gorm/dialects/mysql"
$ go get "github.com/gorilla/mux"
```

1. Go to `/cmd/main` folder and run `go build`.
2. Change the second param of the `gorm.Open` function to your *`user`:`password`@/`database_name`?charset=utf8&parseTime=True&loc=Local*