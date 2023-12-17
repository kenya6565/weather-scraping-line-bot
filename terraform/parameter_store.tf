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
