package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"strings"
)

func parseElement(res string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		return "", fmt.Errorf("解析失败: %w", err)
	}

	var buy_cart_text string
	btn_buy_now := doc.Find("#submit.buy-now-announce").Text()
	btn_add_to_ubb_cart := doc.Find("#submit.add-to-cart-ubb-announce").Text()
	btn_add_to_cart := doc.Find("#submit.add-to-cart-announce").Text()
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

func readFileContent(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main1() {
	content, err := readFileContent("./amazon/golang/abuyun/B08962MFDH/50afdd7a-74fe-4038-9be3-804e4a70a71b.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(parseElement(string(content)))
}

func main() {
	html := `<body>

				<div id="div1">DIV1aaaaaa</div>
				<span id="submit.buy-now-announce.add-to-cart" class="a-button-text-add-cart" aria-hidden="true">Add to cart</span>
				<span id="submit.buy-now-announce" class="a-button-text" aria-hidden="true"> Buy Now </span>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(dom.Find("#div1").Text())
	fmt.Println(dom.Find("span#submit\\.buy-now-announce\\.add-to-cart").Text())
	fmt.Println(dom.Find("#submit\\.buy-now-announce\\.add-to-cart").Text())
	fmt.Println(dom.Find("#submit\\.buy-now-announce").Text())

	/*
		dom.Find("span\\.a-button-text").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})*/

	/*
		dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
		})*/
}
