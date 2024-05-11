package redis

import (
	"crypto/tls"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

// Open a connection to the Redis server.
func NewClient(addr string, tlsConfig *tls.Config) (*Client, error) {

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return nil, err
	}

	return &Client{conn: conn}, nil
}

// Open a connection to the Redis server with TLS/SSL already enabled.
func NewPreConfigClient(addr string) (*Client, error) {

	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return nil, err
	}

	return &Client{conn: conn}, nil
}

// Send an authentication request to the Redis server.
func (c *Client) Auth(username, password string) (string, error) {
	_, err := c.conn.Write([]byte(fmt.Sprintf("AUTH %s %s\r\n", username, password)))
	if err != nil {
		return "Failed to send AUTH command", err
	}

	response := make([]byte, 1024)
	n, err := c.conn.Read(response)
	if err != nil {
		return "Failed to read response", err
	}

	return string(response[:n]), nil
}

// Close the connection to the Redis server.
func (c *Client) Close() error {
	return c.conn.Close()
}
