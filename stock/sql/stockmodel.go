package mysqldata

import (
	_ "github.com/go-sql-driver/mysql" // 引入包，不使用，使其调用init函数注册mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StockM struct {
	gorm.Model
	Name   string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
	Region int
}

type FundM struct {
	gorm.Model
	Name string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。

	StockM []StockM `gorm:"-"`
}

type FundStockM struct {
	FundId  int
	StockId int

	StockCount int
	StockPer   float32

	StockName string
	FundName  string
}

type FundZqM struct {
	FundId  int
	StockId int

	StockCount int
	StockPer   float32

	StockName string
	FundName  string
}
