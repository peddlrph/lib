package smsgateway

import (
	//"encoding/json"
	//"fmt"
	"io/ioutil"
	//"log"
	"net/http"
)

func GetMessages(url string) ([]byte, error) {

	var response []byte

	res, err := http.Get(url)
	if err != nil {
		return response, err
		//log.Fatal(err)
	}
	defer res.Body.Close()
	response, err = ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return response, err

	//	fmt.Println("Response: ", string(response))

	//mesgs := Item{}
	//json.Unmarshal([]byte(response), &mesgs)

	// return mesgs
	//fmt.Println("Mesgs: ", mesgs)
	//fmt.Println("%T ", mesgs)

	//fmt.Println(string(response))
	//return res.Body
	//fmt.Printf("%s", robots)
}
