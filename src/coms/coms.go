package coms

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Open xRIT file
func Init(path string) {
	fmt.Println("Opening \"" + path + "\"")

	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(data)
}
