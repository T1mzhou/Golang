package main

import (
	"encoding/json"
	"fmt"
)

/*

{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
}

*/

//成员变量名首字母必须大写
//type IT struct {
//	Company  string
//	Subjects []string
//	IsOk     bool
//	Price    float64
//}

func main() {
	// 创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "tencent"
	m["subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 666.666
	// 编码成json
	// result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))

}
