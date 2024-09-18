package monitoring

import (
	"fmt"
	"net/http"
)

type MonitoringInterface interface {
	IncUsersOnline() error
	DecUsersOnline() error
	IncUsersSingUp() error
	IncProducersReserved() error
	IncInstrumentalistsReserved() error
	IncMicrophonesReserved() error
	IncGuitarsReserved() error
	IncCurrenciesReserved() error
}

type Client struct {
	httpClient *http.Client
	url        string
}

func New(url string) *Client {
	return &Client{
		httpClient: &http.Client{},
		url:        url,
	}
}

func madeRequest(method string, u string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request to inc users online: %w", err)
	}
	req.Header.Add("Method", method)

	return req, nil
}

func (c *Client) IncUsersOnline() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncUsersOnline"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest inc users online: %w", err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not inc users online: %w", err)
	}
	return nil
}

func (c *Client) DecUsersOnline() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "DecUsersOnline"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest inc users online: %w", err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not inc users online: %w", err)
	}
	return nil
}

func (c *Client) IncCurrentReserves() (err error) {
	method := "IncCurrentReserves"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest inc users online: %w", err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not inc users online: %w", err)
	}
	return nil
}

func (c *Client) DecCurrentReserves() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "DecCurrentReserves"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest inc users online: %w", err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not inc users online: %w", err)
	}
	return nil
}

func (c *Client) IncUsersSingUp() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncUsersSingUp"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest %s: %w", method, err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not %s: %w", method, err)
	}
	return nil
}

func (c *Client) IncProducersReserved() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncProducersReserved"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest %s: %w", method, err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not %s: %w", method, err)
	}
	return nil
}

func (c *Client) IncInstrumentalistsReserved() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncInstrumentalistsReserved"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest %s: %w", method, err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not %s: %w", method, err)
	}
	return nil
}

func (c *Client) IncMicrophonesReserved() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncMicrophonesReserved"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest %s: %w", method, err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not %s: %w", method, err)
	}
	return nil
}

func (c *Client) IncGuitarsReserved() (err error) {
	//return fmt.Errorf("no url provided %s", c.url)

	method := "IncGuitarsReserved"

	req, err := madeRequest(method, c.url)
	if err != nil {
		return fmt.Errorf("could not make reuqest %s: %w", method, err)
	}

	_, err = c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not %s: %w", method, err)
	}
	return nil
}
