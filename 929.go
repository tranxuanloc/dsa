package main

import "strings"

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{})
loop:
	for _, email := range emails {
		plus := false
		n := len(email)
		var emailBuilder strings.Builder
		for i := 0; i < n; i++ {
			ch := email[i]
			if ch == '+' {
				plus = true
				continue
			}
			if ch == '@' {
				emailBuilder.WriteString(email[i:])
				uniqueEmails[emailBuilder.String()] = struct{}{}
				continue loop
			}
			if ch == '.' || plus {
				continue
			}
			emailBuilder.WriteByte(ch)
		}
	}
	return len(uniqueEmails)
}
