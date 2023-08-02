package main

import (
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	remotesOriginRepr = "remotes/origin/"

	// prefix which the organization uses for maintaining their own branch
	acmeCorpBranchSuffix = "-acme-corp"
)

// branchPair represent a parentBranch and acmeCorpBranch pair
type branchPair struct {
	parentBranch   string
	acmeCorpBranch string
}

// PairBranches create pairs of branchPair from given list of branches.
func PairBranches(branches []string) []branchPair {
	branchSet := sets.New(branches...)

	pairs := make([]branchPair, 0)

	for _, branch := range branches {
		if strings.HasSuffix(branch, acmeCorpBranchSuffix) {
			parentBranch := branch[:len(branch)-len(acmeCorpBranchSuffix)]
			if branchSet.Has(parentBranch) {
				pairs = append(pairs, branchPair{
					parentBranch:   parentBranch,
					acmeCorpBranch: branch,
				})
			}
		}
	}

	return pairs
}
