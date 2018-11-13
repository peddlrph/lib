package smsgateway

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Timestamp    int  `json:"timestamp"`
	AirplaneMode bool `json:"is_airplane_mode"`
	Telephony    Telephony
	Telephonies  []Telephony
}

type Telephony struct {
	NetworkRoaming      bool   `json:"is_network_roaming"`
	SimState            string `json:"sim_state"`
	NetworkOperatorName string `json:"network_operator_name"`
	DisplayName         string `json:"display_name"`
	SimSlot             int    `json:"sim_slot"`
}

func GetMessages(ipaddr string, limit string, offset string) ([]byte, error) {

	res, err := http.Get("http://" + ipaddr + ":8080/v1/sms/?limit=" + limit + "&offset=" + offset)
	if err != nil {
		//log.Fatal(err)
		return []byte("error"), err
	}

	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	//fmt.Println(response)
	if err != nil {
		log.Fatal(err)
		return []byte(response), err
	}

	//response = []byte("Hello")

	fmt.Println("response")

	return response, err
	//return []byte("res.Body"), err

}

func WriteMessagesToFile(resp []byte) {

	dateFormat := "2006Jan02-15H"

	currentTime := time.Now()

	outputfile := "./asset/messages/messages" + currentTime.Format(dateFormat) + ".json"

	f, _ := os.Create(outputfile)

	writer := bufio.NewWriter(f)

	defer f.Close()

	_, _ = writer.Write(resp)

	writer.Flush()

	//_ = pie.Render(chart.PNG, writer)
}

func CheckStatus(ip_address string) string {

	res, err := http.Get("http://" + ip_address + ":8080/v1/device/status")
	if err != nil {
		//log.Fatal(err)
		return "Offline"
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//log.Fatal(err)
		return "Offline"
	}

	status := Status{}
	json.Unmarshal([]byte(response), &status)

	if status.Telephonies[0].SimState == "ready" {
		return status.Telephonies[0].SimState
	} else {
		return "Offline"
	}

	return "Offline"
}
