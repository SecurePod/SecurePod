.PHONY: setup destory 

include .env

start:
	@make echo_green TEXT=".envファイルを`./frontend/.env.production`へコピーします";
	@cp .env ./frontend/.env.production;

	@make echo_green TEXT="演習コンテナをビルドします"
	./resources/build.sh

	@make echo_green TEXT="環境を構築します"
	@docker compose up -d --build

	@echo "演習環境へアクセスしてください: http://${NEXT_PUBLIC_API_URL}"

destory:
	@make  echo_green TEXT="Dockerのコンテナを削除します"
	@docker compose down

echo_green:
	@echo "\n\033[0;32m${TEXT}\033[0m\n"