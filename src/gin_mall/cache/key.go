/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-11 19:35:57
 */
package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

func ProductViwKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
