package eulertools

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

// match html tag and replace it with ""
func removeHtmlTag(in string) string {
	// regex to match html tag
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)
	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}

func GetPremise(num int) string {

	client := &http.Client{}

	url := fmt.Sprintf("https://projecteuler.net/minimal=%d", num)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "DSM-Reader/1.0")
	req.Header.Add("Accept", "text/plain, text/html")

	if err != nil {
		fmt.Printf("Euler Premise Provider Error: %s", err.Error())
	}

	resp, _ := client.Do(req)

	body := resp.Body

	buffer := make([]byte, 1024)

	i := 1
	text := ""
	for i != 0 {
		i, _ = body.Read(buffer)
		text += string(buffer)
	}

	return removeHtmlTag(text)
}

func PrintPremise(num int) {
	fmt.Println("\nProblem Premise:")
	fmt.Println()
	fmt.Print(GetPremise(num))
}
