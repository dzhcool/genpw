package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	flag_level int64  // 密码复杂度
	flag_len   int64  // 长度
	flag_num   int64  // 生成个数
	flag_spec  string // 特殊字符定义
	flag_help  bool   // 帮助信息
)

// 定义密码级别
const (
	LevelChar    = 1 // 字母
	LevelCNMix   = 2 // 字母、数字
	levelCNSMix  = 3 // 字母、数字、特殊字符
	levelAdvance = 4 // 更好的生成,保证每种类型都出现，缺省类型
)

// 定义默认字符
const (
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NumStr  = "0123456789"
	SpecStr = "_+-&=@#$%^*()"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	initFlag()

	if flag_help {
		usage()
		return
	}

	stime := time.Now()
	for i := int64(0); i < flag_num; i++ {
		passwd := genPasswd()
		fmt.Println(string(passwd))
	}

	etime := time.Now()
	uptime := etime.Sub(stime)
	fmt.Println("use:", uptime)
}

// 初始化flag
func initFlag() {
	flag.Int64Var(&flag_level, "level", 4, "请输入生成等级(1:字母 2:字母、数字 3:字母、数字、特殊字符随机) 4:复杂混合")
	flag.Int64Var(&flag_len, "len", 14, "请输入生成密码长度(至少6位)")
	flag.Int64Var(&flag_num, "num", 1, "请输入生成密码个数")
	flag.StringVar(&flag_spec, "spec", SpecStr, "可选，定制可包含特殊字符")

	flag.BoolVar(&flag_help, "h", false, "this help")
	flag.BoolVar(&flag_help, "help", false, "this help")
	flag.Parse()

	if flag_len < 6 {
		flag_len = 6
	}
}

// genPasswd
func genPasswd() []byte {
	var passwd []byte

	if flag_level == LevelChar {
		passwd = genPwd(CharStr)
	} else if flag_level == LevelCNMix {
		passwd = genPwd(fmt.Sprintf("%s%s", NumStr, CharStr))
	} else if flag_level == levelCNSMix {
		passwd = genPwd(fmt.Sprintf("%s%s%s", NumStr, CharStr, flag_spec))
	} else {
		passwd = genAdvancePwd()
	}
	return passwd
}

// 生成简单密码
func genPwd(sourceStr string) []byte {
	var passwd []byte = make([]byte, flag_len, flag_len)

	for i := int64(0); i < flag_len; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return passwd
}

// 生成复杂密码
// 特殊字符2位即可，位置随机
// 数字占比30%取整个
func genAdvancePwd() []byte {
	var passwd []byte = make([]byte, flag_len, flag_len)

	// 计算各个类型比例
	var specNum int64 = rand.Int63n(2) + 1
	var ratio float64 = float64(rand.Intn(40) / 100) // 数字占比
	var numericNum int64 = int64(math.Ceil(float64(flag_len-specNum)*ratio) + 1)
	var charNum int64 = flag_len - specNum - numericNum

	for i := int64(0); i < flag_len; i++ {
		var sourceStr string
		sourceStr, charNum, numericNum, specNum = getAdvanceSource(charNum, numericNum, specNum)
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return passwd
}

// 复杂密码随机类型
func getAdvanceSource(charNum, numericNum, specNum int64) (string, int64, int64, int64) {
	typeArr := make([]string, 0, charNum+numericNum+specNum)

	if charNum > 0 {
		for i := int64(0); i < charNum; i++ {
			typeArr = append(typeArr, "c")
		}
	}
	if numericNum > 0 {
		for i := int64(0); i < numericNum; i++ {
			typeArr = append(typeArr, "n")
		}
	}
	if specNum > 0 {
		for i := int64(0); i < specNum; i++ {
			typeArr = append(typeArr, "s")
		}
	}
	if len(typeArr) <= 0 {
		return "", 0, 0, 0
	}
	sourceType := typeArr[rand.Intn(len(typeArr))]

	sourceStr := ""
	switch sourceType {
	case "c":
		sourceStr = CharStr
		charNum--
	case "n":
		sourceStr = NumStr
		numericNum--
	case "s":
		sourceStr = flag_spec
		specNum--
	default:
		sourceStr = CharStr
	}

	return sourceStr, charNum, numericNum, specNum
}

func usage() {
	fmt.Println(`usage: ./genpw -level=4 -len=14 -num=1 -spec="!@#$%^&"`)
	fmt.Println("\t -level \t日志级别，1：字母 2：字母、数字 3：字母、数字、特殊字符随机 4：复杂混合(缺省值)")
	fmt.Println("\t -len   \t密码长度，建议至少6位 缺省值：14")
	fmt.Println("\t -num   \t生成个数，缺省值：1")
	fmt.Println("\t -spec  \t包含特殊字符限定，缺省值：_+-&=@#$%^*)(")
	fmt.Println("")
}
