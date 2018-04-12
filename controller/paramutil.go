package controller

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func MustInt(res *int, arg string) {
	var err error
	*res, err = strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}
}

func MayInt(res *int, arg string, defaltNum int) {
	if arg != "" {
		MustInt(res, arg)
	} else {
		*res = defaltNum
	}
}

func MustUnmarshal(res interface{}, arg string) {
	err := json.Unmarshal([]byte(arg), res)
	if err != nil {
		panic(err)
	}
}

func MayUnmarshal(res interface{}, arg string) {
	if arg != "" {
		MustUnmarshal(res, arg)
	}
}

func MayObjectId(res **bson.ObjectId, arg string) {
	if arg != "" && arg != "None" && arg != "null" {
		id := bson.ObjectIdHex(arg)
		*res = &id
	}
}
