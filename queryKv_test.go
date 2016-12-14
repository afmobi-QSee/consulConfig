package consulConfig

import (
	"testing"
	"fmt"
)

func Test_GetAllKvJson(t *testing.T){
	jstring := GetAllKvJson("127.0.0.1:8500", "HttpServerAuth/")
	fmt.Println(jstring)
}