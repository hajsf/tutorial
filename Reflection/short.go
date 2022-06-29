package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// We'll get a `problem` that:
// Struct field tag `name` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair
// But the code will work, Struct tag supposed to be a key:"value", field:"name"
// No worry about this, as in our code below we are not using Tag attribute

type User struct {
	Name string `name`
	Age  int64  `age`

	// Active bool // uncomit this to get an error at JSONEncode, nothing is defined for `bool` type
}

func main() {
	// Use the correct line from the 2 below depending on either the `bool` field is active or not
	var u User = User{"bob", 10}
	// var u User = User{"bob", 10, true}

	res, err := JSONEncode(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

}

func JSONEncode(v interface{}) ([]byte, error) {
	refObjVal := reflect.ValueOf(v)
	refObjTyp := reflect.TypeOf(v)
	buf := bytes.Buffer{}
	if refObjVal.Kind() != reflect.Struct {
		return buf.Bytes(), fmt.Errorf(
			"val of kind %s is not supported",
			refObjVal.Kind(),
		)
	}
	buf.WriteString("{")
	pairs := []string{}
	for i := 0; i < refObjVal.NumField(); i++ {
		structFieldRefObj := refObjVal.Field(i)
		structFieldRefObjTyp := refObjTyp.Field(i)

		switch structFieldRefObj.Kind() {
		case reflect.String:
			strVal := structFieldRefObj.Interface().(string)
			pairs = append(pairs, `"`+string(structFieldRefObjTyp.Tag)+`":"`+strVal+`"`)
		case reflect.Int64:
			intVal := structFieldRefObj.Interface().(int64)
			pairs = append(pairs, `"`+string(structFieldRefObjTyp.Tag)+`":`+strconv.FormatInt(intVal, 10))
		default:
			return buf.Bytes(), fmt.Errorf(
				"struct field with name %s and kind %s is not supprted",
				structFieldRefObjTyp.Name,
				structFieldRefObj.Kind(),
			)
		}
	}

	buf.WriteString(strings.Join(pairs, ","))
	buf.WriteString("}")

	return buf.Bytes(), nil
}
