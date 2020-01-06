package main

import (
	"time"
	"fmt"
	"strings"
	"github.com/spf13/viper"
)

const(
	AddNewUser = "addnewuser" // cmd id
	AddUser = "adduser"	// cmd id
	ReCal	= "recal"	// cmd datekey(2019_1_3)
)

func recoverfromerror() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

func GetDateKey(curtime *time.Time ) string{
	var t time.Time
	if curtime != nil {
		t = *curtime
	}else{
		t = time.Now()
	}
	str := fmt.Sprintf("%d_%02d_%02d", t.Year(), t.Month(), t.Day())
	return str
}


func GetAddNewUserKey() string{
	return "addnewuser"
}

func GetAddUserKey() string{
	return "adduser"
}

func TrimLineData(lineData string) string{
	index1 := strings.Index(lineData, "\r")
	index2 := strings.Index(lineData, "\n")
	if index1 < 0{
		index1 = index2
	}

	return lineData[:index1]
}

// dipatch line 
func Dispatch(line string){
	defer recoverfromerror()

	dateKey := GetDateKey(nil)
	fmt.Println("curDateKey", dateKey)

	line = TrimLineData(line)
	datas := strings.Split(line, " ")
	switch datas[0] {
	case AddNewUser:
		addnewuserkey := GetAddNewUserKey()
		adduserkey := GetAddUserKey()

		userId := datas[1]

		conn := RedisPool.Get()
		defer conn.Close()

		value, ok := conn.Do("sadd", dateKey + "_" + addnewuserkey, userId)
		if ok != nil {
			fmt.Println( ok, value)
		}
		value, ok = conn.Do("bf.add", dateKey + "_" + adduserkey, userId)	
		if ok != nil {
			fmt.Println( ok )
		}
/*
		for i :=1; i < 1000; i++{
			value, ok := conn.Do("sadd", dateKey + "_" + addnewuserkey, 9000+ i)
			if ok != nil {
				fmt.Println( ok, value)
			}
			value, ok = conn.Do("bf.add", dateKey + "_" + adduserkey, 9000+ i)
			if ok != nil{
				fmt.Println( ok )
			}
		}
		*/
/*
		var offset uint64 = 0

		for{
			pagedata, cursor, error := RedisClient.SScan(dateKey + "_" + addnewuserkey, offset, "*", 10).Result()

			//offset += (uint64)(len(pagedata)) + 1
			offset = cursor
			fmt.Println(pagedata, cursor, error)	
			if error != nil || len(pagedata) == 0 || cursor == 0{
				break
			}
		}
*/
		// 正确的使用方式
		//value, error := conn.Do("sscan", dateKey + "_" + addnewuserkey, 0, "match", "*", "count", "100")
		//fmt.Println(value, error)	

		//value, ok = conn.Do("bf.exists", dateKey + "_" + adduserkey, userId)	
		//if ok != nil {
		//	fmt.Println( ok )
		//}

		//_, ok := RedisClient.SAdd(dateKey + "_" + addnewuserkey, userId).Result()
		//_, ok = RedisClient.SAdd(dateKey + "_" + adduserkey, userId).Result()
	case AddUser:
		adduserkey := GetAddUserKey()

		userId := datas[1]
		conn := RedisPool.Get()
		defer conn.Close()
		value, ok := conn.Do("bf.add", dateKey + "_" + adduserkey, userId)	
		if ok != nil {
			fmt.Println( ok,value )
		}
		//_, ok := RedisClient.PFAdd(dateKey + "_" + adduserkey, userId).Result()
		//if ok == nil {

		//}	
	case ReCal:
		datekey := datas[1]
		BoottimeSettlement(datekey)
	}
}



//定时计算留存存库
func BoottimeTimingSettlement() {
    for {
        now := time.Now()
        // 计算下一个3点
        next := now.Add(time.Hour * 24 + 3)
        next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
        t := time.NewTimer(next.Sub(now))
        <-t.C
		fmt.Printf("定时结算Boottime表数据，结算完成: %v\n",time.Now())
		
        //以下为定时执行的操作
        BoottimeSettlement(GetDateKey(&now))
    }
}

//计算留存率
func BoottimeSettlement(datekey string){
    localTime, err := time.ParseInLocation("2006_01_02", datekey, time.Local)
    if err != nil{
        fmt.Println(err)
        return
	}
	
	datestarttime := viper.GetString("Default.Process_DateStart")

	daytime, _ := time.ParseDuration("-24h")
	yesterday := localTime.Add(daytime)
	beforeyesterday := localTime.Add(daytime*2)

	yesterdaykey := GetDateKey(&yesterday)
	beforeyesterdaykey := GetDateKey(&beforeyesterday)
	_= beforeyesterdaykey

	conn := RedisPool.Get()
	defer conn.Close()
	var offset uint64 = 0
	var relogincount int = 0
	var allcount int  = 0
	for{
		pagedata, cursor, error := RedisClient.SScan(beforeyesterdaykey + "_" + GetAddNewUserKey(), offset, "*", 10).Result()
		for _, value := range pagedata {
			r, _ := conn.Do("bf.exists", yesterdaykey + "_" + GetAddUserKey(), value)
			//fmt.Println(r, error, value)
			if r.(int64) == 1 {
				relogincount = relogincount + 1
			}
		}

		allcount += len(pagedata)

		offset = cursor
		fmt.Println(pagedata, cursor, error)	
		if error != nil || len(pagedata) == 0 || cursor == 0{
			break
		}
	}

	fmt.Println(" next day login ", relogincount, allcount)
	//计算前天的次留
	fmt.Println(datestarttime)
	fmt.Println(localTime)
	fmt.Println(yesterday)
	fmt.Println(beforeyesterday)
}