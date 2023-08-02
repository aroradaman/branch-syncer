package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/sets"
)

func TestPairBranches(t *testing.T) {
	branches := []string{
		"master", "master-acme-corp",
		"main", "main-acme-corp",
		"release-1.23", "release-1.23-acme-corp",
		"release-1.24", "release-1.24-acme-corp",
		"release/1.6", "release/1.6-acme-corp",
		"release/1.5", "release/1.5-acme-corp",
		"release-1.0", "release-1.0-acme-corp",
		"release-1.1", "release-1.1-acme-corp",
		// branches with no -acme-corp suffix
		"release-2.1", "release-2.2", "release-2.3",
		// standalone -acme-corp branches
		"release-3.1-acme-corp", "release-3.2-acme-corp",
	}

	pairs := PairBranches(branches)
	assert.Equal(t, 8, len(pairs))

	actualSet := sets.New(pairs...)
	expectedSet := sets.New([]branchPair{
		{"master", "master-acme-corp"},
		{"main", "main-acme-corp"},
		{"release-1.23", "release-1.23-acme-corp"},
		{"release-1.24", "release-1.24-acme-corp"},
		{"release/1.6", "release/1.6-acme-corp"},
		{"release/1.5", "release/1.5-acme-corp"},
		{"release-1.0", "release-1.0-acme-corp"},
		{"release-1.1", "release-1.1-acme-corp"},
	}...)

	assert.Equal(t, true, actualSet.Equal(expectedSet))

}

func TestShouldStartWith(t *testing.T) {
	assert.Equal(t, true, ShouldStartWith("v")("v1.0.0"))
	assert.Equal(t, false, ShouldStartWith("v")("1.0.0"))

	assert.Equal(t, true, ShouldStartWith("remotes/origin/")("remotes/origin/master-acme-corp"))
	assert.Equal(t, false, ShouldStartWith("remotes/origin/")("master-acme-corp"))
}

func TestShouldContain(t *testing.T) {
	assert.Equal(t, true, ShouldContain("HEAD")("remotes/origin/HEAD"))
	assert.Equal(t, false, ShouldContain("HEAD")("remotes/origin/main-acme-corp"))

	assert.Equal(t, true, ShouldContain("remotes/origin/")("remotes/origin/master-acme-corp"))
	assert.Equal(t, false, ShouldContain("remotes/origin/")("master-acme-corp"))
}

func TestFilter(t *testing.T) {
	var items []string

	items = []string{"v1.0.0", "v2.0.0", "v3.0.0", "v4", "1.0.0", "2.0", "3"}
	assert.Equal(t, []string{"v1.0.0", "v2.0.0", "v3.0.0", "v4"}, Filter(items, ShouldStartWith("v")))

	items = []string{"v1.0.0", "v2.0.0-acme-corp-es1", "v3.0.0-acme-corp", "v4", "1.0.0-acme-corp-es1", "2.0-acme-corp", "3"}
	assert.Equal(t, []string{"v2.0.0-acme-corp-es1", "v3.0.0-acme-corp"}, Filter(items, ShouldStartWith("v"), ShouldContain("acme")))

	items = []string{"v1.0.0", "v2.0.0-acme-corp-es1", "v3.0.0-acme-corp", "v4", "1.0.0-acme-corp-es1", "2.0-acme-corp", "3"}
	assert.Equal(t, []string{"v2.0.0-acme-corp-es1"}, Filter(items, ShouldStartWith("v"), ShouldContain("acme"), ShouldContain("es")))

}

func TestRemovePrefix(t *testing.T) {
	var items []string

	items = []string{
		"remotes/origin/main-acme-corp", "remotes/origin/main", "remotes/origin/release-1.1-acme-corp", "remotes/origin/release-1.1",
		"main-acme-corp", "main", "release-1.1-acme-corp", "release-1.1",
	}
	assert.Equal(t, []string{
		"main-acme-corp", "main", "release-1.1-acme-corp", "release-1.1",
		"main-acme-corp", "main", "release-1.1-acme-corp", "release-1.1",
	}, RemovePrefix(items, "remotes/origin/"))
}
