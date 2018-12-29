package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type INFO struct{
	gs_name string
	gs_pwd string
	git_name string
	git_pwd string
	addr string
	delay string
	cookie string
}


func (info *INFO)init(){
	cfg, _ := goconfig.LoadConfigFile("./info.conf")
	info.gs_name, _ = cfg.GetValue("gitStar", "gs_name")
	info.gs_pwd, _ = cfg.GetValue("gitStar", "gs_pwd")
	info.git_name, _ = cfg.GetValue("gitStar", "git_name")
	info.git_pwd, _ = cfg.GetValue("gitStar", "git_pwd")
	info.addr, _ = cfg.GetValue("gitStar", "addr")
	info.delay, _ = cfg.GetValue("gitStar", "delay")
	info.cookie = ""
}

// http request func
func (info *INFO)http_req(uri string, value url.Values, method string, headers map[string]string) (*http.Response){
	var pv io.Reader
	if value != nil{
		post_value := value.Encode()
		pv = strings.NewReader(post_value)
	}else{
		pv=nil
	}
	req, err := http.NewRequest(method, uri, pv)
	if err != nil{
		os.Exit(1)
	}
	if headers != nil{
		for k,v := range(headers){
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		os.Exit(1)
	}
	return resp
}

// gitStar
// login gitstar website to get your own cookie on it
func (info *INFO)loginGitStar(){
	uri := "http://" + info.addr + "/api/user/login"
	value := url.Values{
		"username":{info.gs_name},
		"password":{info.gs_pwd},
	}
	method := "POST"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	resp := info.http_req(uri, value, method, headers)
	tmp := resp.Header.Get("Set-Cookie")
	info.cookie = tmp
}

func (info *INFO)getGitStarList() []string{
	uri := "http://" + info.addr + "/api/users/" + info.gs_name + "/status/"
	value := url.Values{}
	method := "GET"
	res := make([]string,0)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Cookie"] = info.cookie
	resp := info.http_req(uri, value, method, headers)
	body, _ := ioutil.ReadAll(resp.Body)
	list, _ :=  simplejson.NewJson(body)
	items, _ := list.Array()
	for i, _ := range(items){
		res = append(res, list.GetIndex(i).Get("Repo").MustString())
	}
	return res
}

func (info *INFO)update_gs(){
	uri := "http://" + info.addr + "/star_update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Cookie"] = info.cookie
	for{
		resp := info.http_req(uri, nil, "GET", headers)
		if resp.StatusCode == 200{
			break
		}
	}
}

// github
func (info *INFO)star(repo string){
	uri := "https://api.github.com/user/starred/" + repo
	req, _ := http.NewRequest("PUT", uri, nil)
	req.SetBasicAuth(info.git_name, info.git_pwd)
	req.Header.Add("Content-Length", "0")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil{
		os.Exit(1)
	}

}


func run_star(){
	test := INFO{}
	test.init()
	test.loginGitStar()
	repos := test.getGitStarList()
	fmt.Printf("Get %d repotories to star\n", len(repos))
	for i, repo := range(repos){
		test.star(repo)
		fmt.Printf("[Start %d] Finished\n", i+1)
		tt, _ := strconv.Atoi(test.delay)
		time.Sleep(time.Duration(tt) * time.Second)
	}
	if len(repos) > 0{
		test.update_gs()
	}
}

func main() {
	run_star()
	fmt.Printf("Auto Star Finished\nPlease press Enter to quit\n")
	for{
		var b []byte = make([]byte, 1)
		for {
			_, err := os.Stdin.Read(b)
			if err != nil{
				fmt.Printf("Please press Enter to quit\n")
			}
			if string(b) == "\r\n" || string(b) == "\n"{
				os.Exit(0)
			}
		}
	}
}
