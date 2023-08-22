# Cloud Functionsのリソース
resource "google_cloudfunctions_function" "weather_notifier" {
  name                  = "weather-notifier"
  available_memory_mb   = 256
  runtime               = "go113"
  source_archive_bucket = google_storage_bucket.cloudfunctions_bucket.name
  source_archive_object = google_storage_bucket_object.source_archive.name

  trigger_http = true

  entry_point = "WeatherNotifierFunction"
}

resource "google_storage_bucket" "cloudfunctions_bucket" {
  name     = "weather-notifier-functions-bucket"
  location = "us-central1"
}

resource "google_storage_bucket_object" "source_archive" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.cloudfunctions_bucket.name
  source = "${path.module}/../function-source.zip"
}

# Cloud Schedulerのジョブ
resource "google_cloud_scheduler_job" "daily_weather_check" {
  name     = "daily-weather-check"
  schedule = "0 0 * * *" # 毎日0時に実行

  http_target {
    uri                  = google_cloudfunctions_function.weather_notifier.https_trigger_url
    http_method          = "GET"
  }
}

# Firestoreの設定は事前に終えていたのと今後変更する事がないのでterraformに含まない
