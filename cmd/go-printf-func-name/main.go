package main

import "golang.org/x/tools/go/analysis/singlechecker"
import "github.com/Cooperjie/SastGo"

func main()  {
	singlechecker.Main(analyzer.Analyzer)
}