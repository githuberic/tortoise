package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	cookieFile  = "cookies.txt"
	productFile = ""
	proxyURL    = "http://your-proxy-address" // 替换为实际的代理服务器地址
	resultsFile = "amazon_results.txt"
	maxFailures = 3
)

type Cookie struct {
	Value     string
	FailCount int
}

var (
	successCount uint64
	failCount    uint64
)

func readCookie() []Cookie {
	var cookieList []Cookie
	var cookies []string
	cookies, err := readFileSplitLineRetArr(cookieFile)
	if err != nil {
		panic(err)
	} else {
		for _, value := range cookies {
			cookieList = append(cookieList, Cookie{Value: value})
		}
	}

	return cookieList
}

func readProd() []string {
	var prods []string
	prods, err := readFileSplitLineRetArr(productFile)
	if err != nil {
		panic(err)
	}
	return prods
}

func readFileSplitLineRetArr(filePath string) ([]string, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	buf := bufio.NewScanner(fi)
	var arrLine []string
	for {
		//文件读完了,退出for
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		if strings.TrimSpace(line) != "" {
			arrLine = append(arrLine, line)
		}
	}
	return arrLine, nil
}

func setHeaders(cookie string) http.Header {
	headers := http.Header{}
	headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	headers.Set("accept-language", "zh-CN,zh;q=0.9")
	headers.Set("cache-control", "max-age=0")
	headers.Set("device-memory", "8")
	headers.Set("downlink", "1.3")
	headers.Set("dpr", "1.8")
	headers.Set("ect", "3g")
	headers.Set("rtt", "500")
	headers.Set("sec-ch-device-memory", "8")
	headers.Set("sec-ch-dpr", "1.8")
	headers.Set("sec-ch-ua", "Not A(Brand\";v=\"99\", \"Google Chrome\";v=\"121\", \"Chromium\";v=\"121\"")
	headers.Set("sec-ch-ua-mobile", "?0")
	headers.Set("sec-ch-ua-platform", "macOS")
	headers.Set("sec-ch-viewport-width", "1600")
	headers.Set("sec-fetch-dest", "document")
	headers.Set("sec-fetch-mode", "navigate")
	headers.Set("sec-fetch-site", "none")
	headers.Set("sec-fetch-user", "?1")
	headers.Set("sec-gpc", "1")
	headers.Set("upgrade-insecure-requests", "1")
	headers.Set("viewport-width", "1600")
	headers.Set("user-agent", "")
	headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	headers.Set("cookie", cookie)
	return headers
}

func saveResponse(res string) {
	ioutil.WriteFile("response.txt", []byte(res), 0644)
}

func parseElement(res string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		fmt.Println("解析失败：", err)
		return
	}
	element := doc.Find("span#productTitle").Text()
	fmt.Println("获取元素值：", element)
}

func selCookie(cookies []Cookie) Cookie {
	var cookieReturn Cookie
	for _, item := range cookies {
		if item.FailCount >= maxFailures {
			continue // 跳过失败的cookie
			log.Printf("Max failures reached for cookie: %s, discarding it", item.Value)
		}
		cookieReturn = item
	}
	return cookieReturn

	// 生成0到len(arr)-1之间的随机索引
	// randomIndex := rand.Intn(len(cookies))
	// 使用随机索引选取数组中的元素
	// selectedItem := cookies[randomIndex]
	/*
		if selectedItem.FailCount >= maxFailures {
			log.Printf("Max failures reached for cookie: %s, discarding it", selectedItem.Value)
		}*/
}

func fetch(asin string, cookies []Cookie, wg *sync.WaitGroup) {
	defer wg.Done()

	url := getAmzProdUrl(asin)
	cookie := selCookie(cookies)
	var headers = setHeaders(cookie.Value)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败：", err)
		return
	}
	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		cookie.FailCount++
		fmt.Println("请求失败：", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cookie.FailCount++
		fmt.Println("读取响应失败：", err)
		return
	}

	saveResponse(string(body))
	parseElement(string(body))
}

func getAmzProdUrl(asin string) string {
	str := fmt.Sprintf("https://www.amazon.com/dp/{%s}?th=1&psc=1", asin)
	return str
}

// 获取代理服务器地址
func getProxy() (*url.URL, error) {
	resp, err := http.Get(proxyURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	proxyURL, err := url.Parse(string(body))
	if err != nil {
		return nil, err
	}

	return proxyURL, nil
}

func main() {
	// 设置随机数种子以获取随机结果（通常在程序开始时设置一次即可）
	rand.Seed(time.Now().UnixNano())

	var prods = readProd()
	var cookies = readCookie()

	successCount := 0
	failCount := 0

	// 使用一个WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 设置并发数为10
	concurrency := 20
	// 信号量控制并发数
	semaphore := make(chan struct{}, concurrency)

	for _, val := range prods {
		wg.Add(1)
		go func(asin string) {
			semaphore <- struct{}{}        // 拉低信号量，确保不超过设定的并发数
			defer func() { <-semaphore }() // 处理完后释放信号量
			fetch(asin, cookies, &wg)
		}(val)
	}

	// 等待所有任务完成
	wg.Wait()

	fmt.Printf("成功请求次数：%d, 失败请求次数：%d", successCount, failCount)
}
