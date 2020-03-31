package simplivity

import (
	"github.com/HewlettPackard/simplivity-go/ovc"
)

type Config struct {
	OVCIP           string
	Username        string
	Password        string
	CertificatePath string
	Client          *ovc.Client
}

// SetClient sets the SimpliVity OVC client.
func (c *Config) SetClient() error {
	client, err := ovc.NewClient(c.Username, c.Password, c.OVCIP, c.CertificatePath)
	if err != nil {
		return err
	}

	c.Client = client
	return nil
}
