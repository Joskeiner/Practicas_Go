package main

func main() {
	servidor := NewServer(":3000")
	servidor.handle("/", handleRoot)
	servidor.handle("/Home", servidor.AddMidleware(handleHome, AuthCheck(), Logging()))
	servidor.Listen()
}
