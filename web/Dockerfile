FROM node:18-alpine as builder

# Build Time Environment
ARG REACT_APP_RENTHOME_GOOGLE_OAUTH_CLIENT_ID
ARG REACT_APP_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID
ARG REACT_APP_RENTHOME_API_ADDR

# Set working directory
WORKDIR /app

# Install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm ci

# Add source code
COPY ./ ./

# Build
RUN npm run build

FROM nginx:1.20.0-alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /app/build ./
ENTRYPOINT [ "nginx", "-g", "daemon off;" ]

