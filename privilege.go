package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	privilegeAPIEndpoint = "service/rest/beta/security/privileges"
)

// Privilege data
type Privilege struct {
	Actions         []string `json:"actions"`
	ContentSelector string   `json:"contentSelector"`
	Description     string   `json:"description"`
	Domain          string   `json:"domain"`
	Format          string   `json:"string"`
	Name            string   `json:"name"`
	ReadOnly        bool     `json:"readOnly"`
	Repository      string   `json:"repository"`
	Type            string   `json:"type"`
}

func (c client) PrivilegeCreate(p Privilege) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s", privilegeAPIEndpoint, strings.ToLower(p.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) PrivilegeRead(name string) (*Privilege, error) {
	body, resp, err := c.Get(privilegeAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read privileges: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var privileges []Privilege
	if err := json.Unmarshal(body, &privileges); err != nil {
		return nil, fmt.Errorf("could not unmarshal privileges \"%s\": %v", name, err)
	}

	for _, p := range privileges {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, nil
}

func (c client) PrivilegeUpdate(name string, p Privilege) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s", privilegeAPIEndpoint, p.Type, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) PrivilegeDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", privilegeAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete privilege \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
