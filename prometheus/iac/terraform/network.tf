resource "google_compute_network" "prometheus" {
  name                    = var.name
  auto_create_subnetworks = false

  project = var.gcp_project
}

resource "google_compute_subnetwork" "prometheus" {
  ip_cidr_range = "192.168.0.0/24"
  name          = "${var.name}-nodes-w-exporter"
  network       = google_compute_network.prometheus.id

  project = var.gcp_project
  region  = var.gcp_region
}

