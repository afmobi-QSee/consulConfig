package consulConfig

import (
	"fmt"
	"encoding/json"
)

type ConsulCatalogServices map[string][]string
type ConsulServices []ConsulService
type ConsulService struct {
	Node				string		`json:"Node"`
	Address				string		`json:"Address"`
	ServiceID			string		`json:"ServiceID"`
	ServiceName			string		`json:"ServiceName"`
	ServiceTags			[]string	`json:"ServiceTags"`
	ServiceAddress			string		`json:"ServiceAddress"`
	ServicePort			int		`json:"ServicePort"`
	ServiceEnableTagOverride	bool	 	`json:"ServiceEnableTagOverride"`
	CreateIndex			int	 	`json:"CreateIndex"`
	ModifyIndex			int	 	`json:"ModifyIndex"`
}

func (service *ConsulService) HasTag(tag string)(bool){
	for _, v := range service.ServiceTags {
		if v == tag {
			return true
		}
	}
	return false
}

func GetServices(host, serviceName string)(*ConsulServices){
	consulServices := ConsulServices{}
	url := "http://" + host + "/v1/catalog/service/"
	resp := httpGet(url + serviceName)
	if resp == "" {
		return nil
	}
	err := json.Unmarshal([]byte(resp), &consulServices)
	if err != nil {
		fmt.Println("Consul KV Error: ", err.Error())
		return nil
	}

	return &consulServices
}

type ConsulServicesRequest struct {
	ServiceName	string		`json:"serviceName"`
	Addrs		[]string	`json:"addrs"`
}
