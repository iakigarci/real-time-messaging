# Use a specific version of Node.js Alpine for better stability
FROM node:18.19-alpine3.19

# Add necessary system packages
RUN apk add --no-cache \
    dumb-init \
    ca-certificates

# Create app directory
WORKDIR /app

# Copy package files first for better caching
COPY package*.json ./

# Install dependencies
RUN npm install --production \
    && npm cache clean --force

# Copy application files
COPY . .

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose gRPC port
EXPOSE 50051

# Use dumb-init as entrypoint to handle signals properly
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Start the service
CMD ["node", "server.js"] 