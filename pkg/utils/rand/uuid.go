package rand

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/gjing1st/hertz-admin/pkg/utils/slice"
	"github.com/google/uuid"
	"strings"
	"time"
)

// GenerateUUID32
// @description: 生成uid 32位
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 11:37
// @success:
func GenerateUUID32() string {
	now := time.Now().UnixNano()
	timeByte := make([]byte, 8)
	binary.BigEndian.PutUint64(timeByte, uint64(now))
	nowStr := hex.EncodeToString(timeByte)
	return nowStr + S(16, false)
}

// GoogleUUID32
// @description: google/uuid
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 13:08
// @success:
func GoogleUUID32() string {
	var l = 32
	uuid := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(uuid) > l {
		uuid = uuid[:l]
	} else if len(uuid) < l {
		uuid = uuid + S(l-len(uuid), false)
	}
	return uuid
}

// GenerateUUID20
// @description: 20位uuid
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 11:39
// @success:
func GenerateUUID20() string {
	now := time.Now().UnixNano()
	timeByte := make([]byte, 8)
	binary.BigEndian.PutUint64(timeByte, uint64(now))
	nowStr := hex.EncodeToString(timeByte)
	newNowStr := slice.ReverseString(nowStr[2:])
	return S(6, false) + newNowStr
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
	uuid := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(uuid) > l {
		uuid = uuid[:l]
	} else if len(uuid) < l {
		uuid = uuid + S(l-len(uuid), false)
	}
	return uuid + S(4, false)
}

func GooGleUUID() string {
	return uuid.NewString()
}

func GenerateRuleName() string {
	now := time.Now().UnixNano()
	timeByte := make([]byte, 8)
	binary.BigEndian.PutUint64(timeByte, uint64(now))
	nowStr := hex.EncodeToString(timeByte)
	newNowStr := slice.ReverseString(nowStr[2:])
	return LS(6) + newNowStr
}

// GenerateUUid
// @Description 生成指定位数的uuid
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/7/19 16:12
func GenerateUUid(l int) string {
	str := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(str) > l {
		str = str[:l]
	} else if len(str) < l {
		str = str + S(l-len(str), false)
	}
	return str
}
