resource "aws_cloudwatch_event_rule" "daily_trigger" {
  name        = "DailyTrigger-${terraform.workspace}"
  description = "Trigger at 24:00 JST daily"
  schedule_expression = "cron(0 15 * * ? *)" # UTCで15時 (JSTで24時)
}

# 対象のlambda関数を指定
resource "aws_cloudwatch_event_target" "trigger_weather_lambda" {
  rule = aws_cloudwatch_event_rule.daily_trigger.name
  arn  = aws_lambda_function.weather_lambda.arn
  # リソースに変更があった場合のみ再作成する
  lifecycle {
    create_before_destroy = true
  }
}

# EventBridgeからLambda関数を呼び出すためのIAM権限を設定
resource "aws_lambda_permission" "allow_cloudwatch_to_call_weather_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch-${terraform.workspace}"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.weather_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.daily_trigger.arn
}
