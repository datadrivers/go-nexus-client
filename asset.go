package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	assetAPIEndpoint = "service/rest/v1/assets"
)

type AssetResponse struct {
	Items             []Asset     `json:"items,omitempty"`
	ContinuationToken interface{} `json:"continuationToken,omitempty"`
}

type Asset struct {
	DownloadUrl string `json:"download_url,omitempty"`
	Path        string `json:"path,omitempty"`
	ID          string `json:"id,omitempty"`
	Repository  string `json:"repository,omitempty"`
	Format      string `json:"format,omitempty"`
}

func jsonUnmarshalAssetResponse(data []byte) (*AssetResponse, error) {
	var assetResponse AssetResponse
	if err := json.Unmarshal(data, &assetResponse); err != nil {
		return nil, fmt.Errorf("could not unmarshal assetResponse: %v", err)
	}
	return &assetResponse, nil
}

func jsonUnmarshalAsset(data []byte) (*Asset, error) {
	var asset Asset
	if err := json.Unmarshal(data, &asset); err != nil {
		return nil, fmt.Errorf("could not unmarshal Asset: %v", err)
	}
	return &asset, nil
}

func (c client) AssetRead(id string) (*Asset, error) {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", assetAPIEndpoint, id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("could not delete Asset '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	asset, err := jsonUnmarshalAsset(body)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (c client) AssetDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", assetAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete asset '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) AssetList(repository string) ([]Asset, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s?repository=%s", assetAPIEndpoint, repository), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s' asset list : HTTP: %d, %s",
			repository, resp.StatusCode, string(body))
	}

	assetResponse, err := jsonUnmarshalAssetResponse(body)
	if err != nil {
		return nil, err
	}

	list := assetResponse.Items
	for assetResponse.ContinuationToken != nil && assetResponse.ContinuationToken != "" {
		body, resp, err := c.Get(fmt.Sprintf("%s?repository=%s&continuationToken=%s", assetAPIEndpoint, repository, assetResponse.ContinuationToken), nil)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("could not read repository '%s' asset list : HTTP: %d, %s",
				repository, resp.StatusCode, string(body))
		}

		assetResponse, err := jsonUnmarshalAssetResponse(body)
		if err != nil {
			return nil, err
		}

		list = append(list, assetResponse.Items...)
	}

	return assetResponse.Items, nil
}
