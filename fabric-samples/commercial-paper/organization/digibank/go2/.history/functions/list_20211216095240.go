package functions

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ListAll() error {
	files, err := ioutil.ReadDir("../identity/user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("LIST all user:")

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return err
}

// err = filepath.Walk("../identity/",
// 	func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(path)
// 		fmt.Println("Name:", info.Name())

// 		return nil
// 	})
// if err != nil {
// 	log.Println(err)
// }
