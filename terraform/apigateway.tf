resource "aws_api_gateway_rest_api" "line_webhook_api" {
  name        = "LineWebhookAPI"
  description = "API for LINE webhook"
}

resource "aws_api_gateway_resource" "line_webhook_resource" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id
  parent_id   = aws_api_gateway_rest_api.line_webhook_api.root_resource_id
  path_part   = "webhook"
}

# /webhookエンドポイントにPOSTリクエストが送られると、設定したLambda関数が実行
resource "aws_api_gateway_method" "line_webhook_method" {
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id   = aws_api_gateway_resource.line_webhook_resource.id
  http_method   = "POST"
  authorization = "NONE"
}
