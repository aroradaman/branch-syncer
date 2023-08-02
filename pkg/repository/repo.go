package repository

type Repository struct {
	remoteUrl string
	localPath string
}

// NewFromUrl clones a repo
func NewFromUrl(remoteUrl string) (*Repository, error) {
	clonePath := getTempDir()

	err := Clone(remoteUrl, clonePath)
	if err != nil {
		return nil, err
	}

	return &Repository{remoteUrl: remoteUrl, localPath: clonePath}, nil
}

// RemoteUrl returns remoteUrl
func (r *Repository) RemoteUrl() string {
	return r.remoteUrl
}

// LocalPath returns localPath
func (r *Repository) LocalPath() string {
	return r.localPath
}

// Checkout : git checkout
func (r *Repository) Checkout(options CheckoutOptions) error {
	return Checkout(options, r.localPath)
}

// Pull : git pull
func (r *Repository) Pull(options PullOptions) error {
	if options.RemoteUrl == "" {
		options.RemoteUrl = r.RemoteUrl()
	}

	return Pull(options, r.LocalPath())
}

// Push : git push
func (r *Repository) Push(options PushOptions) error {
	if options.RemoteUrl == "" {
		options.RemoteUrl = r.RemoteUrl()
	}

	return Push(options, r.LocalPath())
}

// GetBranches : git branch -a
func (r *Repository) GetBranches() ([]string, error) {
	return GetBranches(r.localPath)
}

// GetTags : git tag
func (r *Repository) GetTags() ([]string, error) {
	return GetTags(r.localPath)
}
