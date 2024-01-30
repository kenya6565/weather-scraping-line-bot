resource "null_resource" "build_lambda" {
  triggers = {
    build_trigger = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = <<EOF
      cd ../app/cmd
      GOOS=linux go build -o bootstrap
      zip -j bootstrap.zip bootstrap
    EOF
  }
}

resource "aws_lambda_function" "weather_lambda" {
  function_name = "WeatherLambda"
  runtime       = "provided.al2"
  handler       = "bootstrap"

  filename      = "../app/cmd/bootstrap.zip"

  role = aws_iam_role.lambda_role.arn

  depends_on = [null_resource.build_lambda]
  environment {
    variables = {
      AWS_EXECUTION_ENV = "AWS_Lambda"
    }
  }

  # loggroupを作成する
  tracing_config {
    mode = "PassThrough"
  }
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "ssm_get_parameter" {
  name        = "SSMGetParameter"
  description = "Allow lambda function to get parameters from SSM"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "ssm:GetParameter",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ssm_get_parameter_attach" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.ssm_get_parameter.arn
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Lambda関数をトリガーするエンドポイントを作成
resource "aws_api_gateway_resource" "line_message_resource" {
  rest_api_id = aws_api_gateway_rest_api.line_webhook_api.id
  parent_id   = aws_api_gateway_rest_api.line_webhook_api.root_resource_id
  path_part   = "message"
}

resource "aws_api_gateway_method" "line_message_method" {
  rest_api_id   = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id   = aws_api_gateway_resource.line_message_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "line_message_integration" {
  rest_api_id             = aws_api_gateway_rest_api.line_webhook_api.id
  resource_id             = aws_api_gateway_resource.line_message_resource.id
  http_method             = aws_api_gateway_method.line_message_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.weather_lambda.invoke_arn
}
