resource "aws_api_gateway_rest_api" "line_webhook_api" {
  name        = "LineWebhookAPI"
  description = "API for LINE webhook"
}

resource "aws_api_gateway_resource" "line_webhook_resource" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id
  parent_id   = aws_api_gateway_rest_api.line_webhook_api.root_resource_id
  path_part   = "webhook"
}

resource "aws_api_gateway_integration" "line_webhook_integration" {
  rest_api_id             = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id             = aws_api_gateway_resource.line_webhook_resource.id
  http_method             = aws_api_gateway_method.line_webhook_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.weather_lambda.invoke_arn
}

# /webhookエンドポイントにPOSTリクエストが送られると、設定したLambda関数が実行
resource "aws_api_gateway_method" "line_webhook_method" {
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id   = aws_api_gateway_resource.line_webhook_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_deployment" "prd_deployment" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "prd_stage" {
  stage_name    = "prd"
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  deployment_id = aws_api_gateway_deployment.prd_deployment.id
}
