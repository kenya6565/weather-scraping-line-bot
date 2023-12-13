resource "null_resource" "build_lambda" {
  triggers = {
    build_trigger = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "cd ../app/cmd && GOOS=linux go build -o bootstrap && zip -j bootstrap.zip bootstrap"
  }

  provisioner "local-exec" {
    when    = "destroy"
    command = "rm ../app/cmd/bootstrap.zip"
  }
}

resource "aws_lambda_function" "weather_lambda" {
  function_name = "WeatherLambda"
  runtime       = "provided.al2"
  handler       = "bootstrap"

  filename         = "../app/cmd/bootstrap.zip"
  source_code_hash = filebase64sha256("../app/cmd/bootstrap.zip")

  role = aws_iam_role.lambda_role.arn

  depends_on = [null_resource.build_lambda]
}
