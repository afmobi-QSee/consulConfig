package consulConfig

import (
	"testing"
	"fmt"
)

func Test_httpGet(t *testing.T){
	result := httpGet("http://www.github.com")
	if result == ""{
		t.Error("httpGet test failed")
	}
}

func Test_GetLocalIp(t *testing.T){
	ip := GetLocalIp();
	fmt.Print(ip)
}