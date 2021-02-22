package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	jso, err := ioutil.ReadFile(fmt.Sprintf("ceping_item.json")) //datadict/ceping_item.json
	if err != nil {
		fmt.Println(err)
	}
	jss := make([]interface{}, 0)
	_ = json.Unmarshal(jso, &jss)
	//			logger.Error(jss["Android"])
	//var items []string
	//result := make(map[string]interface{})
	i := 0
	str := ""
	for _, i2 := range jss {
		if i2.(map[string]interface{})["type"] == "Android" { //Android   SDK   iOS
			//if i2.(map[string]interface{})["item"].(string) == "sec_infos" {
			//	continue
			//}
			str = str + i2.(map[string]interface{})["item"].(string) + ","
			/*	fmt.Println(i2.(map[string]interface{})["item"])
				items = append(items, i2.(map[string]interface{})["item"].(string))
				result[i2.(map[string]interface{})["item"].(string)] = i2.(map[string]interface{})["desc"].(string)*/
			i++
		}
	}
	fmt.Println(i)
	fmt.Println(str)
}
