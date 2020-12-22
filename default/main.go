package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Input default input type
// func Input() []string {
// 	var s string
// 	var arr []string
// 	for i := 0; i < 1; i++ {
// 		if sc.Scan() {
// 			s = sc.Text()
// 			arr = append(arr, s)
// 		}
// 	}
// 	return arr
// }

func run() error {
	var r io.Reader
	filename := "./test.txt"
	switch filename {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("%#v", err)
			return err
		}
		defer f.Close()
		r = f
	}
	t, err := ioutil.ReadFile(r)
	if err != nil {
		fmt.Printf("%#v", err)
		return err
	}
	fmt.Printf("%#v", t)
	return nil
}

func main() {
	run()
}
