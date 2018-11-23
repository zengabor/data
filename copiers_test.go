package data

import "testing"

func TestChangeAwareCopierString(t *testing.T) {
	stringTests := []struct {
		in, out, wantOut string
		want             bool
	}{
		{"", "", "", false},
		{"abc", "abc", "abc", false},
		{"", "abc", "", true},
		{"abc", "", "abc", true},
		{"abc", "def", "abc", true},
	}
	for i, test := range stringTests {
		var c changeAwareCopier
		c.copyString(test.in, &test.out)
		if c.DestinationUpdated() != test.want {
			t.Errorf("#%d want %v, got %v", i, test.want, c.DestinationUpdated())
		}
		if test.wantOut != test.out {
			t.Errorf("#%d want %v, got %v", i, test.wantOut, test.out)
		}
	}
}

func TestChangeAwareCopierStringIfAny(t *testing.T) {
	stringTests := []struct {
		in, out string
		want    bool
	}{
		{"", "", false},
		{"abc", "abc", false},
		{"", "abc", false},
		{"abc", "", true},
		{"abc", "def", true},
	}
	for i, test := range stringTests {
		var c changeAwareCopier
		if c.copyStringIfAny(test.in, &test.out); c.DestinationUpdated() != test.want {
			t.Errorf("#%d want %v, got %v", i, test.want, c.DestinationUpdated())
		}
	}
}

func TestChangeAwareCopierStringWithWarning(t *testing.T) {
	stringTests := []struct {
		in, out string
		want    bool
	}{
		{"", "", false},
		{"abc", "abc", false},
		{"", "abc", false},
		{"abc", "", true},
		{"abc", "def", true},
	}
	for i, test := range stringTests {
		var c changeAwareCopier
		if c.copyStringWithWarning(test.in, &test.out, "warning"); c.DestinationUpdated() != test.want {
			t.Errorf("#%d want %v, got %v", i, test.want, c.DestinationUpdated())
		}
	}
}

func TestChangeAwareCopierStringIfMissing(t *testing.T) {
	stringTests := []struct {
		in, out string
		want    bool
	}{
		{"", "", false},
		{"abc", "abc", false},
		{"", "abc", false},
		{"abc", "", true},
		{"abc", "def", false},
	}
	for i, test := range stringTests {
		var c changeAwareCopier
		if c.copyStringIfMissing(test.in, &test.out); c.DestinationUpdated() != test.want {
			t.Errorf("#%d want %v, got %v", i, test.want, c.DestinationUpdated())
		}
	}
}

func TestDiffAwareCopierStringIfMissing(t *testing.T) {
	stringTests := []struct {
		in, out, wantOut string
		want             bool
	}{
		{"", "", "", false},
		{"abc", "abc", "abc", false},
		{"", "abc", "abc", true},
		{"abc", "", "abc", false},
		{"abc", "def", "def", true},
	}
	for i, test := range stringTests {
		var c diffAwareCopier
		c.copyStringIfMissing(test.in, &test.out)
		if c.Different() != test.want {
			t.Errorf("#%d want %v, got %v", i, test.want, c.Different())
		}
		if test.wantOut != test.out {
			t.Errorf("#%d want %v, got %v", i, test.wantOut, test.out)
		}
	}
}
