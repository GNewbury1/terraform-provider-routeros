package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPAddress struct {
	ID              string `json:".id,omitempty"`
	ActualInterface string `json:"actual-interface,omitempty"`
	Address         string `json:"address"`
	Comment         string `json:"comment,omitempty"`
	Disabled        string `json:"disabled"`
	Dynamic         string `json:"dynamic,omitempty"`
	Interface       string `json:"interface"`
	Invalid         string `json:"invalid,omitempty"`
	Network         string `json:"network"`
}

func (c *Client) GetIPAddress(id string) (*IPAddress, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/ip/address/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateIPAddress(ip_address *IPAddress) (*IPAddress, error) {
	reqBody, err := json.Marshal(ip_address)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rest/ip/address", c.HostURL), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateIPAddress(id string, ip_address *IPAddress) (*IPAddress, error) {
	reqBody, err := json.Marshal(ip_address)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/rest/ip/address/%s", c.HostURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteIPAddress(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/rest/ip/address/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	res := IPAddress{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	return nil
}
