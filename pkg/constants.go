package JAA

import (
	"fmt"
	"os"
	"sync"
)

type Asset struct {
	Host    string
	IP      string
	Details string
}

type Assets []Asset

var AssetList Assets

type DnsJson struct {
	Status   int  `json:"Status"`
	TC       bool `json:"TC"`
	RD       bool `json:"RD"`
	RA       bool `json:"RA"`
	AD       bool `json:"AD"`
	CD       bool `json:"CD"`
	Question []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
}

// func DOHIP(domain string) []string {
// 	req, err := http.NewRequest("GET", "https://cloudflare-dns.com/dns-query", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	q := req.URL.Query()
// 	q.Add("ct", "application/dns-json")
// 	q.Add("name", domain)
// 	req.URL.RawQuery = q.Encode()
// 	client := &http.Client{}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	res, err := client.Do(req)
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 	}
// 	var dnsResult DnsJson
// 	if err := json.Unmarshal(body, &dnsResult); err != nil {
// 		log.Fatal(err)
// 	}
// 	var temp []string
// 	for _, dnsAnswer := range dnsResult.Answer {
// 		temp = append(temp, dnsAnswer.Data)
// 	}
// 	if len(temp) == 0 {
// 		return []string{"NA"}
// 	}
// 	return temp
// }

type Result struct {
	Name      string
	StdOutput string
	StdError  string
}
type ResChan chan Result

var ResultsChannel = make(chan Result)

type Command struct {
	Name       string   `yaml:"name"`
	Concurrent bool     `yaml:"concurrent"`
	Commands   []string `yaml:"commands"`
}

type Profile struct {
	APIVersion float64 `yaml:"apiVersion" json:"apiVersion"`
	BasicInfo  struct {
		Author   string `yaml:"author" json:"author"`
		Name     string `yaml:"name" json:"name"`
		TargetOs string `yaml:"targetOs" json:"targetOS"`
	} `yaml:"basicInfo" json:"basicInfo"`
	UserInputs map[string]string `yaml:"userInputs"`
	Concurrent bool              `yaml:"concurrent"`
	Commands   []Command         `yaml:"commands"`
}

type Job struct {
	Name     string
	Async    bool
	Commands []string
	Status   string
}

var Wg sync.WaitGroup
var MAX int
var Sem chan int
var Flax bool

// Functions..
func WaitG() {
	Flax = false
	Wg.Wait()
	Flax = true
}

// return Profile{BasicInfo: t.BasicInfo, Concurrent: t.Concurrent, Commands: Xcommands}

// func ParseChunks(t Profile) []Command {
// 	var Commands []Command
// 	var userInputs map[string]string
// 	if t.Concurrent {
// 		Wg.Wait()
// 		go func(t Profile) {
// 			x := AskandParseInputs(t)
// 		}(t)
// 		Wg.Done()
// 	} else {

// 	}
// 	for _, v := range t.Commands {
// 		// for v.Commands
// 	}
// }

// func AskInputs(t Profile) []Command {
// 	var Commands []Command
// 	for key, value := range t.UserInputs {
// 		fmt.Println("[User-Input] User Input required for " + key)
// 		fmt.Scanln(&value)
// 	}
// 	for _, cmd := range t.Commands {
// 		cmdx := strings.ReplaceAll(cmd, "%{"+key+"}", value)
// 		Commands := append(Commands, Command{Name: t.BasicInfo.Name, Command: cmdxs, TargetOS: t.BasicInfo.TargetOS,Async:})
// 	}
// 	return Commands
// }
func CheckError(err error) {
	fmt.Println("Error occured.")
	fmt.Println(err.Error())
	os.Exit(1)
}
