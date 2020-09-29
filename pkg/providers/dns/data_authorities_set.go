package dns

import (
	"context"
	"fmt"
	"sort"
	"strings"

	dnsv2 "github.com/akamai/AkamaiOPEN-edgegrid-golang/configdns-v2"
	"github.com/akamai/terraform-provider-akamai/v2/pkg/akamai"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAuthoritiesSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAuthoritiesSetRead,
		Schema: map[string]*schema.Schema{
			"contract": {
				Type:     schema.TypeString,
				Required: true,
			},
			"authorities": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceAuthoritiesSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := akamai.Meta(m)
	logger := meta.Log("[Akamai DNS]", "dataSourceDNSAuthoritiesRead")

	contractid := strings.TrimPrefix(d.Get("contract").(string), "ctr_")
	// Warning or Errors can be collected in a slice type
	var diags diag.Diagnostics

	logger.WithField("contractid", contractid).Debug("Start Searching for authority records")

	ns, err := dnsv2.GetNameServerRecordList(contractid)
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("error looking up ns records for %s", contractid),
			Detail:   err.Error(),
		})
	}
	logger.WithField("records", ns).Debug("Searching for records")

	sort.Strings(ns)
	d.Set("authorities", ns)
	d.SetId(contractid)

	return diags
}