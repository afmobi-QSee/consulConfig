package consulConfig

import (
	"encoding/json"
	"strconv"
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
)

type Check struct {
	Tcp string `json:"tcp"`
	Interval string `json:"interval"`
	Timeout string `json:"timeout"`
}

type Service struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Port int `json:"port"`
	Tags []string `json:"tags"`
	Check []Check `json:"checks"`
}

type ConsulConfig struct {
	Url string
	ServiceName string
	Ip string
	Port int
	Tag []string
}

func NewConsulConfig(url, serviceName, ip string, port int, tag []string) *ConsulConfig{
	return &ConsulConfig{Url:url, ServiceName:serviceName, Ip:ip, Port:port, Tag:tag}
}

func (this *ConsulConfig)Register() error{
	check := &Check{Tcp:this.Ip + ":" + strconv.Itoa(this.Port), Interval:"10s", Timeout:"1s"}
	service := &Service{Id:this.ServiceName + this.Ip + strconv.Itoa(this.Port), Name:this.ServiceName, Address:this.Ip, Port:this.Port, Tags: this.Tag, Check:[]Check{*check}}
	jbody, _ := json.Marshal(service)
	client := &http.Client{}
	fmt.Print("PUT", "http://" + this.Url + "/v1/agent/service/register", string(jbody))
	req, err := http.NewRequest("PUT ", "http://" + this.Url + "/v1/agent/service/register ", bytes.NewReader(jbody))
	if err != nil{
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (this *ConsulConfig)DeRegister() error {
	resp, err := http.Get("http://" + this.Url + "/v1/agent/service/deregister/" + this.ServiceName + this.Ip + strconv.Itoa(this.Port))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}