package client

func (c *Client) SearchLocation(lp *LocationParam) ([]*Location, error) {
	resp, err := c.do(lp.Encode())
	if err != nil {
		return nil, err
	}

	return resp.Locations, nil
}

func (c *Client) QueryLiveWeather(lwp *LiveWeatherParam) (*Now, error) {
	resp, err := c.do(lwp.Encode())
	if err != nil {
		return nil, err
	}

	return resp.Now, nil
}
