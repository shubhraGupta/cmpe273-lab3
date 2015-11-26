package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Data struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var serv [3]int

func putData(portno string, key string, val string) {

	url := "http://localhost:" + portno + "/keys/" + key + "/" + val
	//http://localhost:3002/keys/1
	fmt.Println(url)

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, nil)

	resp, _ := client.Do(req)
	fmt.Println(resp.StatusCode)

}

func getData(portno string, key string) Data {

	url := "http://localhost:" + portno + "/keys/" + key

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var f Data
	//var f F

	err = json.Unmarshal(body, &f)
	if err != nil {
		panic(err)
	}

	//fmt.Println(f)
	return f

}

func main() {

	portno := []string{"3000", "3001", "3002"}

	hash := crc32.ChecksumIEEE
	i := 0

	for i = 0; i < 3; i++ {
		keyHash := int(hash([]byte(portno[i])))
		fmt.Println(keyHash, "\t")
		serv[i] = keyHash + 50000000
	}

	pairs := map[string]string{
		"1":  "a",
		"2":  "b",
		"3":  "c",
		"4":  "d",
		"5":  "e",
		"6":  "f",
		"7":  "g",
		"8":  "h",
		"9":  "i",
		"10": "j",
	}

	for i = 1; i <= 10; i++ {

		a := strconv.Itoa(i)
		keyHash := int(hash([]byte(a)))
		fmt.Println(keyHash, "\t", i)

		val, _ := pairs[a]

		if keyHash > serv[0] && keyHash < serv[1] {
			putData(portno[1], a, val)
		} else if keyHash > serv[1] && keyHash < serv[2] {
			putData(portno[2], a, val)
		} else {
			putData(portno[0], a, val)
		}
	}

	fmt.Println()
	fmt.Println("Now getting Data")
	u := Data{}

	for i = 1; i <= 10; i++ {
		a := strconv.Itoa(i)
		keyHash := int(hash([]byte(a)))

		if keyHash > serv[0] && keyHash <= serv[1] {
			u = getData(portno[1], a)
		} else if keyHash > serv[1] && keyHash <= serv[2] {
			u = getData(portno[2], a)
		} else {
			u = getData(portno[0], a)
		}

		fmt.Println(u)

	}

}
