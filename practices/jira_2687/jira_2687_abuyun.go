package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"github.com/iunary/fakeuseragent"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	cookieFile  = "/Users/guoqingliu/job/pingpong/jira_2687/US_90001.cookie"
	productFile = "/Users/guoqingliu/job/pingpong/jira_2687/products.txt"
	resultsFile = "abuyun_buybox_mt_log.txt"
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
	// 自定义生成http-header/user-agent
	randomAgent := fakeuseragent.RandomUserAgent()
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
	headers.Set("user-agent", randomAgent)
	headers.Set("cookie", cookie)
	return headers
}

func writeLogging(asin string, er error, content string) error {
	file, err := os.OpenFile(resultsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	var contentNew string = ""
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	contentNew += formattedTime
	if asin != "" {
		contentNew += ",Asin=" + asin
	}
	if er != nil {
		contentNew += ",Error=" + er.Error()
	}
	contentNew += "," + content
	contentNew += "\n"
	_, err = file.WriteString(contentNew)
	if err != nil {
		return fmt.Errorf("Failed write to file: %w", err)
	}
	return nil
}

func saveResponse(asin string, res string) error {
	// 创建一个UUID
	uuidStr := uuid.New().String()
	filePath := fmt.Sprintf("./amazon/golang/abuyun/%s/%s.html", asin, uuidStr)
	dirPath := filepath.Dir(filePath)
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("Failed to create directory: %w", err)
		}
	}

	if err := ioutil.WriteFile(filePath, []byte(res), 0644); err != nil {
		return fmt.Errorf("Failed write to file: %w", err)
	}
	return nil
}

func parseElement(res string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		return "", fmt.Errorf("解析失败: %w", err)
	}

	var buy_cart_text string
	btn_buy_now := doc.Find("#submit\\.buy-now-announce").Text()
	btn_add_to_ubb_cart := doc.Find("#submit\\.add-to-cart-ubb-announce").Text()
	btn_add_to_cart := doc.Find("#submit\\.add-to-cart-announce").Text()
	btn_setup_now := doc.Find("#rcx-subscribe-submit-button-announce").Text()

	if btn_buy_now != "" {
		buy_cart_text = btn_buy_now
	} else if btn_add_to_ubb_cart != "" {
		buy_cart_text = btn_add_to_ubb_cart
	} else if btn_add_to_cart != "" {
		buy_cart_text = btn_add_to_cart
	} else if btn_setup_now != "" {
		buy_cart_text = btn_setup_now
	}
	return buy_cart_text, nil
}

func selCookie(cookies []Cookie) Cookie {
	var cookieReturn Cookie
	for _, item := range cookies {
		if item.FailCount >= maxFailures {
			continue // 跳过失败的cookie
			//log.Printf("Max failures reached for cookie: %s, discarding it", item.Value)
		}
		cookieReturn = item
	}
	return cookieReturn
}

func fetch(asin string, cookies []Cookie, wg *sync.WaitGroup) {
	defer wg.Done()

	if asin != "" {
		asin = strings.ReplaceAll(asin, "\"", "")
	}
	url := getAmzProdUrl(asin)
	cookie := selCookie(cookies)
	var headers = setHeaders(cookie.Value)

	transport, err := getProxy()
	if err != nil {
		writeLogging(asin, err, "Error parsing proxy URL")
	}

	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		writeLogging(asin, err, "创建请求失败")
		return
	}
	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		cookie.FailCount++
		writeLogging(asin, err, "请求失败")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cookie.FailCount++
		writeLogging(asin, err, "读取响应失败")
		return
	}

	err = saveResponse(asin, string(body))
	if err != nil {
		writeLogging(asin, err, "写入响应到文件失败")
	}

	buyboxText, err := parseElement(string(body))
	if err != nil {
		writeLogging(asin, err, "解析响应失败")
	} else {
		writeLogging(asin, nil, "BuyBox="+buyboxText)
	}
}

func getAmzProdUrl(asin string) string {
	return fmt.Sprintf("https://www.amazon.com/dp/%s?th=1&psc=1", asin)
}

func getProxy() (*http.Transport, error) {
	// 代理服务器地址和端口，以及认证信息（如果需要的话）
	proxyURL := "http://HW1Q953296669L2D:822B0A8AC3A7CA2D@http-dyn.abuyun.com:9020"

	// 解析代理URL
	parsedProxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	// 创建一个新的Transport，并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(parsedProxy),
	}
	return transport, nil
}

func main() {
	// 设置随机数种子以获取随机结果（通常在程序开始时设置一次即可）
	rand.Seed(time.Now().UnixNano())

	writeLogging("", nil, ">>>阿布云(购物车)-多线程")

	var prods = readProd()
	var cookies = readCookie()

	successCount := 0
	failCount := 0

	// 使用一个WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 设置并发数为10
	concurrency := 3
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
