package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"micro-service/service"
)

// 使用proto
func main() {
	user := &service.User{
		Name:   "bbz",
		Age:    21,
		Sex:    1,
		Status: 1,
	}
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshal)
	err = proto.Unmarshal(marshal, user)
	fmt.Println(user.String())
}
