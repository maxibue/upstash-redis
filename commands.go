package redis

import (
	"fmt"
	"strings"
)

// Send a PING command to the Redis server.
func (c *Client) Ping() (string, error) {
	_, err := c.conn.Write([]byte("PING\r\n"))
	if err != nil {
		return "", err
	}

	response := make([]byte, 1024)
	n, err := c.conn.Read(response)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(response[:n]), "+"), nil
}

// Get the value of a key from the Redis server.
func (c *Client) Get(key string) (string, error) {
	_, err := c.conn.Write([]byte(fmt.Sprintf("GET %s \r\n", key)))
	if err != nil {
		return "", err
	}

	getResponse := make([]byte, 1024)
	n, err := c.conn.Read(getResponse)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(strings.SplitAfter(string(getResponse[:n]), "\r\n")[1], "\"", ""), nil
}

// Set the value of a key on the Redis server. If the first return value is not "OK", an error occurred.
func (c *Client) Set(key, value string) (string, error) {
	setCmd := fmt.Sprintf("SET %s \"%s\" \r\n", key, strings.ReplaceAll(value, " ", ""))
	_, err := c.conn.Write([]byte(setCmd))
	if err != nil {
		return "", err
	}

	setResponse := make([]byte, 1024)
	n, err := c.conn.Read(setResponse)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(setResponse[:n]), "+"), nil
}

// Set the value of a key on the Redis server with an expiration time.
func (c *Client) SetWithEx(key, value string, expiration int) (string, error) {
	setCmd := fmt.Sprintf("SET %s \"%s\" EX %v \r\n", key, strings.ReplaceAll(value, " ", ""), expiration)
	_, err := c.conn.Write([]byte(setCmd))
	if err != nil {
		return "", err
	}

	setResponse := make([]byte, 1024)
	n, err := c.conn.Read(setResponse)
	if err != nil {
		return "", err
	}

	return string(setResponse[:n]), nil
}

// Get the time-to-live of a key on the Redis server.
func (c *Client) TTL(key string) (string, error) {
	_, err := c.conn.Write([]byte(fmt.Sprintf("TTL %s \r\n", key)))
	if err != nil {
		return "", err
	}

	getResponse := make([]byte, 1024)
	n, err := c.conn.Read(getResponse)
	if err != nil {
		return "", err
	}

	return strings.TrimLeft(string(getResponse[:n]), ":"), nil
}
