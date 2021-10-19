package go_interesting_code

func main() {
	defer func() {
		for {
		}
	}()
}

func mainV1() {
	defer func() {
		select {}
	}()
}
