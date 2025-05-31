# Stage 1: Build Next.js app
FROM node:20-alpine AS build

# Đặt thư mục làm việc trong container
WORKDIR /app

# Copy package.json và lock file
COPY package*.json ./

# Cài đặt dependencies
RUN npm install

# Copy toàn bộ mã nguồn
COPY . .

# Build ứng dụng Next.js (dựa theo next.config.js sẽ xuất ra thư mục 'out')
RUN npm run build

# Stage 2: Serve static site bằng Nginx
FROM nginx:alpine

# Copy output từ stage build
COPY --from=build /app/out /usr/share/nginx/html

# Copy file cấu hình Nginx (nếu bạn có)
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Khởi động nginx khi container chạy
CMD ["nginx", "-g", "daemon off;"]
