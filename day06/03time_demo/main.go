package main

import (
	"fmt"
	"time"
)

//时间

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//time.Unix
	ret := time.Unix(1625758751, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	//时间间隔
	fmt.Println(time.Second)
	//now + 24小时
	fmt.Println(now.Add(24 * time.Hour))
	////定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t) //每秒钟执行一次
	//}

	//格式化时间 把语言中的时间对象 转换为字符串类型的时间
	//2019-08-03
	fmt.Println(now.Format("2006-01-02"))
	//2009/02/03 11:55:02
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//2009/02/03 11:55:02 AM
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))

	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2019-08-03")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

//时区
func f2() {
	now := time.Now() //本地时间
	fmt.Println(now)
	//明天的这个时间
	//按照指定格解析一个字符串格式的时间
	//time.Parse("2006-01-02 15:04:05", "2021-07-12 23:07:32")
	//按照东八区的时区和格式去解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
		return
	}
	//按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-07-12 23:07:32", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	//时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	//f1()
	f2()
}
