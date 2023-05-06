package urlCheck

import "net/url"

func CheckIsUrl(str string) (bool, error) { // проверка на url адрес
	parsedUrl, err := url.Parse(str)
	return err == nil && parsedUrl.Scheme != "" && parsedUrl.Host != "", err
}
