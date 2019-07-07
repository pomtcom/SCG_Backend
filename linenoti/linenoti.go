package linenoti

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func LineNotification(lineMessage string){
	urlAddr := "https://notify-api.line.me/api/notify"
	lineToken := "LINE TOKEN"
	fmt.Println("URL:>", urlAddr)

	form := url.Values{}
	form.Add("message", lineMessage)

	req, err := http.NewRequest("POST", urlAddr, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer " + lineToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Line notification message is sent")
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
