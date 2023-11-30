package main

func main() {
	server := ProvideServer()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
