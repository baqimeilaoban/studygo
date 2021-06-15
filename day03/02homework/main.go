package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//1.判断字符串中汉字的数量
	//难点：判断一个字符是汉字
	s1 := "hello沙河"
	//1.1依次拿到字符串的字符
	var count int
	for _, c := range s1 {
		//1.2判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	//1.3把汉字出现的次数累加得到最终结果
	fmt.Println(count)

	//2.how do you do单词出现的次数
	//2.1把字符串按照空格切割得到切片
	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	//2.2遍历切片存储到一个map
	for _, w := range s3{
		if _, ok := m1[w]; !ok{
			//1.如果原来map中不存在w这个key那么次数=1
			m1[w] = 1
		}else {
			//2.如果map中存在w这个key，那么次数+1
			m1[w]++
		}
	}
	//2.3累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	//3.回文判断：字符串从左往右和从右往左读都是一样的，那么就是回文
	//上海自来水来自海上
	ss := "山西运煤车煤运西山"
	//把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, v := range ss {
		r = append(r, v)
	}
	for i := 0; i < len(r)/2; i++ {
		//山 ss[0] ss[len(ss)-1]
		//西 ss[1] ss[len(ss)-1-1]
		//运 ss[2] ss[len(ss)-1-2]
		//...
		//c ss[i] ss[len(ss)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
