provider "akamai" {
  edgerc = "../../test/edgerc"
}

resource "akamai_edge_hostname" "edgehostname" {
  contract            = "ctr_2"
  group               = "grp_2"
  product             = "prd_2"
  edge_hostname       = "test.akamaized.net"
  ip_behavior         = "IPV6_PERFORMANCE"
  status_update_email = []
}

output "edge_hostname" {
  value = akamai_edge_hostname.edgehostname.edge_hostname
}