databox 是为了解决 处理复杂 golang 结构编写的代码

我们经常需要对复杂的结构做处理，比如取结构内部的一个值，或者设置一个值，这时候golang 编写代码需要层层断言，才能实现。

解决这个痛点，就可以把golang 当php 用。

package main

import "github.com/codewangz/databox/utils"

func main() {
  
  var data interface{}
	data = map[string]string{"b":"aa"}
	//data = []interface{}{0,1}
	//data := []interface{}{}
	//data["a"] = 1
	dbbox := NewDataBox(data)
	//d := dbbox.createData([]string{"a","a"},"111")
	//fmt.Println(d)
	dbbox.Set("a","aaa")
	dbbox.Set("b.1","111")
	dbbox.Set("b.0","222")
	dbbox.Set("b.1","222")

	//dbbox.Set("c","aaa")
	//dbbox.Set("b.1","ccc")
	//dbbox.Set("b.1","222")
	//dbbox.Set("b.0","222")
	fmt.Println(dbbox.Data())
	fmt.Println(dbbox.Get("d"))
}


