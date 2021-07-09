databox 是为了解决 处理复杂 golang 结构编写的代码

我们经常需要对复杂的结构做处理，比如取结构内部的一个值，或者设置一个值，这时候golang 编写代码需要层层断言，才能实现。

解决这个痛点，就可以把golang 当php 用。


package main

import "github.com/codewangz/databox/utils"

func main() {
  
        //map 增加新的值
	a := map[string]interface{}{"0":"0","1":1,"2":"2","3":"3"}
	dbx := utils.NewDataBox(a)
	dbx.Set("a","2")
	fmt.Println(dbx.Data())
	//结果 map[0:0 1:1 2:2 3:3 a:2]

	//slice index 五号位置 设置为 5
	b := []interface{}{1,2,3}
	dbx = utils.NewDataBox(b)

	dbx.Set("5","5")
	fmt.Println(dbx.Data())
	// 结果 [1 2 3 <nil> <nil> 5]

	//空变量设置值
	var c interface{}
	dbx = utils.NewDataBox(c)
	dbx.Set("a","a")
	dbx.Set("b.0","b")
	fmt.Println(dbx.Data())
	// 结果 map[a:a b:[b]]
}


