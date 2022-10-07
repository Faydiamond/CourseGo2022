package main

import (
	"fmt"
	"net/http"
	"time"
)

func verifyServer(server string, myChanel chan string) {
	_, err := http.Get(server)
	if err != nil {
		//fmt.Println(server, " se encuentra caido. ")
		myChanel <- server + " se encuentra caido. "
	} else {
		//fmt.Println(server, " se encuentra funcionando. ")
		myChanel <- server + " se encuentra funcionando. "
	}
}

func main() {
	channel := make(chan string)
	init := time.Now()
	servers := []string{
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
		"https://www.soho.co/",
		"http://www.falso125*.com",
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
		"https://www.soho.co/",
		"http://www.falso125*.com",
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
		"https://www.soho.co/",
		"http://www.falso125*.com",
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
		"https://www.soho.co/",
		"http://www.falso125*.com",
		"https://www.youtube.com/",
		"https://www.udemy.com/",
		"https://www.radioacktiva.com/",
		"https://www.soho.co/",
		"http://www.falso125*.com",
	}

	for _, server := range servers {
		go verifyServer(server, channel)
	}
	finish := time.Since(init)

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	fmt.Println("El tiempo de ejecucion fue: ", finish)
}
