package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

// Config 结构体嵌套
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0 参数的校验
	// 0.1 传进来的参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		// 新创建一个错误
		err = errors.New("data param should be a pointer")
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针(因为配置文件中各种键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}
	// 1 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.New("failed to read file")
		return
	}
	// 将字节类型的文件内容转换为字符串
	lineSlice := strings.Split(string(b), "\n")
	// 2 一行一行的读取数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 2.2 如果是[开头的就表示是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line%d syntax error", idx+1)
				return
			}
			// 把这一行的首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是开头就是=分割的键值对
			// 2.3.1 以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2.3.2 根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			// 拿到嵌套结构体的值信息
			sValue := v.Elem().FieldByName(structName)
			// 拿到嵌套结构体的类型信息
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 2.3.3 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				// tag信息是存储在类型信息中的
				filed := sType.Field(i)
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			// 2.3.4 如果key=tag，给这个字段赋值
			// 2.3.4.1 根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 2.3.4.2 对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("行%d语法错误，传入参数有误，需要为整形", idx+1)
					return err
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("行%d语法错误，传入参数有误，需要为布尔型", idx+1)
					return err
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("行%d语法错误，传入参数有误，需要为浮点型", idx+1)
					return err
				}
				fileObj.SetFloat(valueFloat)
			}
		}

	}

	return nil
}

func main() {
	var cfg Config
	err := loadIni("day07/ini_demo/conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
