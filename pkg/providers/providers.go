package providers

import (
	// This is where providers are import so they can register themselves
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/appsec"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/botman"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/cloudlets"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/cps"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/datastream"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/dns"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/edgeworkers"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/gtm"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/iam"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/imaging"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/networklists"
	_ "github.com/akamai/terraform-provider-akamai/v2/pkg/providers/property"
)
