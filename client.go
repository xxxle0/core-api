package coreapi

type CoreAPIClient struct {
	baseURL      string
	scans        IScanAPI
	repositories IRepositoryAPI
}

type ICoreAPIClient interface {
	Scans() IScanAPI
}

func NewCoreAPIClient(baseURL string) ICoreAPIClient {
	scanClient := NewScanAPI(baseURL)
	repositoryClient := NewRepositoryAPI(baseURL)
	return CoreAPIClient{
		baseURL:      baseURL,
		scans:        scanClient,
		repositories: repositoryClient,
	}
}

func (c CoreAPIClient) Scans() IScanAPI {
	return c.scans
}

func (c CoreAPIClient) Repositories() IRepositoryAPI {
	return c.repositories
}
