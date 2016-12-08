package consulConfig

import (
	"encoding/json"
	"strings"
	"fmt"
	"encoding/base64"
)

type ConsulKvs []ConsulKv
type ConsulKv struct {
	LockIndex	int		`json:"LockIndex"`
	Key		string		`json:"Key"`
	Flags		int		`json:"Flags"`
	Value		string		`json:"Value"`
	CreateIndex	int		`json:"CreateIndex"`
	ModifyIndex	int		`json:"ModifyIndex"`
}

var getRecurseFlag = "recurse"

func GetAllKvPairs(host string, keyPrefix string)*ConsulKvs{
	url := "http://"+host+"/v1/kv/" + keyPrefix + "?" + getRecurseFlag
	resp := httpGet(url)
	newKvs, kvs := &ConsulKvs{}, &ConsulKvs{}

	json.Unmarshal([]byte(resp), kvs)
	for _, v := range *kvs{
		dec, err := base64.StdEncoding.DecodeString(v.Value)
		if err != nil {
			fmt.Println("Consul KV Error: ", v.Key, v.Value, "can not be decoded.", err.Error())
			v.Value = ""
			continue
		}
		if (strings.HasPrefix(v.Key, keyPrefix)){
			v.Key = strings.Replace(v.Key, keyPrefix, "", 1)
		}
		v.Value = string(dec)
		*newKvs = append(*newKvs, v)
	}
	return newKvs
}

func GetAllKvJson(host, keyPrefix string) (string){
	ckvs := GetAllKvPairs(host, keyPrefix)
	kvMap := map[string]string{}
	for _, v := range *ckvs {
		kvMap[v.Key] = v.Value
	}
	result, err := json.Marshal(kvMap)
	if err != nil {
		fmt.Println("Consul KV Error: ", err.Error())
		return ""
	}
	return string(result)
}