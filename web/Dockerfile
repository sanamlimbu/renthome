FROM node:18-alpine as builder

# Set working directory
WORKDIR /app

# Build environment variables
ARG VITE_RENTHOME_API_ADDRESS
ENV VITE_RENTHOME_API_ADDRESS ${VITE_RENTHOME_API_ADDRESS}

ARG VITE_RENTHOME_PUBLIC_ADDRESS
ENV VITE_RENTHOME_PUBLIC_ADDRESS ${VITE_RENTHOME_PUBLIC_ADDRESS}

ARG VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID
ENV VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID ${VITE_RENTHOME_GOOGLE_OAUTH_CLIENT_ID}

ARG VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID
ENV VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID ${VITE_RENTHOME_FACEBOOK_OAUTH_CLIENT_ID}

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
COPY --from=builder /app/dist ./
COPY ./config/nginx.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT [ "nginx", "-g", "daemon off;" ]

