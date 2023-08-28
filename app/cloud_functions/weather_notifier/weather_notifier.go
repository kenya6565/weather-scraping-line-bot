package weather_notifier

import (
    "fmt"
    "net/http"
)

func WeatherNotifierFunction(w http.ResponseWriter, r *http.Request) {
    // TODO: Cloud Schedulerからのリクエストを受け取るためのロジックを記述
    fmt.Fprint(w, "Weather notifier function executed!")
}
