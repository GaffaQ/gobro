package main

import (
	"fmt"
	"regexp"
)

func main() {

	regex := regexp.MustCompile(`g([a-z])a`)
	fmt.Println(regex.MatchString("gfa"))
	fmt.Println(regex.MatchString("gGa"))
	fmt.Println(regex.MatchString("gda"))

	fmt.Println(regex.FindAllString("gda gfa gca Gfa gFa gqg askdo gga ", 10))

}
