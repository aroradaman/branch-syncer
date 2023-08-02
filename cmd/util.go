package main

import (
	"strings"
)

type Predicate func(string) bool

// ShouldStartWith returns a prefix match predicate
func ShouldStartWith(s string) Predicate {
	return func(item string) bool {
		if strings.HasPrefix(item, s) {
			return true
		}
		return false
	}
}

// ShouldContain returns string contains predicate
func ShouldContain(s string) Predicate {
	return func(item string) bool {
		if strings.Contains(item, s) {
			return true
		}
		return false
	}
}

// Filter filters a list of string bases on the predicates
func Filter(items []string, predicates ...Predicate) []string {
	fItems := make([]string, 0)

	for _, item := range items {
		skip := false
		for i := 0; i < len(predicates); i++ {
			if !predicates[i](item) {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		fItems = append(fItems, item)
	}
	return fItems
}

// RemovePrefix removes prefix from each item in the list of items
func RemovePrefix(items []string, prefix string) []string {
	fItems := make([]string, 0)
	for _, item := range items {
		fItem := item
		if strings.HasPrefix(item, prefix) {
			fItem = fItem[len(prefix):]
		}

		fItems = append(fItems, fItem)
	}
	return fItems

}
