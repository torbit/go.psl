package psl

import (
	"testing"
)

func TestPublicSuffixData(t *testing.T) {

	checkPublicSuffix := func(domain, suffix string) {
		ret := RegisteredDomain(domain)
		if ret != suffix {
			t.Errorf("PublicSuffix(%q) got %q expected %q\n", domain, ret, suffix)
		}
	}

	// test data from http://publicsuffix.org/list/test.txt
	// NULL input.
	checkPublicSuffix("", "")
	// Mixed case.
	checkPublicSuffix("COM", "")
	checkPublicSuffix("example.COM", "example.com")
	checkPublicSuffix("WwW.example.COM", "example.com")
	// Leading dot.
	checkPublicSuffix(".com", "")
	checkPublicSuffix(".example", "")
	checkPublicSuffix(".example.com", "")
	checkPublicSuffix(".example.example", "")
	// Unlisted TLD.
	checkPublicSuffix("example", "")
	checkPublicSuffix("example.example", "")
	checkPublicSuffix("b.example.example", "")
	checkPublicSuffix("a.b.example.example", "")
	// Listed, but non-Internet, TLD.
	//checkPublicSuffix("local", "");
	//checkPublicSuffix("example.local", "");
	//checkPublicSuffix("b.example.local", "");
	//checkPublicSuffix("a.b.example.local", "");
	// TLD with only 1 rule.
	checkPublicSuffix("biz", "")
	checkPublicSuffix("domain.biz", "domain.biz")
	checkPublicSuffix("b.domain.biz", "domain.biz")
	checkPublicSuffix("a.b.domain.biz", "domain.biz")
	// TLD with some 2-level rules.
	checkPublicSuffix("com", "")
	checkPublicSuffix("example.com", "example.com")
	checkPublicSuffix("b.example.com", "example.com")
	checkPublicSuffix("a.b.example.com", "example.com")
	checkPublicSuffix("uk.com", "")
	checkPublicSuffix("example.uk.com", "example.uk.com")
	checkPublicSuffix("b.example.uk.com", "example.uk.com")
	checkPublicSuffix("a.b.example.uk.com", "example.uk.com")
	checkPublicSuffix("test.ac", "test.ac")
	// TLD with only 1 (wildcard) rule.
	checkPublicSuffix("cy", "")
	checkPublicSuffix("c.cy", "")
	checkPublicSuffix("b.c.cy", "b.c.cy")
	checkPublicSuffix("a.b.c.cy", "b.c.cy")
	// More complex TLD.
	checkPublicSuffix("jp", "")
	checkPublicSuffix("test.jp", "test.jp")
	checkPublicSuffix("www.test.jp", "test.jp")
	checkPublicSuffix("ac.jp", "")
	checkPublicSuffix("test.ac.jp", "test.ac.jp")
	checkPublicSuffix("www.test.ac.jp", "test.ac.jp")
	checkPublicSuffix("kyoto.jp", "")
	checkPublicSuffix("c.kyoto.jp", "")
	checkPublicSuffix("b.c.kyoto.jp", "b.c.kyoto.jp")
	checkPublicSuffix("a.b.c.kyoto.jp", "b.c.kyoto.jp")
	checkPublicSuffix("pref.kyoto.jp", "pref.kyoto.jp")     // Exception rule.
	checkPublicSuffix("www.pref.kyoto.jp", "pref.kyoto.jp") // Exception rule.
	checkPublicSuffix("city.kyoto.jp", "city.kyoto.jp")     // Exception rule.
	checkPublicSuffix("www.city.kyoto.jp", "city.kyoto.jp") // Exception rule.
	// TLD with a wildcard rule and exceptions.
	checkPublicSuffix("om", "")
	checkPublicSuffix("test.om", "")
	checkPublicSuffix("b.test.om", "b.test.om")
	checkPublicSuffix("a.b.test.om", "b.test.om")
	checkPublicSuffix("songfest.om", "songfest.om")
	checkPublicSuffix("www.songfest.om", "songfest.om")
	// US K12.
	checkPublicSuffix("us", "")
	checkPublicSuffix("test.us", "test.us")
	checkPublicSuffix("www.test.us", "test.us")
	checkPublicSuffix("ak.us", "")
	checkPublicSuffix("test.ak.us", "test.ak.us")
	checkPublicSuffix("www.test.ak.us", "test.ak.us")
	checkPublicSuffix("k12.ak.us", "")
	checkPublicSuffix("test.k12.ak.us", "test.k12.ak.us")
	checkPublicSuffix("www.test.k12.ak.us", "test.k12.ak.us")
}

func TestRegisteredDomain(t *testing.T) {
	if RegisteredDomain("www.google.com") != "google.com" {
		t.Fail()
	}
	if RegisteredDomain("www.google.co.uk") != "google.co.uk" {
		t.Fail()
	}
	if RegisteredDomain("something.unknown") != "" {
		t.Fail()
	}
	if RegisteredDomain("co.uk") != "" {
		t.Fail()
	}
}

func TestPublicSuffix(t *testing.T) {
	if PublicSuffix("www.google.com") != "com" {
		t.Fail()
	}
	if PublicSuffix("www.google.co.uk") != "co.uk" {
		t.Fail()
	}
	if PublicSuffix("something.unknown") != "" {
		t.Fail()
	}
	if PublicSuffix("co.uk") != "co.uk" {
		t.Fail()
	}
}

func ExampleRegisteredDomain() {
	RegisteredDomain("www.google.com")    // "google.com"
	RegisteredDomain("www.google.co.uk")  // "google.co.uk"
	RegisteredDomain("something.unknown") // ""
	RegisteredDomain("co.uk")             // ""
}

func ExamplePublicSuffix() {
	PublicSuffix("www.google.com")    // "com"
	PublicSuffix("www.google.co.uk")  // "co.uk"
	PublicSuffix("something.unknown") // ""
	PublicSuffix("co.uk")             // "co.uk"
}
