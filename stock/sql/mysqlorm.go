package mysqldata

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 引入包，不使用，使其调用init函数注册mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	OrmMysqlDB *gorm.DB
)

//
func OrmConnectDB() bool {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stock?charset=utf8mb4")
	OrmMysqlDB = db
	//MysqlDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stock?charset=utf8mb4")
	if err != nil {
		fmt.Println("创建数据库对象失败")
		return false
	}
	registermodel()
	fmt.Println("连接数据库成功")
	return true
}

//
func OrmCloseDB() {
	if OrmMysqlDB != nil {
		OrmMysqlDB.Close()
	}

	OrmMysqlDB = nil
}

func registermodel() {
	//OrmMysqlDB.CreateTable(&FundStock{})
	//OrmMysqlDB.DropTable(&FundStock{})
	OrmMysqlDB.AutoMigrate(&StockM{})
	OrmMysqlDB.AutoMigrate(&FundM{})
	OrmMysqlDB.AutoMigrate(&FundStockM{})
	OrmMysqlDB.AutoMigrate(&FundZqM{})

	OrmMysqlDB.Model(&FundStockM{}).AddUniqueIndex("idx_fund_stock", "fund_id", "stock_id")
	OrmMysqlDB.Model(&FundZqM{}).AddUniqueIndex("idx_fund_stock", "fund_id", "stock_id")

	stock := &StockM{}
	stock.ID = 2
	stock.Name = "stock2"
	AddStock(stock)
	OrmMysqlDB.Delete(stock)

	fund := &FundM{}
	fund.ID = 2
	fund.Name = "fund311"
	AddFund(fund)

	fundstock := &FundStockM{}
	fundstock.StockId = 1
	fundstock.FundId = 1
	AddFundStock(fundstock)
}

//
func AddStock(stock *StockM) {
	findfund := StockM{}
	OrmMysqlDB.First(&findfund, stock.ID)
	if findfund.ID > 0 {
		OrmMysqlDB.Model(stock).Update()
	} else {
		OrmMysqlDB.Create(stock)
	}
}

//
func AddFund(fund *FundM) {
	findfund := FundM{}
	OrmMysqlDB.First(&findfund, fund.ID)
	if findfund.ID > 0 {
		OrmMysqlDB.Model(fund).Update()
	} else {
		OrmMysqlDB.Create(fund)
	}
}

//
func AddFundStock(fundstock *FundStockM) {
	findfund := FundStockM{}
	OrmMysqlDB.Where("fund_id = ? AND stock_id = ?", fundstock.FundId, fundstock.StockId).First(&findfund)
	if findfund.FundId > 0 {
		//OrmMysqlDB.Model(fund).Update()
	} else {
		OrmMysqlDB.Create(fundstock)
	}
}

//
func AddFundZZ(fundzz *FundZqM) {
	findfund := FundZqM{}
	OrmMysqlDB.Where("fund_id = ? AND stock_id = ?", fundzz.FundId, fundzz.StockId).First(&findfund)
	if findfund.FundId > 0 {
		//OrmMysqlDB.Model(fund).Update()
	} else {
		OrmMysqlDB.Create(fundzz)
	}
}
