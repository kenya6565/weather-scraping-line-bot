#!/bin/bash
# .env ファイルから環境変数を読み込み、それらを TF_VAR_ で始まる環境変数に設定
export $(grep -v '^#' app/cmd/.env | xargs -d '\n' -I {} echo "TF_VAR_{}" | sed 's/=/=/g')
