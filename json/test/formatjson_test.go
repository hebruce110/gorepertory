package example

import (
	"fmt"
	"gorepertory/json"
	"testing"
)

func TestEncodeToJson(t *testing.T)  {
	var a map[string]interface{}
	a=map[string]interface{}{
		"k1":"hellow orld!",
		"k2":2332,
		"k3":false,
		"k4":nil,
		"k5":[]string{"f","i","j"},
		"k6":map[string]interface{}{"foo":"bar"},
	}
	s:=formatjson.Encode(a)
	fmt.Println(s)
}

func TestDecodeJson(t *testing.T)  {
	s:= `{
    "array": [
        1,
        "2",
        3
    ],
    "int": 10,
    "float": 5.15,
    "bignum": 9223372036854775807,
    "string": "simplejson",
    "bool": true
	}`
	i:=formatjson.Decode(s)
	fmt.Println(i)
	m := i.(map[string]interface{})
	for k,v := range m {
		fmt.Println(k,"=>",v)
	}

}