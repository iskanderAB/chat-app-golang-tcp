package main

import (
	"bufio"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	// "math/rand"
	"os"
)

func main() {
	// max := 100
	// min := 1
	// secret := rand.Intn(max - min) + min 

	// fmt.Println(secret)
	reader := bufio.NewReader(os.Stdin)

	for {
		entry , err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line := strings.Fields(entry)[0]
		num , err := strconv.Atoi(line)
		fmt.Println(reflect.TypeOf(num))
		fmt.Println(line)

	}
}