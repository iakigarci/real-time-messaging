# Real Time Messaging

A distributed real-time messaging system built with microservices architecture using WebSocket, NATS, and PostgreSQL.

## ⚡ Features

- Real-time message delivery using WebSocket and events
- Distributed architecture with microservices following DDD
- Message persistence in PostgreSQL
- NATS for message streaming
- Containerized deployment with Docker
- Health monitoring and automatic recovery
- Swagger for API documentation and testing
- Air for docker automatic build
- Authentication service with gRPC

## 🏗️ Architecture

The system consists of three main microservices:

1. **Authentication Service**
   - Handles user authentication and authorization
   - Provides gRPC endpoints for user management
   - Manages user sessions and tokens

2. **WebSocket Producer Service**
   - Handles incoming WebSocket connections
   - Processes and validates messages
   - Publishes messages to NATS

3. **WebSocket Consumer Service**
   - Subscribes to NATS topics
   - Processes and stores messages in PostgreSQL
   - Manages message delivery to connected clients

### Infrastructure Components

- **PostgreSQL**: Persistent message storage and user data
- **NATS**: Message streaming and pub/sub
- **Docker**: Containerization and orchestration

### Software Design

The services are written in Go, with the auth service following a flat design and the WebSocket services following Domain-Driven Design (DDD) principles.

#### DDD Service Structure (WebSocket Services)
```
service/
├── api/           # API definitions and interfaces
├── cmd/           # Application entry points
├── config/        # Configuration files and structures
├── docs/          # API documentation (Swagger)
├── internal/      # Private application code
│   ├── domain/    # Domain models and business logic
│     ├── ports/     # Ports (interfaces) for external communication
│     ├── services/  # Application services
│   ├── adapters/  # Adapters implementing ports
│   └── mocks/     # Mocks for testing

├── pkg/           # Public packages that can be imported by other services
└── tmp/           # Temporary files and logs
```

#### Flat Design Structure (Auth Service)
```
service/
├── src/           # Source code
└── test/          # Test files
```

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Make (optional, for convenience commands)

### Environment Setup

1. Copy the environment template:
   ```bash
   cp .env.example .env
   ```

2. Configure your environment variables in `.env`

### Running the Application

Start all services using Docker Compose:

```bash
docker-compose up -d
```

The services will be available at:
- Auth Service: `localhost:50051` (gRPC)
- Producer Service: `http://localhost:${PRODUCER_HTTP_PORT}`
- Consumer Service: `http://localhost:${CONSUMER_HTTP_PORT}`
- PostgreSQL: `localhost:${POSTGRES_PORT}`
- NATS: `localhost:${NATS_PORT}`

## 📦 Service Details

### Authentication Service
- Handles user authentication and authorization
- Provides gRPC endpoints for user management
- Manages user sessions and tokens
- Port: 50051 (gRPC)

### WebSocket Producer Service
- Handles WebSocket connections
- Validates and processes incoming messages
- Publishes messages to NATS topics
- Port: ${PRODUCER_HTTP_PORT}

### WebSocket Consumer Service
- Subscribes to NATS topics
- Processes and stores messages
- Manages WebSocket connections for message delivery
- Port: ${CONSUMER_HTTP_PORT}

## 🔧 Development

### Building Services

```bash
# Build all services
docker-compose build --no-cache

# Build specific service
docker-compose build auth-service
docker-compose build ws-producer-app
docker-compose build ws-consumer-app
```

### Viewing Logs

```bash
# View all logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f auth-service
docker-compose logs -f ws-producer-app
docker-compose logs -f ws-consumer-app
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
