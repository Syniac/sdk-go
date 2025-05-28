package syniac

import "encoding/json"

type Instance struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Region   string `json:"region"`
	Image    string `json:"image"`
	InstanceType string `json:"instance_type"`
}

func (c *Client) ListInstances() ([]Instance, error) {
	data, err := c.request("GET", "/v1/instances", nil)
	if err != nil {
		return nil, err
	}

	var instances []Instance
	err = json.Unmarshal(data, &instances)
	if err != nil {
		return nil, err
	}

	return instances, nil
}

func (c *Client) GetInstance(id string) (*Instance, error) {
	data, err := c.request("GET", "/v1/instances/" + id, nil)
	if err != nil {
		return nil, err
	}

	var instance Instance
	err = json.Unmarshal(data, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
