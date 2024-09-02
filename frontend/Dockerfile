# Use a Node image to build the application
FROM node:20-alpine as build

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json postcss.config.js tailwind.config.js  ./
RUN npm install

# Copy the source code
COPY . .

# Build the application
RUN npm build

# Use a lightweight web server to serve the frontend
FROM nginx:alpine

# Copy the build files to the NGINX web server's root
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port
EXPOSE 80

# Start NGINX server
CMD ["nginx", "-g", "daemon off;"]
