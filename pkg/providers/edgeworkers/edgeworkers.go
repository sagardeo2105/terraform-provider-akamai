//go:build all || edgeworkers
// +build all edgeworkers

package edgeworkers

import "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/registry"

func init() {
	registry.RegisterProvider(Subprovider())
}
