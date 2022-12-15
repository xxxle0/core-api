package coreapi

type CoreAPIClient[R any, A any] struct {
	baseURL      string
	scans        IScanAPI[R]
	repositories IRepositoryAPI[A]
}

type ICoreAPIClient[R, A any] interface {
	Scans() IScanAPI[R]
	Repositories() IRepositoryAPI[A]
}

func NewCoreAPIClient[R, A any](baseURL string) ICoreAPIClient[R, A] {
	scanClient := NewScanAPI[R](baseURL)
	repositoryClient := NewRepositoryAPI[A](baseURL)
	return CoreAPIClient[R, A]{
		baseURL:      baseURL,
		scans:        scanClient,
		repositories: repositoryClient,
	}
}

func (c CoreAPIClient[R, _]) Scans() IScanAPI[R] {
	return c.scans
}

func (c CoreAPIClient[_, A]) Repositories() IRepositoryAPI[A] {
	return c.repositories
}
