package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" orm:"pk"`
	Sex  uint   `json:"sex"  orm:"null"`
}

func main() {
	user := &User{"xmge", 1}
	//通过反射获取type定义，传入参数必须为指针类型
	s := reflect.TypeOf(user).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
		fmt.Println(s.Field(i).Tag.Get("json"))
		fmt.Println(s.Field(i).Tag.Get("orm"))

		// blablabla....  根据相应的tag做相应的处理

	}
}
