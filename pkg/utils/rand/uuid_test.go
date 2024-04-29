package rand

import (
	"fmt"
	"github.com/gjing1st/hertz-admin/pkg/utils"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestUid(t *testing.T) {
	for i := 0; i < 5; i++ {
		//s := GenerateUUID20()
		//fmt.Println("len=", len(s), "===", s)
		//time.Sleep(time.Nanosecond)
		id := uuid.New()
		fmt.Println(id.String(), "============", len(id.String()), id.Version())
	}
	//fmt.Println(time.Now().UnixNano())

	id := uuid.New()
	ids := strings.ReplaceAll(id.String(), "-", "")
	fmt.Println(ids, "============", len(ids), id.Version())

}

func TestName(t *testing.T) {
	fmt.Println("------")
	var dest []map[string]interface{}
	dest = make([]map[string]interface{}, 4)
	ss1, ss2, ss3, ss4 := make(map[string]interface{}, 1), make(map[string]interface{}, 1), make(map[string]interface{}, 1), make(map[string]interface{}, 1)
	ss1["db_name"] = "mysql"
	dest[0] = ss1
	ss2["db_name"] = "gaf"
	dest[1] = ss2
	ss3["db_name"] = "sys"
	dest[2] = ss3
	ss4["db_name"] = "information_schema"
	dest[3] = ss4
	//dest[2]["db_name"] = "sys"
	fmt.Println("--------dest", dest)
	//var dest1 = make([]map[string]interface{}, 0)
	for i := 0; i < len(dest)-1; i++ {
		fmt.Println("---------------------------------------------------")
		dbName := utils.String(dest[i]["db_name"])
		if dbName == "mysql" || dbName == "information_schema" || dbName == "performance_schema" || dbName == "sys" {
			dest = append(dest[:i], dest[i+1:]...)
			i--
		}
		//dest1 = append(dest1, dest[i])

	}
	fmt.Println("res===", dest)
}
