package quotas

import (
	"github.com/DashuOps/gophercloud"
)

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
