package aaa

import "fmt"

//bbbb
var Dbconn string

func init() {
	Dbconn = "mysql"
}

//lk ifc
func AddPeter() {
	fmt.Println("inner...............")
}
