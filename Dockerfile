# ────────────────────────────────────────────
# Stage 1：編譯階段（builder）
# 使用完整的 Go 工具鏈編譯出靜態二進位檔
# ────────────────────────────────────────────

# GO_VERSION 由 CI/CD 的 build-args 傳入（預設 1.26），統一版本管理
# 要升版只需修改 ci-cd.yml 的 env.GO_VERSION，Dockerfile 不需動
ARG GO_VERSION=1.26

# 以官方 Go Alpine 映像作為編譯基底，版本由上方 ARG 決定
FROM golang:${GO_VERSION}-alpine AS builder

# 設定容器內的工作目錄，後續所有指令都在此路徑下執行
WORKDIR /app

# 先只複製 go.mod 與 go.sum，利用 Docker 快取層機制：
# 只要這兩個檔案沒變，下一行的 go mod download 就不會重新執行
COPY go.mod go.sum ./

# 依照 go.sum 下載並驗證所有相依套件，寫入模組快取
RUN go mod download

# 將專案所有原始碼複製進容器（排除 .dockerignore 列出的項目）
COPY . .

# 編譯 Go 程式：
#   CGO_ENABLED=0  停用 CGO，產生純靜態二進位，不依賴 C 函式庫
#   GOOS=linux     跨平台編譯目標設為 Linux
#   -o server      輸出執行檔命名為 server
#   .              編譯當前目錄（main package）
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ────────────────────────────────────────────
# Stage 2：執行階段（runtime）
# 只保留執行所需的最小化映像，降低映像體積與攻擊面
# ────────────────────────────────────────────

# 以最小化的 Alpine 3.20 作為執行基底（不含 Go 工具鏈，映像更小）
FROM alpine:3.20

# 安裝時區資料（tzdata），確保應用程式能正確處理時區設定
# --no-cache 避免在映像內留下 APK 快取，進一步縮小體積
RUN apk add --no-cache tzdata

# 設定執行階段的工作目錄
WORKDIR /app

# 從 builder 階段複製編譯好的執行檔到執行映像
COPY --from=builder /app/server .

# 從 builder 階段複製 SQL migration 檔案，供啟動時執行資料庫版本控制
COPY --from=builder /app/migrations ./migrations

# 宣告容器對外開放的埠號（8000），供 docker run -p 或 docker-compose 對應
EXPOSE 8000

# 容器啟動時執行的預設指令：直接執行編譯好的 server 二進位檔
CMD ["./server"]
