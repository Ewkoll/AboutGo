package tags

import (
	"fmt"
	"reflect"
)

func AnalysisTags(data interface{}) {
	t := reflect.TypeOf(data)
	NumField := t.NumField()
	for i := 0; i < NumField; i++ {
		fmt.Println(t.Field(i).Name, t.Field(i).Tag)
		tag := t.Field(i).Tag
		orm, ok := tag.Lookup("orm")
		jsondata := tag.Get("json")
		fmt.Println(orm, ok, jsondata)
	}
}
