package main

func main() {
	msg := "Hello"
	say(&msg)
	println(msg)
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}
