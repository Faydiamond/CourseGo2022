package main

import (
	"fmt"
	"net/http"
	"time"
)

func verifyServer(server string) {

	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, " se encuentra caido!")
	} else {
		fmt.Println(server, " se encuentra funcionando!")
	}
}

func main() {
	iniit := time.Now()
	servers := []string{
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
	}

	for _, server := range servers {
		verifyServer(server)
	}
	finish := time.Since(iniit)
	fmt.Println("El tiempo transcurrido es: ", finish)
}
