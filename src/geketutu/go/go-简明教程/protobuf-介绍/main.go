/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-23 22:05:23
 */
package main

import (
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	test := &Student{
		Name:   "Zhouzw",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error : ", err)
	}
	if test.GetName() != newTest.GetName() {
		log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
	}

}
