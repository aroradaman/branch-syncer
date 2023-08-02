package main

import (
	"k8s.io/klog/v2"

	"github.com/aroradaman/branch-syncer/pkg/repository"
)

// main is the entrypoint
func main() {
	klog.InitFlags(nil)
	klog.Info("loaded main")

	rebaseRepository("git@github.com:acme-corp/kubernetes.git")
}

// rebaseRepository find all the target branches and rebase them with the parent one
func rebaseRepository(repoUrl string) {
	repo, err := repository.NewFromUrl(repoUrl)
	if err != nil {
		klog.ErrorS(err, "failed to clone repository")
	}

	branches, err := repo.GetBranches()
	if err != nil {
		klog.ErrorS(err, "failed to get branches")
	}

	pairs := PairBranches(branches)
	for _, pair := range pairs {
		_ = rebaseReference(repo, pair.acmeCorpBranch, pair.parentBranch)
	}
}

// rebaseReference will rebase the given branch
func rebaseReference(repo *repository.Repository, reference, base string) error {
	var err error

	err = repo.Checkout(repository.CheckoutOptions{
		Reference: reference,
		Create:    false,
	})
	if err != nil {
		klog.ErrorS(err, "failed checkout repository")
	}

	err = repo.Pull(repository.PullOptions{
		Reference: base,
		Rebase:    true,
		RebaseOptions: repository.RebaseOptions{
			Merges: true,
		},
	})

	if err != nil {
		klog.ErrorS(err, "failed to pull")
	}

	err = repo.Push(repository.PushOptions{
		Force: true,
	})
	if err != nil {
		klog.ErrorS(err, "failed to push")
	}

	return err
}
