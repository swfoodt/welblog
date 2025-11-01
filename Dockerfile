FROM nginx:alpine

# 复制本地构建好的文件
COPY ./public /usr/share/nginx/html

# 复制nginx配置
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]