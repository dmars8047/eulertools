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

func GetPremise(num int, agentName string) string {

	client := &http.Client{}

	url := fmt.Sprintf("https://projecteuler.net/minimal=%d", num)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", agentName)
	req.Header.Add("Accept", "text/plain, text/html")

	if err != nil {
		fmt.Printf("Euler Premise Provider Error: %s", err.Error())
	}

	resp, _ := client.Do(req)

	body := resp.Body

	buffer := make([]byte, 1024)

	text := ""
	for true {
		i, _ := body.Read(buffer)

		if i != 0 {
			text += string(buffer)
		} else {
			break
		}
	}

	return removeHtmlTag(text)
}

func PrintPremise(num int, agentName string) {
	fmt.Println("\nProblem Premise:")
	fmt.Print(GetPremise(num, agentName))
}
