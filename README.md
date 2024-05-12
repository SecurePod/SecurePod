# SecurePod

## 準備

プライベートIPv4アドレスの確認

Mac, Linuxの場合、`ifconfig`コマンド、Windowsの場合、`ipconfig`コマンドを実行して、プライベートIPv4アドレスを確認してください。

確認後、.env.productionファイルをコピーして.envファイルを作成し、確認したプライベートIPv4アドレスを設定してください。

```bash
# 例
NEXT_PUBLIC_API_URL="192.168.3.2"
NEXT_PUBLIC_API_PORT="8080"
```

これで準備は完了です

## アプリケーションの起動

Mac, Linuxの場合、以下のコマンドを実行してください。

```bash
make start
```

Windowsの場合以下のスクリプトを実行してください。

```bash
```
