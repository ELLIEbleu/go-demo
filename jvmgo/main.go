package main

import (
	"flag"
	"fmt"
	"os"
	"rsc.io/quote"
	"strings"
	"test/jvmgo/classpath"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	XjreOption  string
	Class       string
	Args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage:.%s [-options] class [args...]\n", os.Args[0])
}

func main() {
	//cmd := parseCmd()
	//
	//if cmd.VersionFlag {
	//	fmt.Println("version 0.0.1")
	//} else if cmd.HelpFlag || cmd.Class == "" {
	//	printUsage()
	//} else {
	//	startJVM(cmd)
	//}
	fmt.Printf(quote.Go())
	//fmt.Println(1<<64)        ### constant 18446744073709551616 overflows int
	var u uint
	fmt.Println(u)
	u--
	fmt.Println(u)
	//hey(4)



}


func hey(i int) {
	if i <= 0 {
		panic("panicked on hey")
	}
	there(i - 1)
}

func there(i int) {
	if i <= 0 {
		panic("panicked on hey")
	}
	hey(i - 1)
}

func withMap(m map[string]int, i int) {
	if i&1 == 0 {
		m["key"] = 10
	} else {
		v := m["key"]
		if v == 0 {
			m["key"] = 11
		}
	}

}

func test()  {

}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cp, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("read class faild, %s\n", cmd.Class)
	}
	fmt.Printf("class data:%v\n", classData)
}

//version1

//func startJVM(cmd *Cmd) {
//	fmt.Printf("classpath:%s class:%s args:%v\n",
//		cmd.CpOption, cmd.Class, cmd.Args)
//
//}
