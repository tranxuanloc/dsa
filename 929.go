package main

import "strings"

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{})
	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := strings.Split(parts[0], "+")[0]
		local = strings.ReplaceAll(local, ".", "")
		uniqueEmails[local+"@"+parts[1]] = struct{}{}
	}
	return len(uniqueEmails)
}
