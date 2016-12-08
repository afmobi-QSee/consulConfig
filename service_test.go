package consulConfig

import "testing"

func Test_service(t *testing.T){
	consulConfig := NewConsulConfig("172.17.40.21:8500", "testServiceName", "10.2.227.131", 6379, []string{"testTag"})
	consulConfig.Register()
	consulConfig.DeRegister()
}
