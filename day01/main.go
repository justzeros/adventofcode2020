package main

import (
		  "fmt"
//		  "io"
			"strings"
			"strconv"	
		  "io/ioutil"
			"sort"
)

func main() {
	dat := read_file()
	for _,i := range dat {
		for _,j := range dat {
			for _,k := range dat {
				if (i + j + k) == 2020 {
					fmt.Println("Found product: ", i* j* k)
				}
			}		
		}
	}
}

func find_match(val int, dat []int) bool {
	for _,i := range dat {
		if i == val {return true}		
	}
	return false
}

func read_file() (report_values []int) {
	dat,err := ioutil.ReadFile("./expensereport.txt")
	if err != nil {
					panic(err)
	}

	for _,i := range strings.Split(string(dat),"\n") {
					if i == "" {continue}
					j, err := strconv.Atoi(i)
					if err != nil {
									panic(err)
					}
					report_values = append(report_values,j)
	}
	sort.Ints(report_values)
	return
}
