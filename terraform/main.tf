resource "google_storage_bucket" "cloudfunctions_bucket" {
  name     = "weather-notifier-functions-bucket"
  location = "us-central1"
}

resource "google_storage_bucket_object" "source_archive" {
  name   = "cloud_functions.zip"
  bucket = google_storage_bucket.cloudfunctions_bucket.name
  source = "${path.module}/../app/workspace/serverless_function_source_code/cloud_functions.zip"
}

# Cloud SchedulerからのHTTPリクエストを発火点にAPIを叩いて通知を行うCloud Functions
resource "google_cloudfunctions_function" "weather_notifier" {
  name                  = "WeatherNotifierFunction"
  available_memory_mb   = 256
  runtime               = "go120"
  source_archive_bucket = google_storage_bucket.cloudfunctions_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive.name
  trigger_http          = true
  entry_point           = "WeatherNotifierFunction"
  environment_variables = {
  }
}

# LINE webhookからのPOSTリクエストを発火点にフォローなどのイベントがあった際に通知を行うCloud Functions
resource "google_cloudfunctions_function" "line_webhook" {
  name                  = "LineWebhookFunction"
  available_memory_mb   = 256
  runtime               = "go120"
  source_archive_bucket = google_storage_bucket.cloudfunctions_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive.name
  trigger_http          = true
  entry_point           = "LineWebhookFunction"
  environment_variables = {
  }
}

# Cloud Schedulerのジョブ
resource "google_cloud_scheduler_job" "daily_weather_check" {
  name     = "daily-weather-check"
  schedule = "0 0 * * *"
  http_target {
    uri                  = google_cloudfunctions_function.weather_notifier.https_trigger_url
    http_method          = "GET"
  }
}

# Firestoreの設定は事前に終えていたのと今後変更する事がないのでterraformに含まない
