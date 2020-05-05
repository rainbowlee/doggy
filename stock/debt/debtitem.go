package debt 

import (
	"fmt"
)


type DebtItem struct {
	//Id				int
	DebtName		string
	DebtPrice		float32
	StockName		string
	StockPrice		float32
}

func (debtitem *DebtItem) Estimate(curstockprice float32, debtperstock float32) {
	var stockcount = 1000.0/debtperstock;
	var stockvalue = 100/debtitem.DebtPrice*curstockprice
	var stockmoney = stockcount*curstockprice

	fmt.Println( "need stock count ", stockcount, " stockmony " , stockmoney, " stockvalue value ", stockvalue)
}

func (debtitem *DebtItem) ToString() {
	str :=fmt.Sprintf("debtname %s debtprice %f stockname %s stockprice %f",debtitem.DebtName,debtitem.DebtPrice, debtitem.StockName,debtitem.StockPrice)
	fmt.Println(str)
}

func NewDebt(debtName string, debtprice float32, stockname string, stockprice float32) DebtItem {  
    e := DebtItem {debtName, debtprice, stockname, stockprice}
    return e
}
