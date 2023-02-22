locals {
  firewall_rules = {
    "default" : {
      allow_rules : [
      ]
      deny_rules : [
        {
          protocol : "icmp"
        }
      ]
      source_ranges : [
        "0.0.0.0/0"
      ]
    },
    "ssh-with-iap" : {
      allow_rules : [
        {
          protocol : "tcp"
          ports : ["22"]
        }
      ],
      deny_rules : [],
      source_ranges : [
        // art ip
        "81.198.100.65/32",
        // iap address range (https://cloud.google.com/iap/docs/using-tcp-forwarding#create-firewall-rule)
        "35.235.240.0/20"
      ]
    },
    "prometheus" : {
      name : "prometheus",
      allow_rules : [
        {
          protocol : "tcp"
          ports : ["9100"]
        }
      ],
      deny_rules : [],
      source_ranges : [
        "81.198.100.65/32",
      ]
    }
  }
}

resource "google_compute_firewall" "this" {
  for_each = local.firewall_rules

  name    = each.key
  network = google_compute_network.prometheus.name

  project = var.gcp_project

  dynamic "allow" {
    for_each = each.value.allow_rules

    content {
      protocol = allow.value["protocol"]
      ports    = can(allow.value["ports"]) ? allow.value["ports"] : []
    }
  }

  dynamic "deny" {
    for_each = each.value.deny_rules
    content {
      protocol = deny.value["protocol"]
      ports    = can(deny.value["ports"]) ? deny.value["ports"] : []
    }
  }

  source_ranges = each.value.source_ranges
}
