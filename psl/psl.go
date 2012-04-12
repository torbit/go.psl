// regdom-libs/public suffix list for Go.
package psl

import (
	"strings"
)

func reverse(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func matchRules(labels []string) int {

	matched := 0

	r := pslRules
	for _, label := range labels {
		if match, exists := r.children[label]; exists {
			// matching label
			r = match
			if r.exception {
				break
			}
			matched++
		} else if match, exists := r.children["*"]; exists {
			// wildcard
			matched++
			r = match
			break
		} else {
			// no more matches
			break
		}
	}

	return matched

}

func splitDomain(domain string) []string {
	if len(domain) == 0 || domain[0] == '.' {
		return nil
	}

	domain = strings.ToLower(domain)
	labels := strings.Split(domain, ".")
	reverse(labels)

	return labels
}

func PublicSuffix(domain string) string {

	labels := splitDomain(domain)
	if labels == nil {
		return ""
	}

	matched := matchRules(labels)

	if matched > 0 {
		reverse(labels[0:matched])
		return strings.Join(labels[0:matched], ".")
	}

	return ""
}

func RegisteredDomain(domain string) string {

	labels := splitDomain(domain)
	if labels == nil {
		return ""
	}

	matched := matchRules(labels)

	if matched > 0 && matched < len(labels) {
		matched++
		reverse(labels[0:matched])
		return strings.Join(labels[0:matched], ".")
	}

	return ""
}
