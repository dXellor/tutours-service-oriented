FROM node:22.12.0 as build
WORKDIR /src

COPY package.json ./
COPY package-lock.json ./
RUN npm ci

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Copy the rest of the application files
COPY . .

# Build the Angular app for production
RUN ng build

# Stage 2: Set up Nginx to serve the Angular app
FROM nginx:latest

# Copy the Angular build output from the build stage to Nginx's html directory
COPY --from=build /src/dist/explorer /usr/share/nginx/html
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80 for serving the frontend
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
