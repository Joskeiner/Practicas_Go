package main

func main() {
	servidor := NewServer(":3000")
	servidor.handle("/", "GET", handleRoot)
	servidor.handle("/create", "POST", PostResquest)
	servidor.handle("/user", "POST", UserPostResquest)
	servidor.handle("/Home", "POST", servidor.AddMidleware(handleHome, AuthCheck(), Logging()))
	servidor.Listen()
}
