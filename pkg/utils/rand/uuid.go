package rand

import (
	"github.com/google/uuid"
	"strings"
)

// GoogleUUID32
// @description: google/uuid
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 13:08
// @success:
func GoogleUUID32() string {
	var l = 32
	randomStr := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(randomStr) > l {
		randomStr = randomStr[:l]
	} else if len(randomStr) < l {
		randomStr = randomStr + generateRandomString(l-len(randomStr))
	}
	return randomStr
}

// GoogleUUID20
// @description: 20位uuid
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 13:11
// @success:
func GoogleUUID20() string {
	var l = 16
	randomStr := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(randomStr) > l {
		randomStr = randomStr[:l]
	} else if len(randomStr) < l {
		randomStr = randomStr + generateRandomString(l-len(randomStr))
	}
	return randomStr + generateRandomString(4)
}

func GooGleUUID() string {
	return uuid.NewString()
}
