package consulConfig

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net"
)

func httpGet(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return ""
	}

	return string(body)
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}
	var localIp string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIp = ipnet.IP.String()
				break
			}
		}
	}
	return localIp
}