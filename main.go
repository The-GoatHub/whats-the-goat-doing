package main

func main() {
	server := ProvideHardcodedServer() // or ProvideInMemoryServer()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
