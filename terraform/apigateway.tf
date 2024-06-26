resource "aws_api_gateway_rest_api" "line_webhook_api" {
  name        = "LineWebhookAPI-${terraform.workspace}"
  description = "API for LINE webhook in ${terraform.workspace} environment"
}

# ユーザーからのwebhook event(フォローやメッセージ)を受け取るためのエンドポイントの設定
resource "aws_api_gateway_resource" "line_webhook_resource" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id
  parent_id   = aws_api_gateway_rest_api.line_webhook_api.root_resource_id
  path_part   = "webhook"
}

resource "aws_api_gateway_method" "line_webhook_method" {
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id   = aws_api_gateway_resource.line_webhook_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

# line webhookから受け取ったeventをlambda関数に渡すための設定
resource "aws_api_gateway_integration" "line_webhook_integration" {
  rest_api_id             = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id             = aws_api_gateway_resource.line_webhook_resource.id
  http_method             = aws_api_gateway_method.line_webhook_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.weather_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "deployment" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id

  depends_on = [
    aws_api_gateway_integration.line_webhook_integration,
    aws_api_gateway_method.line_webhook_method,
  ]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "stage" {
  stage_name    = terraform.workspace
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  deployment_id = aws_api_gateway_deployment.deployment.id
}
