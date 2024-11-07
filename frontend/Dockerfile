# 使用 Node.js 作為基礎映像
FROM node:18-alpine AS build

# 設定工作目錄
WORKDIR /app

# 複製 package.json 和 package-lock.json 並安裝依賴
COPY package.json package-lock.json ./
RUN npm install

# 複製所有檔案並進行建置
COPY . .
RUN npm run build

# 使用輕量級的 nginx 映像來提供服務
FROM nginx:alpine

# 複製 build 出來的檔案到 nginx 的靜態檔案目錄
COPY --from=build /app/dist /usr/share/nginx/html

# 暴露端口 80
EXPOSE 80

# 啟動 nginx
CMD ["nginx", "-g", "daemon off;"]
