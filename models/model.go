package models

import (
	"fmt"
	"log"

	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
	"os"
	"time"
)

// Open ...
func Open() *gorm.DB {
	driver := "mysql"

	dataSource := os.Getenv("CLEARDB_DATABASE_URL")
	if dataSource != "" {
		dataSource = convertDataSource(dataSource)
	} else {
		dataSource = "root:pass@tcp(mysql:3306)/my_goal"
	}
	databaseConnect :=  dataSource + "?parseTime=true&charset=utf8"

	for i := 0; i < 30; i++ {
		var db *gorm.DB
		var err error
		if db, err = gorm.Open(driver, databaseConnect); err != nil {
			log.Println(err.Error())
			log.Println("データベースと接続できませんでした。")
			time.Sleep(time.Second * 10)
		} else {
			return db
		}
	}
	return nil
}

func convertDataSource(ds string) (result string) {
	parse, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", parse.User.String(), parse.Host, parse.Path)
	return result
}
