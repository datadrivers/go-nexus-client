package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	assetAPIEndpoint = "service/rest/v1/assets"
)

type AssetList struct {
	Items             []Asset `json:"items,omitempty"`
	ContinuationToken string  `json:"continuationToken,omitempty"`
}

type Asset struct {
	DownloadUrl string `json:"downloadUrl,omitempty"`
	Path        string `json:"path,omitempty"`
	ID          string `json:"id,omitempty"`
	Repository  string `json:"repository,omitempty"`
	Format      string `json:"format,omitempty"`
}

func jsonUnmarshalAssetList(data []byte) (*AssetList, error) {
	var assetList AssetList
	if err := json.Unmarshal(data, &assetList); err != nil {
		return nil, fmt.Errorf("could not unmarshal assetList: %v", err)
	}
	return &assetList, nil
}

func jsonUnmarshalAsset(data []byte) (*Asset, error) {
	var asset Asset
	if err := json.Unmarshal(data, &asset); err != nil {
		return nil, fmt.Errorf("could not unmarshal Asset: %v", err)
	}
	return &asset, nil
}

func (c client) AssetRead(id string) (*Asset, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/%s", assetAPIEndpoint, id), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("could not read Asset '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
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

	assetList, err := jsonUnmarshalAssetList(body)
	if err != nil {
		return nil, err
	}

	list := assetList.Items
	for assetList.ContinuationToken != "" {
		body, resp, err := c.Get(fmt.Sprintf("%s?repository=%s&continuationToken=%s", assetAPIEndpoint, repository, assetList.ContinuationToken), nil)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("could not read repository '%s' asset list : HTTP: %d, %s",
				repository, resp.StatusCode, string(body))
		}

		assetResponse, err := jsonUnmarshalAssetList(body)
		if err != nil {
			return nil, err
		}

		list = append(list, assetResponse.Items...)
	}

	return assetList.Items, nil
}
