package database

import (
	//"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	//_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	//"log"
)

//var SqlDB *sql.DB
var MogSession *mgo.Session
var Redis redis.Conn
func init() {
}
//var mgodb *mgo.Database
func Config(mogo_url string,redis_url string) {
	var err error
	var mgoerr error
	// SqlDB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test?parseTime=true")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = SqlDB.Ping()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	MogSession, mgoerr = mgo.Dial(mogo_url)
	if mgoerr != nil {
		panic(mgoerr)
	}
	MogSession.SetMode(mgo.Monotonic, true)
	Redis, err = redis.Dial("tcp", redis_url)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	//  defer mogSession.Close()
	//  session.SetMode(mgo.Monotonic, true)
	//  mgodb = session.DB("egg_cnode")
	//  countNum, _ :=mgodb.C("users").Count()
	//  log.Println(countNum)
}
