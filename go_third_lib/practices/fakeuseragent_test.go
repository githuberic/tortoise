package practices

import (
	"fmt"
	"github.com/iunary/fakeuseragent"
	"testing"
)

func TestTypeBasic(t *testing.T) {
	// Get a random user agent for Chrome
	chromeAgent := fakeuseragent.GetUserAgent(fakeuseragent.BrowserChrome)
	fmt.Println("Chrome User Agent:", chromeAgent)

	// Get a random user agent from the supported browsers
	randomAgent := fakeuseragent.RandomUserAgent()
	fmt.Println("Random User Agent:", randomAgent)
}
