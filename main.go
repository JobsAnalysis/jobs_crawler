package main

import (
	"log"

	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/eddieivan01/nic"
)

func main() {
	// recommend to use
	random := browser.Random()
	log.Printf("Random: %s", random)
	/*
			a = session.get(
		    url="https://www.lagou.com/jobs/list_爬虫工程师?labelWords=&fromSearch=true&suginput=",
		    headers={
				"useragent": UserAgent().random})
	*/
	session := &nic.Session{}
	resp, err := session.Get("https://www.lagou.com/jobs/list_golang?labelWords=&fromSearch=true&suginput=", nic.H{
		Headers: nic.KV{
			"useragent": random,
		},
	})
	if err != nil {
		log.Printf("Random: %s", err)
	}
	log.Printf("Random: %s", resp.Response.Header)
	time.Sleep(5 * time.Second)
	url := "https://www.lagou.com/jobs/positionAjax.json"
	ua := browser.Random()
	r, _ := session.Post(
		url,
		nic.H{
			Headers: nic.KV{
				"user-agent":   ua,
				"referer":      "https://www.lagou.com/jobs/list_%E6%95%B0%E6%8D%AE%E5%88%86%E6%9E%90?city=%E5%85%A8%E5%9B%BD&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot",
				"Origin":       "https://www.lagou.com",
				"Accept":       "application/json, text/javascript, */*; q=0.01",
				"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			},
			Data: nic.KV{
				"first":               "true",
				"pn":                  "1",
				"kd":                  "golang",
				"city":                "上海",
				"needAddtionalResult": "false",
			},
		})
	log.Printf("Random: %s", r.Response.Body)
	// content = r.content

}
