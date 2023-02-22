locals {
  necessary_services = [
    "iap.googleapis.com",
    #    "networkmanagement.googleapis.com",
  ]
}

resource "google_project_service" "this" {
  for_each = toset(local.necessary_services)

  service = each.value

  project = var.gcp_project
}

