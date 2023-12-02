resource "aws_lambda_function" "weather_lambda" {
  function_name = "WeatherLambda"
  runtime       = "provided.al2"
  handler       = "bootstrap" # Lambda ハンドラ関数を指定

  # Lambda 関数のソースコードをZIPファイルとして指定
  filename         = "path/to/your/bootstrap.zip"
  source_code_hash = filebase64sha256("path/to/your/bootstrap.zip")

  # Lambda 関数に適切なロールを割り当て
  role = aws_iam_role.lambda_role.arn
}

resource "aws_lambda_function" "line_webhook_lambda" {
  function_name = "LineWebhookLambda"
  runtime       = "provided.al2"
  handler       = "bootstrap" # Lambda ハンドラ関数を指定

  filename         = "path/to/your/bootstrap.zip"
  source_code_hash = filebase64sha256("path/to/your/bootstrap.zip")

  role = aws_iam_role.lambda_role.arn
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
      },
    ],
  })
}

# 必要なIAMポリシーのアタッチ
resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
