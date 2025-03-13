
resource "google_compute_network" "vpc" {
  name                    = var.vpc_name
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "general_subnet" {
  name          = "asia-seoul-subnet"
  ip_cidr_range = var.subnet_cidr
  region        = var.region
  network       = google_compute_network.vpc.id
}

resource "google_compute_subnetwork" "connector_subnet" {
  name    = "serverless-vpc-connector"
  region  = var.region
  network = google_compute_network.vpc.id
}

resource "google_compute_firewall" "allow_internal" {
  name    = "allow-internal"
  network = google_compute_network.vpc.name
  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "udp"
    ports    = ["0-65535"]
  }

  allow {
    protocol = "icmp"
    ports    = ["0-65535"]
  }
  source_ranges = [var.subnet_cidr, var.connector_cidr]
}

resource "google_compute_firewall" "allow_connector" {
  name      = "allow-connector"
  network   = google_compute_network.vpc.name
  direction = "INGRESS"
  allow {
    protocol = "tcp"
    ports    = ["6379"]
  }

  source_ranges = [var.connector_cidr]
}


resource "google_compute_firewall" "allow_health_checks" {
  name    = "allow-health-checks"
  network = google_compute_network.vpc.name
  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }
  source_ranges = ["35.191.0.0/16", "130.211.0.0/22"] # Google Health check ranges
}

resource "google_vpc_access_connector" "connector" {
  name   = "asia-seoul-access-connector"
  region = var.region
  subnet {
    name = google_compute_subnetwork.connector_subnet.name
  }
}

resource "google_redis_instance" "redis" {
  name               = "redis"
  memory_size_gb     = var.redis_size
  region             = var.region
  authorized_network = google_compute_network.vpc.id
}


resource "google_artifact_registry_repository" "repo" {
  location      = var.region
  repository_id = var.artifact_registry_repo_name
  format        = "DOCKER"
}

resource "google_service_account" "github_deployer" {
  account_id   = "github-actions-deployer"
  display_name = "Github Action Deployer"
}

resource "google_project_iam_member" "run_admin" {
  project = var.project_id
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.github_deployer.email}"
}


resource "google_project_iam_member" "artifact_registry_writer" {
  project = var.project_id
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.github_deployer.email}"
}

resource "google_project_iam_member" "service_account_user" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.github_deployer.email}"
}

resource "google_project_iam_member" "storage_admin" {
  project = var.project_id
  role    = "roles/storage.admin"
  member  = "serviceAccount:${google_service_account.github_deployer.email}"
}

resource "google_service_account_key" "github_deployer_key" {
  service_account_id = google_service_account.github_deployer.name
}
