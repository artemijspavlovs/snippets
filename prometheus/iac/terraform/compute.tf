resource "google_service_account" "compute" {
  account_id   = var.name
  display_name = "Prometheus Node"

  project = var.gcp_project
}

// to install ops agent: https://registry.terraform.io/modules/terraform-google-modules/cloud-operations/google/latest/submodules/agent-policy#usage
resource "google_compute_instance" "prometheus_node" {
  machine_type = "e2-micro"
  name         = "${var.name}-node-w-exporter"

  project = var.gcp_project
  zone    = "${var.gcp_region}-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-2204-jammy-v20230214"
    }
  }

  network_interface {
    network = google_compute_network.prometheus.name

    subnetwork         = google_compute_subnetwork.prometheus.id
    subnetwork_project = google_compute_subnetwork.prometheus.project

    access_config {}
  }

  service_account {
    email  = google_service_account.compute.email
    scopes = ["cloud-platform"]
  }
}
