package mysqldata

import(
    _  "github.com/go-sql-driver/mysql" // 引入包，不使用，使其调用init函数注册mysql
    "database/sql"
    "fmt"
)

var (
	MysqlDB sql.DB
)

func ConnectDB() bool {
    MysqlDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stock?charset=utf8mb4")
    if err != nil {
        fmt.Println("创建数据库对象失败")
        return false
    }
    // 实际去尝试连接数据库
    err = MysqlDB.Ping()

    if err != nil {
        fmt.Println("连接数据库失败")
        return false
    }

    fmt.Println("连接数据库成功")

	return true;
}
