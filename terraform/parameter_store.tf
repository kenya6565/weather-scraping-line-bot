resource "aws_ssm_parameter" "line_channel_secret" {
  name  = "/app/line_channel_secret"
  type  = "SecureString"
  value = var.line_channel_secret
}

resource "aws_ssm_parameter" "line_access_token" {
  name  = "/app/line_access_token"
  type  = "SecureString"
  value = var.line_access_token
}

resource "aws_ssm_parameter" "firebase_project_id" {
  name  = "/app/firebase_project_id"
  type  = "SecureString"
  value = var.firebase_project_id
}

resource "aws_ssm_parameter" "google_application_credentials" {
  name  = "/app/google_application_credentials"
  type  = "SecureString"
  value = jsonencode({
    type = var.type,
    project_id = var.project_id,
    private_key_id = var.private_key_id,
    private_key = var.private_key,
    client_email = var.client_email,
    client_id = var.client_id,
    auth_uri = var.auth_uri,
    token_uri = var.token_uri,
    auth_provider_x509_cert_url = var.auth_provider_x509_cert_url,
    client_x509_cert_url = var.client_x509_cert_url,
    universe_domain = var.universe_domain
  })
}
