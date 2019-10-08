package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage int
	endPage   int

	pageLen  int
	pageType bool

	printDest string

	inFileName string
}

type spArgs selpgArgs

var progname string

func main() {
	sa := spArgs{}

	progname = os.Args[0]

	inputArgs(&sa)

	handleArgs(&sa)
	//process_input(sa)
	// fmt.Println(sa)

	process(&sa)
}

func inputArgs(sa *spArgs) {
	flag.IntVarP(&sa.startPage, "s", "s", -1, "startPage")
	flag.IntVarP(&sa.endPage, "e", "e", -1, "endPage")

	flag.IntVarP(&sa.pageLen, "l", "l", -1, "len of page")
	flag.BoolVarP(&sa.pageType, "f", "f", false, "seperated by page")

	flag.StringVarP(&sa.printDest, "d", "d", "", "printer")

	flag.Parse()

	flag.Usage = func() {
		fmt.Println(sa)
		fmt.Println(
			"USAGE: \nselpg -s startPage -e endPage [ -f | -l lines_per_page ]" +
				" [ -d dest ] [ inFileName ]\n")
	}

	if flag.NArg() == 1 {
		sa.inFileName = flag.Arg(0)
	} else {
		sa.inFileName = ""
	}

}

func handleArgs(sa *spArgs) {
	//page error
	if sa.startPage < 1 || sa.startPage > sa.endPage {
		// flag.Usage("wrong page number")
		flag.Usage()
		os.Exit(1)
	}

	//pageLen pageType error
	if sa.pageLen != -1 && sa.pageType != false {
		// flag.Usage("-f -l both exist")
		flag.Usage()
		os.Exit(2)
	}
	if sa.pageLen == -1 && sa.pageType == false {
		sa.pageLen = 72
	}

	//don't exist input file
	if sa.inFileName != "" {
		_, err := os.Stat(sa.inFileName)
		if err != nil {
			// flag.Usage("input file not exist")
			flag.Usage()
			os.Exit(3)
		}
	}

}

func process(sa *spArgs) {
	var fin *os.File
	if sa.inFileName == "" {
		fin = os.Stdin
	} else {
		var err error
		fin, err = os.Open(sa.inFileName)
		if err != nil {
			fmt.Println("can't open input file")
			os.Exit(4)
		}
		defer fin.Close()
	}
	buffer := bufio.NewReader(fin)

	var fout io.WriteCloser
	//no printer
	if sa.printDest == "" {
		fout = os.Stdout
	}

	if sa.pageType == false {
		pageCnt := 1
		lineCnt := 0
		for {
			line, crc := buffer.ReadString('\n')
			if crc != nil {
				break
			}
			lineCnt++
			if lineCnt > sa.pageLen {
				lineCnt = 1
				pageCnt++
			}
			// fmt.Println("pageCnt is %d, lineCnt is %d", pageCnt, lineCnt)
			if sa.startPage <= pageCnt && pageCnt <= sa.endPage {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Println("can't write")
					os.Exit(5)
				}
			}
		}

	} else {
		pageCnt := 0
		for {
			line, crc := buffer.ReadString('\n')
			if crc != nil {
				break
			}
			pageCnt++
			if sa.startPage <= pageCnt && pageCnt <= sa.endPage {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Println("can't write")
					os.Exit(6)
				}
			}
		}
	}

}
