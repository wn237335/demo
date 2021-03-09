package main

import (
	"context"
	_ "encoding/json"
	"fmt"
	"github.com/asticode/go-astichartjs"
	_ "github.com/asticode/go-astilectron"
	_ "github.com/asticode/go-astilectron-bootstrap"
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

// Exploration represents the results of an exploration
type Exploration struct {
	Dirs       []Dir              `json:"dirs"`
	Files      *astichartjs.Chart `json:"files,omitempty"`
	FilesCount int                `json:"files_count"`
	FilesSize  string             `json:"files_size"`
	Path       string             `json:"path"`
}

// PayloadDir represents a dir payload
type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

/*func main()  {
	dir,_ := os.Getwd()
	fmt.Println("当前路径：",dir)
	//文件相对路径
//	dirPath := filepath.Dir(osPath)

	//文件名

//	fileName := filepath.Base(osPath)
}*/

/*func main()  {
	path := ""
	// If no path is provided, use the user's home dir

		var u *user.User
		u, err := user.Current()
		if err != nil {
			return
		}
		path = u.HomeDir
		fmt.Println(path)

	// Read dir
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(path); err != nil {
		return
	}

	// Init exploration
	e := Exploration{
		Dirs: []Dir{},
		Path: path,
	}

	// Add previous dir
	if filepath.Dir(path) != path {
		e.Dirs = append(e.Dirs, Dir{
			Name: "..",
			Path: filepath.Dir(path),
		})
	}

	// Loop through files
	var sizes []int
	var sizesMap = make(map[int][]string)
	var filesSize int64
	for _, f := range files {
		if f.IsDir() {
			e.Dirs = append(e.Dirs, Dir{
				Name: f.Name(),
				Path: filepath.Join(path, f.Name()),
			})
		} else {
			var s = int(f.Size())
			sizes = append(sizes, s)
			sizesMap[s] = append(sizesMap[s], f.Name())
			e.FilesCount++
			filesSize += f.Size()
		}
	}

	fmt.Println(e)

	// Prepare files size
	if filesSize < 1e3 {
		e.FilesSize = strconv.Itoa(int(filesSize)) + "b"
	} else if filesSize < 1e6 {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024), 'f', 0, 64) + "kb"
	} else if filesSize < 1e9 {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024), 'f', 0, 64) + "Mb"
	} else {
		e.FilesSize = strconv.FormatFloat(float64(filesSize)/float64(1024*1024*1024), 'f', 0, 64) + "Gb"
	}

	// Prepare files chart
	sort.Ints(sizes)
	if len(sizes) > 0 {
		e.Files = &astichartjs.Chart{
			Data: &astichartjs.Data{Datasets: []astichartjs.Dataset{{
				BackgroundColor: []string{
					astichartjs.ChartBackgroundColorYellow,
					astichartjs.ChartBackgroundColorGreen,
					astichartjs.ChartBackgroundColorRed,
					astichartjs.ChartBackgroundColorBlue,
					astichartjs.ChartBackgroundColorPurple,
				},
				BorderColor: []string{
					astichartjs.ChartBorderColorYellow,
					astichartjs.ChartBorderColorGreen,
					astichartjs.ChartBorderColorRed,
					astichartjs.ChartBorderColorBlue,
					astichartjs.ChartBorderColorPurple,
				},
			}}},
			Type: astichartjs.ChartTypePie,
		}
		var sizeOther int
		for i := len(sizes) - 1; i >= 0; i-- {
			for _, l := range sizesMap[sizes[i]] {
				if len(e.Files.Data.Labels) < 4 {
					e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizes[i])
					e.Files.Data.Labels = append(e.Files.Data.Labels, l)
				} else {
					sizeOther += sizes[i]
				}
			}
		}
		if sizeOther > 0 {
			e.Files.Data.Datasets[0].Data = append(e.Files.Data.Datasets[0].Data, sizeOther)
			e.Files.Data.Labels = append(e.Files.Data.Labels, "other")
		}
	}
	return

}*/

/*func main() {
	sysType := runtime.GOOS

	if sysType == "linux" {
		// LINUX系统
		fmt.Println(sysType)
	}

	if sysType == "windows" {
		// windows系统
		fmt.Println(sysType)

	}
}
*/

//加固任务表
type Task struct {
	UUID       string `genji:"uuid"`        //uid
	ClientKey  string `genji:"client_key"`  //客户密钥
	ClientName string `genji:"client_name"` //客户名称
	UserId     int    `genji:"user_id"`     //用户id
	CreatedAt  string `genji:"created_at"`  //创建时间
	EndAt      string `genji:"end_at"`      //结束时间
	Status     string `json:"status"`       //任务状态（排队中 处理中 成功 失败）
	Type       string `genji:"type"`        //加固类型  bticode
	SubmitType string `genji:"submit_type"` //提交类型  WEB  or  GUI
	TaskName   string `genji:"task_name"`   //任务名称
	Version    string `genji:"version"`     //版本号
	AllDate    string `genji:"all_date"`    //所有任务数据
}

func main() {

	db, err := genji.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db = db.WithContext(context.Background())

	// Query some documents
	//res, err := db.Query("SELECT *FROM task WHERE client_key = ? LIMIT ? OFFSET ?", "ct-rQuBr9KAjIUIzXwOP8n3LilFsprfyPzm", 1,2)
	as := fmt.Sprintf("SELECT *FROM task WHERE client_key = '%s' LIMIT %d OFFSET %d", "ct-rQuBr9KAjIUIzXwOP8n3LilFsprfyPzm", 2, 2)
	res, err := db.Query(as)
	//	res, err := db.Query("SELECT *FROM task WHERE client_key = ? LIMIT 2 OFFSET 2", "ct-rQuBr9KAjIUIzXwOP8n3LilFsprfyPzm")
	//	res, err := db.Query(as)
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
