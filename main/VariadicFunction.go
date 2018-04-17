package main

func main() {
	says("This", "is", "a", "book")
	says("Hi")
}

func says(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}
