package main

import (
	"GoSmallCase/reptile/request"
	"GoSmallCase/reptile/scheduler"
	"GoSmallCase/reptile/xpath"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var dispatcher scheduler.Dispatcher
var num int32
var file *os.File

func main() {
	cookieMap := request.AnalysisCookieString(`NOWCODERUID=792B95DA084A8262FBFACFDF9BFD55A3; NOWCODERCLINETID=674ECA51D47406EE734D6C1D29455B91; gr_user_id=57829fd4-ffc3-44b2-84ae-6c62740e272a; grwng_uid=55ffc44b-39f2-43c5-bcd2-b41579e94fe8; c196c3667d214851b11233f5c17f99d5_gr_last_sent_cs1=4459196; _9755xjdesxxd_=32; gdxidpyhxdE=Bj5u7h4RkpAL2lW%2FrRgZ4rc%2BPVMpuvqMYHGIDl7XGWvvtlt%5C9G2mV0b8EYTUC4gOh8gM7Lg2O7riRsKcuHxkqsINm2rPWzwjzzbyVMpmPQUjG6mD8I1c8vCQl0mgeO4bT%2F%5CkgIvnA5A2AMzvPlcvxDRui%5C0T4%2Bk9aC6gGP7pyrMssHQ5%3A1563881227888; Hm_lvt_a808a1326b6c06c437de769d1b85b870=1563495358,1563597344,1563854674,1564026329; c196c3667d214851b11233f5c17f99d5_gr_session_id=567479a5-68f8-4816-8b0d-72aa1bcc4568; c196c3667d214851b11233f5c17f99d5_gr_last_sent_sid_with_cs1=567479a5-68f8-4816-8b0d-72aa1bcc4568; c196c3667d214851b11233f5c17f99d5_gr_session_id_567479a5-68f8-4816-8b0d-72aa1bcc4568=true; callBack=%2Fprofile%2F4459196%2Fresume%3F%26headNav%3Dwww; t=6D5C35983E9F060E28C3F6B14D3CA840; c196c3667d214851b11233f5c17f99d5_gr_cs1=4459196; Hm_lpvt_a808a1326b6c06c437de769d1b85b870=1564043403; SERVERID=9e4b74fdb43c9945205776603264d280|1564043408|1564026327`)
	headerMap := make(map[string]string)

	headerMap["User-Agent"] = request.UserAgentPCChrome

	var err error
	file, err = os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	dispatcher = scheduler.Dispatcher{
		BaseUrl: "https://www.nowcoder.com",
		Cookies: cookieMap,
		Headers: headerMap,
	}
	dispatcher.Init(
		"/profile/4459196/wrongset?onlyWrong=0&tags=&page=%d",
		`//td[@class='t-subject-title']/a[@class='test-subject']`,
		each,
		time.Millisecond*500,
		&scheduler.Pagination{StartPage: 26, EndPage: 26})

}

func each(nodes xpath.Nodes) {
	href := nodes.Attr("href")
	for _, v := range href {
		dispatcher.Add(v, `//div[@class='subject-des']`, question)
	}
}

func question(nodes xpath.Nodes) {
	text := nodes.Text()
	for _, v := range text {
		atomic.AddInt32(&num, 1)
		sprintf := fmt.Sprintf("%d. %s\n\n", num, format(v))

		write([]byte(sprintf))
	}
}

func format(str string) string {
	return strings.TrimSpace(str)
}

func write(b []byte) {
	if _, err := file.Write(b); err != nil {
		panic(err)
	}
}
