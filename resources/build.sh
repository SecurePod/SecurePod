#!/bin/bash

# トップディレクトリを設定（このスクリプトが置かれているディレクトリ）
top_dir=$(pwd)

# findコマンドを使用して、すべてのdocker-compose.ymlファイルの場所を探す
find $top_dir -name 'docker-compose.yml' -print0 | while IFS= read -r -d '' file; do
    # docker-compose.ymlが見つかったディレクトリに移動
    dir=$(dirname "$file")
    echo "Building in $dir"
    cd "$dir"
    # docker compose build を実行
    docker compose build
    # 元のディレクトリに戻る
    cd $top_dir
done
