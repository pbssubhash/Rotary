package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	JAA "github.com/pbssubhash/JustAnotherAutomator"
)

func main() {
	template := flag.String("c", "Default", "-c ../webrecon.yml")
	threads := flag.Int("t", 10, " -t 100")
	flag.Parse()
	if *template == "Default" {
		flag.PrintDefaults()
		log.Fatalln("Required argument not present.")

	}
	file, err := os.Open(*template)
	if err != nil {
		log.Fatalln("File has issues.. Check again")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("File has issues.. Check again")
	}
	JAA.MAX = *threads
	JAA.Sem = make(chan int, *threads)

	// var wg sync.WaitGroup
	// flax := false
	// wg.Add(1)
	// go func(result JAA.ResChan) {
	// 	for x := range result {
	// 		fmt.Println("Command Output: ")
	// 		if x.StdError == "nil" {
	// 			fmt.Println("[" + x.Name + "] Output: " + x.StdOutput)
	// 		} else {
	// 			fmt.Println("[Error] [" + x.Name + "] Output: " + x.StdError)
	// 		}
	// 	}

	// }(JAA.ResultsChannel)
	JAA.ParseProfile(data)
	// fmt.Println(profile)
	JAA.WaitG()
	// wg.Wait()
	// for {
	// 	if JAA.Flax {
	// 		break
	// 	}
	// }
}
