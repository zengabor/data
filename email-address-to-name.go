package data

import "strings"

func EmailAddressToName(emailAddress string) string {
	return strings.Title(
		strings.TrimSpace(
			strings.Replace(
				strings.Replace(
					strings.Split(strings.ToLower(emailAddress), "@")[0],
					".", " ", -1),
				"_", " ", -1)))
}
