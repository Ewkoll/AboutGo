package tags_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ewkoll/aboutgo/schema"
	"github.com/ewkoll/aboutgo/tags"
)

func TestAnalysisTags(t *testing.T) {
	var userName string
	typeOfData := reflect.TypeOf(userName)
	fmt.Println(typeOfData.Align(),
		typeOfData.FieldAlign(),
		typeOfData.Name(),
		typeOfData.PkgPath(),
		typeOfData.Size(),
		typeOfData.String(),
		typeOfData.Kind())
	// panic: reflect: Len of non-array type string
	// Align
	methodCount := typeOfData.NumMethod()
	for i := 0; i < methodCount; i++ {
		fmt.Println(typeOfData.Method(i))
	}

	user := schema.User{}
	tags.AnalysisTags(user)
}
