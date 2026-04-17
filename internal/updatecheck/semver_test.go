package updatecheck

import "testing"

func TestParseSemver(t *testing.T) {
	cases := []struct {
		in   string
		want semver
		ok   bool
	}{
		{"0.3.2", semver{0, 3, 2}, true},
		{"1.0.0", semver{1, 0, 0}, true},
		{"10.20.30", semver{10, 20, 30}, true},
		{" 0.3.2 ", semver{0, 3, 2}, true},
		{"0.3.2\n", semver{0, 3, 2}, true},
		{"v0.3.2", semver{}, false},
		{"0.3.2-rc.1", semver{}, false},
		{"0.3.2+abc", semver{}, false},
		{"0.3", semver{}, false},
		{"0.3.2.4", semver{}, false},
		{"a.b.c", semver{}, false},
		{"-1.0.0", semver{}, false},
		{"", semver{}, false},
	}
	for _, c := range cases {
		got, ok := parseSemver(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("parseSemver(%q) = (%+v, %v); want (%+v, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestSemverLess(t *testing.T) {
	cases := []struct {
		a, b string
		less bool
	}{
		{"0.3.2", "0.4.0", true},
		{"0.3.2", "0.3.3", true},
		{"0.3.2", "1.0.0", true},
		{"1.0.0", "0.9.9", false},
		{"0.3.2", "0.3.2", false},
		{"0.10.0", "0.9.99", false},
	}
	for _, c := range cases {
		a, _ := parseSemver(c.a)
		b, _ := parseSemver(c.b)
		if got := a.less(b); got != c.less {
			t.Errorf("%s.less(%s) = %v; want %v", c.a, c.b, got, c.less)
		}
	}
}
