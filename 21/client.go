package main

type Client struct {
}

func (c *Client) insertSquareUsbInComputer(com Computer) {
	com.insertInSquarePort()
}
