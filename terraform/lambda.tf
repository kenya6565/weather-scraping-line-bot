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
  function_name = "WeatherLambda-${terraform.workspace}"
  runtime       = "provided.al2"
  handler       = "bootstrap"

  filename      = "../app/cmd/bootstrap.zip"

  role = aws_iam_role.lambda_role.arn

  depends_on = [null_resource.build_lambda]
  environment {
    variables = {
      AWS_EXECUTION_ENV = "AWS_Lambda"
      ENVIRONMENT       = terraform.workspace
    }
  }

  # loggroupを作成する
  tracing_config {
    mode = "PassThrough"
  }
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role-${terraform.workspace}"

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
  name        = "SSMGetParameter-${terraform.workspace}"
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

# api gatewayからLambda関数の呼び出しを許可する権限の追加
resource "aws_lambda_permission" "api_gateway_invoke_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.weather_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.line_webhook_api.execution_arn}/*/*"
}
