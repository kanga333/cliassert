package cliassert

import (
	"testing"
)

func TestContainCase(t *testing.T) {
	cases := []struct {
		CaseString   string
		Input        string
		WantAssert   bool
		WantDescribe string
	}{
		{"t", "test", true, "should contain t"},
		{"n", "test", false, "should contain n"},
	}

	for _, c := range cases {
		cc := NewContainCase(c.CaseString)
		gotAssert := cc.Assert(c.Input)
		if gotAssert != c.WantAssert {
			t.Errorf("Assert got: %v, want: %v:", gotAssert, c.WantAssert)
		}
		gotDescribe := cc.Describe()
		if gotDescribe != c.WantDescribe {
			t.Errorf("Describe got: %v, want: %v:", gotDescribe, c.WantDescribe)
		}
	}
}

func TestNotContainCase(t *testing.T) {
	cases := []struct {
		CaseString   string
		Input        string
		WantAssert   bool
		WantDescribe string
	}{
		{"t", "test", false, "should not contain t"},
		{"n", "test", true, "should not contain n"},
	}

	for _, c := range cases {
		cc := NewNotContainCase(c.CaseString)
		gotAssert := cc.Assert(c.Input)
		if gotAssert != c.WantAssert {
			t.Errorf("Assert got: %v, want: %v:", gotAssert, c.WantAssert)
		}
		gotDescribe := cc.Describe()
		if gotDescribe != c.WantDescribe {
			t.Errorf("Describe got: %v, want: %v:", gotDescribe, c.WantDescribe)
		}
	}
}

func TestRegexCase(t *testing.T) {
	cases := []struct {
		CaseString   string
		Input        string
		WantAssert   bool
		WantDescribe string
	}{
		{"te.t", "test", true, "should match te.t"},
		{"te..t", "test", false, "should match te..t"},
	}

	for _, c := range cases {
		cc := NewRegexCase(c.CaseString)
		gotAssert := cc.Assert(c.Input)
		if gotAssert != c.WantAssert {
			t.Errorf("Assert got: %v, want: %v:", gotAssert, c.WantAssert)
		}
		gotDescribe := cc.Describe()
		if gotDescribe != c.WantDescribe {
			t.Errorf("Describe got: %v, want: %v:", gotDescribe, c.WantDescribe)
		}
	}
}

func TestNotRegexCase(t *testing.T) {
	cases := []struct {
		CaseString   string
		Input        string
		WantAssert   bool
		WantDescribe string
	}{
		{"te.t", "test", false, "should not match te.t"},
		{"te..t", "test", true, "should not match te..t"},
	}

	for _, c := range cases {
		cc := NewNotRegexCase(c.CaseString)
		gotAssert := cc.Assert(c.Input)
		if gotAssert != c.WantAssert {
			t.Errorf("Assert got: %v, want: %v:", gotAssert, c.WantAssert)
		}
		gotDescribe := cc.Describe()
		if gotDescribe != c.WantDescribe {
			t.Errorf("Describe got: %v, want: %v:", gotDescribe, c.WantDescribe)
		}
	}
}

func TestEqualCase(t *testing.T) {
	cases := []struct {
		CaseString   string
		Input        string
		WantAssert   bool
		WantDescribe string
	}{
		{"test", "test", true, "should be equal test"},
		{"tes", "test", false, "should be equal tes"},
	}

	for _, c := range cases {
		cc := NewEqualCase(c.CaseString)
		gotAssert := cc.Assert(c.Input)
		if gotAssert != c.WantAssert {
			t.Errorf("Assert got: %v, want: %v:", gotAssert, c.WantAssert)
		}
		gotDescribe := cc.Describe()
		if gotDescribe != c.WantDescribe {
			t.Errorf("Describe got: %v, want: %v:", gotDescribe, c.WantDescribe)
		}
	}
}
