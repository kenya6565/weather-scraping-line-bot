provider "google" {
  credentials = file(var.service_account_key_path)
  project     = var.gcp_project_id
  region      = "us-central1"
}
