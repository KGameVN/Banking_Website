import { NextConfig } from 'next';

const nextConfig: NextConfig = {
  reactStrictMode: true, // Bật strict mode cho ứng dụng React
  output: 'export', // Chạy Next.js ở chế độ standalone nếu cần, cho Nginx hoặc khi xuất ứng dụng tĩnh
  images: {
    unoptimized: true, // Bỏ qua tối ưu hóa ảnh khi export tĩnh
  },

  // Tạo environment variables, hỗ trợ chuyển đổi giữa môi trường dev và prod
  env: {
    DB_HOST: process.env.DB_HOST || 'localhost',
    DB_PORT: process.env.DB_PORT || '5432',
    DB_USER: process.env.DB_USER || 'user',
    DB_PASSWORD: process.env.DB_PASSWORD || 'password',
    DB_NAME: process.env.DB_NAME || 'mydatabase',
  },

  // Cấu hình rewrites hoặc redirects nếu cần
  async redirects() {
    return [
      {
        source: '/old-path',
        destination: '/new-path',
        permanent: true,
      },
    ];
  },

  // Tùy chỉnh webpack nếu cần
  webpack(config, { isServer }) {
    if (!isServer) {
      config.resolve.fallback = {
        fs: false,
        path: false,
      };
    }
    return config;
  },

  // Tùy chọn để cấu hình ứng dụng phụ thuộc vào môi trường (dev hoặc prod)
  async headers() {
    return [
      {
        source: '/(.*)',
        headers: [
          {
            key: 'X-Content-Type-Options',
            value: 'nosniff',
          },
        ],
      },
    ];
  },
};

export default nextConfig;
