package uri

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func Read(uri string) ([]model.Entry, error) {
	var result []model.Entry
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "file" {
		file, err := os.Open(u.Host + u.Path)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result = append(result, scanner.Text())
		}
		return result, nil
	}

	// file, err := os.Open(uri)
	// if err != nil {
	// 	return nil, err
	// }
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	result = append(result, scanner.Text())
	// }

	return nil, fmt.Errorf("Not a valid scheme")
}
