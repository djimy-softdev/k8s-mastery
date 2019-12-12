FROM node:alpine as builder
WORKDIR /var
COPY . .
RUN npm install && npm run build 

FROM nginx:alpine
COPY --from=builder /var/build /usr/share/nginx/html