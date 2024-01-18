// enable internet access from gcp compute instances
resource "google_compute_router" "this" {
  name = "${var.name}-nat-router"

  project = var.gcp_project
  region  = var.gcp_region

  network = google_compute_network.prometheus.name
}

resource "google_compute_router_nat" "this" {
  name   = "${var.name}-nat"
  router = google_compute_router.this.name

  project = var.gcp_project
  region  = var.gcp_region

  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}
