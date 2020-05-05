package main

import (
	//"flag"

	"fmt"
	"strconv"
	"strings"
	"time"

	//"log"
	//"os"
	//"os/signal"
	//"syscall"
	//"github.com/spf13/viper"
	//"github.com/panjf2000/gnet"
	"github.com/gocolly/colly"
	"github.com/golang/glog"
	jsoniter "github.com/json-iterator/go"
	mysqldata "github.com/rainbowlee/doggy/stock/sql"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type fundstock struct {
	fundid    int
	stockid   int
	stockname string
	stockper  float32
}

type colType int32

const (
	col_index colType = 0
	col_id    colType = 1
	col_name  colType = 2
	col_per   colType = 6
	col_count colType = 7
	col_total colType = 8
)

type funditem struct {
	fundid   int
	fundname string
}

type fundpagejson struct {
	datas      []string
	allRecords int
	pageIndex  int
	pageNum    int
	allPages   int
	allNum     int
	gpNum      int
	hhNum      int
	zqNum      int
	zsNum      int
	bbNum      int
	qdiiNum    int
	etfNum     int
	lofNum     int
	fofNum     int
}

//http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=519712&topline=10&year=&month=3&rt=0.794738968271171 股票持仓
//http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=zqcc&code=003886&year=&rt=0.42976462931540116 债券持仓
//
func DebugOneFundStock(fundid int, funstockitem *mysqldata.FundStockM) {
	glog.Info("fundid ", fundid, "stockid ", funstockitem.StockId, "stockname ", funstockitem.StockName, "per", funstockitem.StockPer)

	mysqldata.AddFundStock(funstockitem)

	stockm := mysqldata.StockM{}
	stockm.ID = uint(funstockitem.StockId)
	stockm.Name = funstockitem.StockName

	mysqldata.AddStock(&stockm)
}

//
func DbugOneFundZZ(fundid int, funstockitem *mysqldata.FundZqM) {
	glog.Info("fundid ", fundid, "stockid ", funstockitem.StockId, "stockname ", funstockitem.StockName, "per", funstockitem.StockPer)

	mysqldata.AddFundZZ(funstockitem)

	stockm := mysqldata.StockM{}
	stockm.ID = uint(funstockitem.StockId)
	stockm.Name = funstockitem.StockName

	mysqldata.AddStock(&stockm)
}

//
func funditemwalk(fundid int, fundname string) {
	glog.Info("funditem:", fundid)

	c := colly.NewCollector()

	var walktype int = 0
	var colcount int = 0
	fundstockitem := mysqldata.FundStockM{}
	fundzzitem := mysqldata.FundZqM{}
	c.OnRequest(func(r *colly.Request) {
		glog.Info("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		glog.Info("Something went wrong:", err)
	})

	c.OnResponseHeaders(func(r *colly.Response) {
		glog.Info("Visited header", r.Request.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		glog.Info("Visited body", string(r.Body))
	})

	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//e.Request.Visit(e.Attr("href"))
	//})
	//c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	c.OnHTML("tr td", func(e *colly.HTMLElement) {
		glog.Info("First column of a table row:", e.Text)

		if walktype == 0 {
			counttype := colType(colcount)
			if col_index == counttype {
				fundstockitem = mysqldata.FundStockM{}
			} else if col_id == counttype {
				stockid, _ := strconv.Atoi(e.Text)
				fundstockitem.StockId = stockid
			} else if col_name == counttype {
				fundstockitem.StockName = e.Text
			} else if col_per == counttype {
				per, _ := strconv.ParseFloat(strings.TrimRight(e.Text, "%"), 32)
				fundstockitem.StockPer = float32(per)
			} else if col_count == counttype {
				countstr := strings.Replace(e.Text, ",", "", -1)
				count, _ := strconv.ParseFloat(countstr, 32)
				fundstockitem.StockCount = int(count)

			} else if col_total == counttype {
				fundstockitem.FundId = fundid
				fundstockitem.FundName = fundname
				DebugOneFundStock(fundid, &fundstockitem)
			}

			colcount = colcount + 1
			colcount = colcount % (int(col_total) + 1)
		} else {
			counttype := colType(colcount)
			if 0 == counttype {
				fundzzitem = mysqldata.FundZqM{}
			} else if 1 == counttype {
				stockid, _ := strconv.Atoi(e.Text)
				fundzzitem.StockId = stockid
			} else if 2 == counttype {
				fundzzitem.StockName = e.Text
			} else if 3 == counttype {
				per, _ := strconv.ParseFloat(strings.TrimRight(e.Text, "%"), 32)
				fundzzitem.StockPer = float32(per)
			} else if 4 == counttype {
				countstr := strings.Replace(e.Text, ",", "", -1)
				count, _ := strconv.ParseFloat(countstr, 32)
				fundzzitem.StockCount = int(count)

				fundzzitem.FundId = fundid
				fundzzitem.FundName = fundname
				DbugOneFundZZ(fundid, &fundzzitem)
			} else if col_total == counttype {

			}

			colcount = colcount + 1
			colcount = colcount % (4 + 1)
		}

	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		glog.Info(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		glog.Info("Finished", r.Request.URL)
	})

	walktype = 0
	intmonth := int(time.Now().Month())
	intmonth = (intmonth - 1) / 3 * 3
	//http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=519712&topline=10&year=&month=3&rt=0.794738968271171
	stdfunditemaddr := "http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=" + strconv.Itoa(fundid) + "&topline=10&year=&month=" + strconv.Itoa(intmonth) + "&rt=0.794738968271171"

	c.Visit(stdfunditemaddr)

	walktype = 1
	colcount = 0
	stdfunditemaddrzqcc := "http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=zqcc&code=" + strconv.Itoa(fundid) + "&year=&rt=0.42976462931540116"
	c.Visit(stdfunditemaddrzqcc)
}

func fundpage(pageaddr string) {
	glog.Info("pageaddr:", pageaddr)

	c := colly.NewCollector()

	var colcount int = 0
	var fundstockitem *fundstock

	c.OnRequest(func(r *colly.Request) {
		glog.Info("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		glog.Info("Something went wrong:", err)
	})

	c.OnResponseHeaders(func(r *colly.Response) {
		glog.Info("Visited header", r.Request.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		datastr := string(r.Body)
		glog.Info("Visited response body", datastr)
		equalIndex := strings.Index(datastr, "=")
		realdatastr := datastr[equalIndex+2 : len(datastr)-1]
		glog.Info("Visited response realdatastr", realdatastr)
		//data := fundpagejson {}
		//code := json.Unmarshal([]byte(realdatastr), &data)

		//glog.Info(code)
		i := 0
		i = i + 1

		beginindex := strings.Index(datastr, "[")
		endindex := strings.Index(datastr, "]")
		if beginindex+1 == endindex {
			return
		}

		datastr2 := datastr[beginindex+2 : endindex]
		glog.Info("Visited response datastr2", datastr2)

		data2split := strings.Split(datastr2, "\",\"")
		for dataitemi := 0; dataitemi < len(data2split); dataitemi++ {
			glog.Info(data2split[dataitemi][0:])

			funditemstr := strings.Split(data2split[dataitemi][0:], ",")
			onefunditem := funditem{}
			fundid, _ := strconv.Atoi(funditemstr[0])
			onefunditem.fundid = fundid
			onefunditem.fundname = funditemstr[1]
			fund := mysqldata.FundM{}
			fund.ID = uint(fundid)
			fund.Name = onefunditem.fundname
			mysqldata.AddFund(&fund)

			funditemwalk(fundid, funditemstr[1])

			glog.Info("fundid ", onefunditem.fundid, " fundname ", onefunditem.fundname)
		}
	})

	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//e.Request.Visit(e.Attr("href"))
	//})
	//c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	c.OnHTML("tr td", func(e *colly.HTMLElement) {
		glog.Info("First column of a table row:", e.Text)

		counttype := colType(colcount)
		if col_index == counttype {
			fundstockitem = &fundstock{}
		} else if col_id == counttype {
			stockid, _ := strconv.Atoi(e.Text)
			fundstockitem.stockid = stockid
		} else if col_name == counttype {
			fundstockitem.stockname = e.Text
		} else if col_per == counttype {
			per, _ := strconv.ParseFloat(strings.TrimRight(e.Text, "%"), 32)
			fundstockitem.stockper = float32(per)
		} else if col_total == counttype {
			//DebugOneFundStock(400015, fundstockitem)
		}

		colcount = colcount + 1
		colcount = colcount % (int(col_total) + 1)
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		glog.Info(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		glog.Info("Finished", r.Request.URL)
	})

	c.Visit(pageaddr)
}

func FundIter() {
	//http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=zzf&st=desc&sd=2019-05-04&ed=2020-05-04&qdii=&tabSubtype=,,,,,&pi=1&pn=50&dx=1&v=0.18916728629687607

	//funditemwalk(3886, "test111")

	now := time.Now().Add(-24 * 60 * 60 * 1e9)
	month := now.Month()
	day := now.Day()

	yearmonthday := fmt.Sprintf("%d-%02d-%02d", now.Year(), month, day)

	for i := 0; i < 1000; i++ {
		var netaddr = "http://fund.eastmoney.com/data/rankhandler.aspx?op=ph&dt=kf&ft=all&rs=&gs=0&sc=zzf&" +
			"st=desc&sd=" + yearmonthday + "&ed=" + yearmonthday + "&qdii=&tabSubtype=,,,,,&pi=" + strconv.Itoa(i+1) + "&pn=50&dx=1&v=0.18916728629687607"
		fundpage(netaddr)

		glog.Info("walk netaddr", netaddr)
	}

}

// walk get fund info
func FundWalk() {
	c := colly.NewCollector()

	var colcount int = 0
	var fundstockitem *fundstock

	c.OnRequest(func(r *colly.Request) {
		glog.Info("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		glog.Info("Something went wrong:", err)
	})

	c.OnResponseHeaders(func(r *colly.Response) {
		glog.Info("Visited header", r.Request.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		glog.Info("Visited response", r.Request.URL)
	})

	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//e.Request.Visit(e.Attr("href"))
	//})
	//c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	c.OnHTML("tr td", func(e *colly.HTMLElement) {
		glog.Info("First column of a table row:", e.Text)

		counttype := colType(colcount)
		if col_index == counttype {
			fundstockitem = &fundstock{}
		} else if col_id == counttype {
			stockid, _ := strconv.Atoi(e.Text)
			fundstockitem.stockid = stockid
		} else if col_name == counttype {
			fundstockitem.stockname = e.Text
		} else if col_per == counttype {
			per, _ := strconv.ParseFloat(strings.TrimRight(e.Text, "%"), 32)
			fundstockitem.stockper = float32(per)
		} else if col_total == counttype {
			//DebugOneFundStock(400015, fundstockitem)
		}

		colcount = colcount + 1
		colcount = colcount % (int(col_total) + 1)
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		glog.Info(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		glog.Info("Finished", r.Request.URL)
	})

	c.Visit("http://fundf10.eastmoney.com/FundArchivesDatas.aspx?type=jjcc&code=400015&topline=100&year=&month=0&rt=0.4153446579006561")
}
