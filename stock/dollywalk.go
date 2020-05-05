package main
/*
import (
	//"flag"
	"fmt"
	"strconv"
	"strings"
	//"log"
	//"os"
	//"os/signal"
    //"syscall"
	//"github.com/spf13/viper"
	//"github.com/panjf2000/gnet"	
	"github.com/gocolly/colly"
)

type  fundstock struct{
	fundid      int
	stockid 	int
	stockname	string
	stockper	float32
}

type colType int32

const (
    col_index     	colType = 0
    col_id      	colType = 1
    col_name     	colType = 2
	col_per      	colType = 6
	col_total		colType = 8
)

func OneFundStock(fundid int,funstockitem *fundstock)  {
	fmt.Println("fundid ", fundid , "stockid ", funstockitem.stockid, "stockname ", funstockitem.stockname, "per", funstockitem.stockper)
}

// walk test
func DollyWalk() {
	c := colly.NewCollector()

	var colcount int = 0
	var fundstockitem *fundstock

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	
	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("Visited header", r.Request.URL)
	})
	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited response", r.Request.URL)
	})
	
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
	//})
	//c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	c.OnHTML("tr td", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
		
		counttype := colType(colcount)
		if col_index == counttype {	
			fundstockitem = &fundstock{}
		} else if col_id == counttype {
			stockid,_:=strconv.Atoi(e.Text)
			fundstockitem.stockid = stockid
		} else if col_name == counttype {
			fundstockitem.stockname = e.Text
		} else if col_per == counttype {
			per,_:=strconv.ParseFloat(strings.TrimRight(e.Text,"%"), 32)			
			fundstockitem.stockper = float32(per)
		} else if col_total == counttype {
			OneFundStock(400015,fundstockitem)
		}
		
		colcount= colcount + 1
		colcount= colcount % (int(col_total)+1)
	})
	
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=400015&topline=100&year=&month=0&rt=0.4153446579006561")
}
*/