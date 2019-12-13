FROM node:alpine as builder
WORKDIR /var
COPY . .
RUN npm install && npm run build 

FROM nginx:alpine
COPY --from=builder /var/build /usr/share/nginx/html
COPY  ./backend.conf.template /etc/nginx/conf.d/
ENTRYPOINT /bin/sh -c "envsubst < /etc/nginx/conf.d/backend.conf.template > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"