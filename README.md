# genpw
随机密码生成工具

## 安装

```console
$ go get github.com/dzhcool/genpw
$ go install github.com/dzhcool/genpw
```

## 使用
usage: ./genpw -level=4 -len=14 -num=1 -spec="!@#$%^&"

        -level 	日志级别，1：字母 2：字母、数字 3：字母、数字、特殊字符随机 4：复杂混合(缺省值)

        -len   	密码长度，建议至少6位 缺省值：14

        -num   	生成个数，缺省值：1

        -spec  	包含特殊字符限定，缺省值：_+-&=@#$%^*)(
