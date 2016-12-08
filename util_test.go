package consulConfig

import "testing"

func Test_httpGet(t *testing.T){
	result := httpGet("http://www.github.com")
	if result == ""{
		t.Error("httpGet test failed")
	}
}