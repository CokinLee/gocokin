package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/**
 * http://denis.papathanasiou.org/2012/10/14/go-golang-and-mongodb-using-mgo/
 */
/*var (
	session *mgo.Session
	db      *mgo.Database
)
*/
/**
 * host: 数据库地址
 * return *mgo.Session
 */
/*func getSession(host string) *mgo.Session {
	if host == "" {
		host = beego.AppConfig.String("mongodb_host")
	}
	if host == "" {
		host = "localhost"
	}
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	return session.Clone()
}
*/
/**
 * collection: 相当于mysql的数据表名
 * dbName: 数据库名
 */
/*
func withCollection(dbName, collection string, fn func(*mgo.Collection) error) error {
	if dbName == "" {
		dbName = beego.AppConfig.String("mongodb_db")
	}
	if dbName == "" {
		dbName = "example"
	}

	session := getSession("")
	defer session.Close()
	c := session.DB(dbName).C(collection)
	return fn(c)
}
*/
type Items map[string]interface{}

type Frds struct {
	Name string
	Age  int
}

type Addr struct {
	Province string
	City     string
}

type Person struct {
	Name    string
	Phone   string
	Age     int
	Address Addr
	Friends []Frds
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	cred := &mgo.Credential{Username: "user", Password: "123456", Source: "test", Service: "mongodb"}
	session.Login(cred)
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("test")
	/*
			b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
			var f interface{}
			err = json.Unmarshal(b, &f)
			m := f.(map[string]interface{})
			err = c.Insert(m)


		if err != nil {
			panic(err)
		}
	*/
	friends := []Frds{Frds{"李四", 28}, Frds{"王五", 26}}
	friends2 := []Frds{Frds{"张三", 25}, Frds{"李四", 28}}
	err = c.Insert(&Person{"王五", "+86153-8116-9639", 25, Addr{"北京", "朝阳区"}, friends2},
		&Person{"Cla", "+86153-8402-8510", 26, Addr{"四川", "成都"}, friends})

	if err != nil {
		panic(err)
	}
	result := []Person{}
	err = c.Find(bson.M{"friends": Items{"$elemMatch": Items{"name": "李四", "age": Items{"$gt": 20}}}}).All(&result)
	if err != nil {

		fmt.Println(err.Error() == "not found")
	}
	fmt.Println(result)
	for _, v := range result {
		fmt.Println(v.Friends)
	}

}
