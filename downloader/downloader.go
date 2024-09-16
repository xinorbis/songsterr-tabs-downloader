package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFile(url string) {
	fileUrl, filename := getFileData(getPageContent(url))

	processError(downloadFile(makeDownloadPath(filename), fileUrl))
	printMessage(filename)
}

func printMessage(filename string) {
	fmt.Println(fmt.Sprintf("Файл %s успешно скачан в директорию загрузок пользователя", filename))
}

func makeDownloadPath(filename string) string {
	userHomeDir, err := os.UserHomeDir()
	processError(err)

	return filepath.Join(userHomeDir, "Downloads", filename)
}

func getPageContent(url string) string {
	return getBody(getContent(url))
}

func getContent(url string) *http.Response {
	resp, err := http.Get(url)
	processError(err)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("При получении данных с сервера songsterr.com произошла ошибка: %v", resp.StatusCode))
	}

	return resp
}

func getBody(resp *http.Response) string {
	body, err := io.ReadAll(resp.Body)
	processError(err)
	defer resp.Body.Close()

	return string(body)
}

func getFileData(content string) (string, string) {
	return makeNameAndUrl(parseRequestedBody(content))
}

func parseRequestedBody(content string) []string {
	regexp, _ := regexp.Compile("\"artist\":\"(.+?)\",.+?\"title\":\"(.+?)\",.+?\"source\":\"(.+?)\",")
	result := regexp.FindStringSubmatch(content)
	if len(result) == 0 {
		panic("Parse error")
	}

	return result
}

func makeNameAndUrl(result []string) (string, string) {
	fileUrl := strings.Replace(result[3], "\\u002F", "/", 3)
	filename := fmt.Sprintf("%s - %s%s", result[1], result[2], path.Ext(fileUrl))

	return fileUrl, filename
}

func processError(err error) {
	if err != nil {
		panic(err)
	}
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url) // todo: remove to func
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
