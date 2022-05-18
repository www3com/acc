package uid

import (
	"fmt"
	"testing"
)

func TestUid(t *testing.T) {
	var i int64
	for i = 1000000000; i < 1000000009; i++ {
		uid := Id2Uid(i)
		id := Uid2Id(uid)
		fmt.Println(i, uid)
		if id != i {
			t.Errorf("测试失败， ID：%v", i)
		}
	}
}
