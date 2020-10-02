# 開発環境Dockerコンテナを起動する
.PHONY: devenv
devenv:
	./build/develop/exec.sh

# コードやドキュメント等の生成を行うタスクを実行する
.PHONY: generate
generate:
	./scripts/go_generate.sh
