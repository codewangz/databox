package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


type dataBox struct {
	data interface{}
}

func NewDataBox(data interface{}) dataBox {
	//复制一份，防止原数据被修改
	return dataBox{Copy(data)}
}

func (dbx *dataBox) Get(key string) interface{} {
	var data interface{}
	paths := strings.Split(key, ".")
	data = dbx.data
	for _, path := range paths {
		dataType := reflect.TypeOf(data).String()
		i, err := strconv.Atoi(path)
		if err == nil && dataType == "[]interface {}" { //list
			tempData := data.([]interface{})
			if len(tempData) > i {
				data = tempData[i]
			} else {
				return nil
			}
		} else if dataType == "map[string]interface {}" {
			tempData := data.(map[string]interface{})
			data = tempData[path]
		} else {
			return nil
		}
	}

	return data
}

func (dbx *dataBox) Set(key string,val interface{}) {
	paths := strings.Split(key, ".")
	dbx.retSet(paths,val)
}

func (dbx *dataBox) retSet(paths []string,val interface{}) {
	data := dbx.createData(paths,val)
	dbx.data = dbx.mergeData(dbx.data,data,paths,0)
}

func (dbx *dataBox) createData(paths []string,val interface{}) interface{}{

		var tempData interface{}
		if index, err := strconv.Atoi(paths[len(paths)-1]) ; err == nil {
			tempData = []interface{}{}
			for i:=0;i <= index; i++ {
				tempData = append(tempData.([]interface{}),nil)
			}
			tempData.([]interface{})[index] = val
		}else{
			tempData = map[string]interface{}{paths[len(paths)-1]:val}
		}

		if len(paths) == 1 {
			return tempData
		}else{
			return dbx.createData(paths[0:len(paths)-1],tempData)
		}
}

func (dbx *dataBox) mergeData(dst interface{},src interface{},paths[]string,deep int) interface{}{
	if dst == nil {
		return src
	}

	if deep == len(paths) {
		return src
	}

	ok,kind := dbx.isSameKind(dst,src)
	if !ok {
		return src
		//panic("结构类型不一致，不能设置")
	}

	if kind == reflect.Slice {
		index,err := strconv.Atoi(paths[deep])
		if err != nil {
			panic(err)
		}
		srcSlice := Copy(src.([]interface{})[index])
		if ok,dstSlice := dbx.isInPath(kind,dst,paths[deep:]); ok {
			srcSlice = dbx.mergeData(dstSlice,srcSlice,paths,deep+1)
		}else{
			copy(src.([]interface{}),dst.([]interface{}))
			dst = src
		}

		dst.([]interface{})[index]  = srcSlice

	}else if kind == reflect.Map {
		srcval := Copy(src.(map[string]interface{})[paths[deep]])
		if ok,dstMap := dbx.isInPath(kind,dst,paths[deep:]); ok {
			srcval = dbx.mergeData(dstMap,srcval,paths,deep+1)
		}
		dst.(map[string]interface{})[paths[0]] = srcval
	}

	return dst

}

func (dbx *dataBox) isSameKind(dst,src interface{}) (bool,reflect.Kind){
	fmt.Println(dst,reflect.ValueOf(dst).Kind() ,src,reflect.ValueOf(src).Kind())
	return reflect.ValueOf(dst).Kind() == reflect.ValueOf(src).Kind(),reflect.ValueOf(src).Kind()
}

func  (dbx *dataBox) isInPath( kind reflect.Kind,dst interface{},path []string) (bool,interface{}){
	if kind == reflect.Map {
		if _,ok := dst.(map[string]interface{})[path[0]];ok {
			return ok,Copy(dst.(map[string]interface{})[path[0]])
		}
	}else if kind == reflect.Slice {
		if i, err := strconv.Atoi(path[0]);err == nil {
			if len(dst.([]interface{})) >i {
				return true,Copy(dst.([]interface{})[i])
			}
			return false,nil
		}else{
			panic(err)
		}
	}
	return false,nil
}

func (dbx *dataBox) Data() interface{}{
	return dbx.data
}

func Copy(data interface{}) interface{}{
	var result interface{}
	jsonByte,err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonByte,&result)
	if err != nil {
		panic(err)
	}
	return result
}


/*
func (dbx *dataBox) GetSlice(key string) []interface{}{
	val := dbx.Get(key)
	return ToSliceInterface(val)
}

func (dbx *dataBox) GetInt64(key string) int64{
	val := dbx.Get(key)
	return ToInt64(val)
}

func (dbx *dataBox) GetString(key string) string{
	val := dbx.Get(key)
	return ItoString(val)
}

func (dbx *dataBox) GetSliceString(key string) []string{
	val := dbx.Get(key)
	return ToSliceString(val)
}

func (dbx *dataBox) GetSliceMap(key string){

}

func (dbx *dataBox) GetMapInterface(key string) map[string]interface{}{
	val := dbx.Get(key)
	return ToMapInterface(val)
}

 */

