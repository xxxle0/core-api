package coreapi

type CoreAPIClient struct {
	baseURL      string
	scans        IScanAPI
	repositories IRepositoryAPI
}

func NewCoreAPIClient(baseURL string) CoreAPIClient {
	scanClient := NewScanAPI(baseURL)
	return CoreAPIClient{
		baseURL: baseURL,
		scans:   scanClient,
	}
}

func (c CoreAPIClient) Scans() IScanAPI {
	return c.scans
}

func (c CoreAPIClient) Repositories() IRepositoryAPI {
	return c.repositories
}
