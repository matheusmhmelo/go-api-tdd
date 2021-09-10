package credential

import "errors"

type Credentials struct {
	ClientID            string `json:"client_id"`
	OrganizationID      string `json:"organization_id"`
	SoftwareStatementID string `json:"software_statement_id"`
}

func (c *Credentials) Validate() error {
	if c.ClientID == "" || c.OrganizationID == "" || c.SoftwareStatementID == "" {
		return errors.New("all values are required")
	}
	return nil
}
