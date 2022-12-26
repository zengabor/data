package data

import "testing"

func TestEmailAddressToName(t *testing.T) {
	var tests = []struct {
		email string
		want  string
	}{
		{"email@inter.net", "Email"},
		{"EMAIL@inter.net", "Email"},
		{"eMaIl@inter.net", "Email"},
		{"sherlock.holmes@inter.net", "Sherlock Holmes"},
		{"sherlock_holmes@inter.net", "Sherlock Holmes"},
	}
	for _, test := range tests {
		if got := EmailAddressToName(test.email); got != test.want {
			t.Errorf(`EmailAddressToName(%q) = %q, want %q`, test.email, got, test.want)
		}
	}
}
