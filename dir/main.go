package main

import (
	"encoding/json"
	"fmt"
	"github.com/StackExchange/wmi"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func GetStorageInfo() map[string]interface{} {
	var storageinfo []storageInfo
	res := make(map[string]interface{})
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return nil
	}
	for _, storage := range storageinfo {
		if storage.Name != "" {
			res[storage.Name] = storage.Size
		}
	}
	return res
}

type Files struct {
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
}

func main() {
	filepath.Walk("D:\\app\\test",
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				fmt.Println("dir:", path)
				return nil
			}

			fmt.Println("file111:", path)

			spl := strings.Split(path, "\\")
			if spl[len(spl)-1] == "engine.json" {
				//	fmt.Println("file111:", path)
				fmt.Println(spl[len(spl)-1])
				jso, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
				}
				jss := make(map[string]interface{}, 0)
				_ = json.Unmarshal(jso, &jss)
				fmt.Println(jss)
				fmt.Println(jss["engine_name"])
			}
			return nil
		})
}

//func main()  {
////	var storageinfo []storageInfo
//	res := make(map[string]interface{})
//	//err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	//for _, storage := range storageinfo {
//	//	if storage.Name != "" {
//	//		res[storage.Name] = storage.Size
//	//	}
//	//}
//	//fmt.Println(res)
//
//
//	var rere []map[string]interface{}
//	//	myfolder := `D:\aaaa\nnn\wails_shield\backend`
//	myfolder := "D:\\mygowork2\\GUI"
//	//	listFile(myfolder)
//	files, _ := ioutil.ReadDir(myfolder)
//	for _, file := range files {
//		if file.IsDir() {
//			//	listFile(myfolder + "/" + file.Name())
//			fmt.Println(myfolder + "/" + file.Name())
//			rffi := map[string]interface{}{
//				"path":  myfolder + "/" + file.Name(),
//				"name":  file.Name(),
//				"isdir": true,
//			}
//			rere = append(rere, rffi)
//		} else {
//			//if pps["isdir"].(string) == "文件夹" {
//			//	continue
//			//}
//			fmt.Println(myfolder + "/" + file.Name())
//			rffi := map[string]interface{}{
//				"path":  myfolder + "/" + file.Name(),
//				"name":  file.Name(),
//				"isdir": false,
//			}
//			rere = append(rere, rffi)
//		}
//	}
////	res := make(map[string]interface{})
//	res["filetree"] = rere
//	fmt.Println(res)
//
//
//
//
//
//
//
//
//}
