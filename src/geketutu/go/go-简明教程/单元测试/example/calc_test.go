/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-24 10:47:51
 */
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Error("出错")
	}
}
