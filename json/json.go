package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonObj struct {
	Data  map[string]interface{} `json:"data"`
	Other map[string]interface{} `json:"-"`
}

func main() {

	uniqMap := make(map[string]struct{})

	b, _ := ioutil.ReadAll(os.Stdin)
	j := b[bytes.Index(b, []byte("\n"))+1:]

	var objs []jsonObj
	json.Unmarshal(j, &objs)

	for _, obj := range objs {
		if value, ok := obj.Data["key"]; ok {
			uniqMap[value.(string)] = struct{}{}
		}
	}
	answer := 0
	for range uniqMap {
		answer++
	}
	fmt.Println(answer)
}
