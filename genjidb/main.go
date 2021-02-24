package main

import (
	"context"
	"fmt"
	"github.com/genjidb/genji"
	"github.com/genjidb/genji/document"
	"log"
)

func inster() string {
	db2, err := genji.Open("my.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer db2.Close()
	db2 = db2.WithContext(context.Background())
	err = db2.Exec("INSERT INTO User (id, name, age) VALUES (?, ?, ?)", 13, "zxcv", 18)
	return "ww"
}

//扫描目录
type FilePath struct {
	CreateAt  string `genji:"create_at"`  //创建时间
	Name      string `genji:"name"`       //用户名
	ClientKey string `genji:"client_key"` //客户密钥
	Path      string `genji:"path"`       //文件地址
}

//引擎地址
type EnginePath struct {
	CreateAt  string `genji:"create_at"`  //创建时间
	UserName  string `genji:"user_name"`  //用户名
	ClientKey string `genji:"client_key"` //客户密钥
	Type      string `genji:"type"`       //引擎类型
	Name      string `genji:"name"`       //引擎名称
	Path      string `genji:"path"`       //引擎地址
}

func main() {

	db, err := genji.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db = db.WithContext(context.Background())

	// Query some documents
	res, err := db.Query("SELECT *FROM enginepath")
	// always close the result when you're done with it
	defer res.Close()

	// Iterate over the results
	err = res.Iterate(func(d document.Document) error {
		// When querying an explicit list of fields, you can use the Scan function to scan them
		// in order. Note that the types don't have to match exactly the types stored in the table
		// as long as they are compatible.

		var u EnginePath
		err = document.StructScan(d, &u)
		if err != nil {
			return err
		}

		fmt.Println(u)

		// Or scan into a map
		var m map[string]interface{}
		err = document.MapScan(d, &m)
		if err != nil {
			return err
		}

		fmt.Println(m)
		return nil
	})
}

/*func main() {
	//	fmt.Println(1111)
	db, err := genji.Open("my.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	db = db.WithContext(context.Background())
	//	err = db.Exec("CREATE TABLE IF NOT EXISTS User")
	//	err = db.Exec("CREATE TABLE IF NOT EXISTS ser")
	//	err = db.Exec("INSERT INTO ser (id, name, age) VALUES (?, ?, ?)", 12, "zxcv", 17)

	//err = db.Update(func(tx *genji.Tx) error {
	//	err = tx.Exec("UPDATE User set age = ? WHERE name = ?", 99, "wangning")
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})

	type User struct {
		UserId     int    `genji:"user_id"` //用户id
		Name       string `genji:"name"`    //用户名
		RealName   string `genji:"real_name"`
		LastLogin  string `genji:"last_login"`  //上一次登录时间
		ClientKey  string `genji:"client_key"`  //客户密钥
		ClientName string `genji:"client_name"` //客户名称
		AllData    string `genji:"all_data"`    //所有数据
	}
	//err = db.Update(func(tx *genji.Tx) error {
	//	//err = tx.Exec("UPDATE User SET real_name = ? WHERE name = ?", "时间", "wangning")
	//	err = tx.Exec("UPDATE user SET all_data = ? WHERE name = ?", "时间", "wangning")
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	res, err := db.Query("SELECT *FROM user WHERE name = ?", "wangning")
	//err = res.Iterate(func(d document.Document) error {
	//	// When querying an explicit list of fields, you can use the Scan function to scan them
	//	// in order. Note that the types don't have to match exactly the types stored in the table
	//	// as long as they are compatible.
	//	var id int
	//	var name string
	//	var age int32
	//	var address struct {
	//		City    string
	//		ZipCode string
	//	}
	//
	//	err = document.Scan(d, &id, &name, &age, &address)
	//	if err != nil {
	//		return err
	//	}
	//
	//	fmt.Println(id, name, age, address)
	//
	//	// It is also possible to scan the results into a structure
	//	var u User
	//	err = document.StructScan(d, &u)
	//	if err != nil {
	//		return err
	//	}
	//
	//	fmt.Println(u)
	//
	//	// Or scan into a map
	//	var m map[string]interface{}
	//	err = document.MapScan(d, &m)
	//	if err != nil {
	//		return err
	//	}
	//
	//	fmt.Println(m)
	//	return nil
	//})

	// Query some documents
	//	res, err := db.Query("SELECT *FROM User WHERE name = ?", "wangning")
	// always close the result when you're done with it
	//	defer res.Close()

	tem := "bb"
	// Iterate over the results
	err = res.Iterate(func(d document.Document) error {
		var u User
		err = document.StructScan(d, &u)
		if err != nil {
			return err
		}

		fmt.Println(u.LastLogin)
		fmt.Println(u.AllData)
		fmt.Println(u.ClientKey)
		fmt.Println(u.ClientName)
		fmt.Println(u.RealName)
		fmt.Println(u.Name)
		fmt.Println(u.UserId)

		//var id int
		//var name string
		//var age int32
		//err = document.Scan(d, &id, &name, &age)
		//if err != nil {
		//	return err
		//}
		//if age == 18 {
		//	tem = "aa"
		//	fmt.Println(232323)
		//}
		//fmt.Println(id, name, age)
		// It is also possible to scan the results into a structure
		//	var m map[string]interface{}
		//	err = document.MapScan(d, &m)
		//	if err != nil {
		//		return err
		//	}
		//	fmt.Println(m)
		return nil
	})

	fmt.Println(tem)

}
*/
