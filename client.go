package coreapi

type CoreAPIClient[R any] struct {
	baseURL      string
	scans        IScanAPI[R]
	repositories IRepositoryAPI
}

type ICoreAPIClient[R any] interface {
	Scans() IScanAPI[R]
	Repositories() IRepositoryAPI
}

func NewCoreAPIClient[R any](baseURL string) ICoreAPIClient[R] {
	scanClient := NewScanAPI[R](baseURL)
	repositoryClient := NewRepositoryAPI(baseURL)
	return CoreAPIClient[R]{
		baseURL:      baseURL,
		scans:        scanClient,
		repositories: repositoryClient,
	}
}

func (c CoreAPIClient[R]) Scans() IScanAPI[R] {
	return c.scans
}

func (c CoreAPIClient[R]) Repositories() IRepositoryAPI {
	return c.repositories
}
