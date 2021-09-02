package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Common struct {
	Version int    `json:"version"`
	Kind    string `json:"kind"`
}

type FTP struct {
	Common
	UserName string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type WEB struct {
	Common
	URL string `json:"url"`
}

func parse(rawjson []byte) (interface{}, error) {
	type Both struct {
		*FTP
		*WEB
	}
	v := Both{}
	if err := json.Unmarshal(rawjson, &v); err != nil {
		return nil, err
	}

	if v.FTP != nil {
		return *v.FTP, nil
	}

	if v.WEB != nil {
		return *v.WEB, nil
	}

	return nil, errors.New("unknown config format")
}

func Read(path string) (interface{}, error) {
	rawjson, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parse(rawjson)
}

func CheckType(name string, config interface{}) {
	switch config.(type) {
	case WEB:
		fmt.Println(name, "is WEB config")
	case FTP:
		fmt.Println(name, "is FTP config")
	default:
		fmt.Println("Unknown config")
	}
}

func main() {
	tmp, err := Read("./v1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	CheckType("v1", tmp)

	tmp, err = Read("./v2.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	CheckType("v2", tmp)
}
