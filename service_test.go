package consulConfig

import "testing"

func Test_service(t *testing.T){
	consulConfig := NewConsulConfig("127.0.0.1:8500", "testServiceName", "10.2.227.131", 6379, []string{"testTag"})
	consulConfig.Register()
	consulConfig.DeRegister()
}
