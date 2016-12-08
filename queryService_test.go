package consulConfig

import (
	"testing"
	"fmt"
)

func Test_GetServices(t *testing.T){
	consulServices := GetServices("127.0.0.1:8500", "imc")
	fmt.Println(consulServices)
}