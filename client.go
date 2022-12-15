package coreapi

type CoreAPIClient struct {
	baseURL      string
	Scans        IScanAPI
	Repositories IRepositoryAPI
}

type ICoreAPIClient interface {
	IScanAPI
	IRepositoryAPI
}

func NewCoreAPIClient(baseURL string) ICoreAPIClient {
	scanClient := NewScanAPI(baseURL)
	repositoryClient := NewRepositoryAPI(baseURL)
	return CoreAPIClient{
		baseURL:      baseURL,
		Scans:        scanClient,
		Repositories: repositoryClient,
	}
}
