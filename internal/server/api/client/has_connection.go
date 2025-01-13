package client

// HasConnection проверка наличия подключения.
func (c Client) HasConnection() bool {
	_, err := c.Get(c.Host + "/ping")
	return err == nil
}
