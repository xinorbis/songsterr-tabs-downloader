package main

import (
	"bufio"
	"fmt"
	"os"
	"songster/downloader"
	"strings"
)

func main() {
	fmt.Println("Утилита для скачивания табулатур с сервиса songsterr.com")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("Для того, что бы скачать файл .gpt укажите url страницы с интересующим Вас произведением и нажмите enter")

		url, _ := reader.ReadString('\n')
		if strings.TrimSpace(url) == "" {
			fmt.Println("Адрес не может быть пустым")
			continue
		}
		url = strings.TrimRight(url, "\r\n")
		downloader.GetFile(url)
	}
}
