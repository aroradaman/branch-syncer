package repository

type CheckoutOptions struct {
	Reference string
	Create    bool
}

type RebaseOptions struct {
	Merges bool
}

type PullOptions struct {
	RemoteUrl string
	Reference string
	Force     bool

	Rebase        bool
	RebaseOptions RebaseOptions
}

type PushOptions struct {
	RemoteUrl string
	Reference string
	Force     bool
}
