resource "null_resource" "build_lambda" {
  triggers = {
    build_trigger = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "cd ../app/cmd && GOOS=linux go build -o bootstrap && zip -j bootstrap.zip bootstrap"
  }
}

resource "aws_lambda_function" "weather_lambda" {
  function_name = "WeatherLambda"
  runtime       = "provided.al2"
  handler       = "bootstrap" # Lambda ハンドラ関数を指定

  # Lambda 関数のソースコードをZIPファイルとして指定
  filename         = "../app/cmd/bootstrap.zip"
  source_code_hash = filebase64sha256("../app/cmd/bootstrap.zip")

  # Lambda 関数に適切なロールを割り当て
  role = aws_iam_role.lambda_role.arn

  depends_on = [null_resource.build_lambda]
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
