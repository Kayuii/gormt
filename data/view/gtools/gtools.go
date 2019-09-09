package gtools

import (
	"fmt"
	"os/exec"

	"github.com/kayuii/gormt/data/config"

	"github.com/xxjwxc/public/mysqldb"
	"github.com/xxjwxc/public/tools"
)

//Execute 开始执行
func Execute() {

	orm := mysqldb.OnInitDBOrm(config.GetMysqlConStr())
	defer orm.OnDestoryDB()

	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)

	pkg := OnGetPackageInfo(orm, OnGetTables(orm))
	pkg.SetPackage("model")
	str := pkg.Generate()

	path := config.GetOutDir() + "/" + config.GetMysqlDbInfo().Database + ".go"
	tools.WriteFile(path,
		[]string{str}, true)

	fmt.Println("formatting differs from goimport's:")
	cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
	fmt.Println(string(cmd))

	fmt.Println("formatting differs from gofmt's:")
	cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
	fmt.Println(string(cmd))

}
