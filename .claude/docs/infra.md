## Dockerfile 管理

多階段建構（`golang:1.25-alpine` → `alpine:3.20`），最終 image 不含 Go 工具鏈。`config.yaml` 必須複製進 image，ENTRYPOINT 帶 `-config config.yaml`。需要修改時先讀現有 `Dockerfile` 再編輯。

## Makefile 管理

所有 target 宣告 `.PHONY`，標準 targets：`run`、`build`、`test`、`lint`、`mysql-up`、`mysql-down`、`mysql-reset`。`mysql-*` 透過 `docker compose` 操作，不直接執行 `docker run`。需要修改時先讀現有 `Makefile` 再編輯。

## docker-compose 管理

專案根目錄，MySQL 8.0，具名 volume，使用 `clinicalmate` user（不用 root）。`config.yaml` database 區塊需與 compose 環境變數保持一致。需要修改時先讀現有 `docker-compose.yaml` 再編輯。
