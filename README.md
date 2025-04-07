# Real Time Messaging

A distributed real-time messaging system built with microservices architecture using WebSocket, NATS, and PostgreSQL.

## ⚡ Features

- Real-time message delivery using WebSocket
- Distributed architecture with microservices following DDD
- Message persistence in PostgreSQL
- NATS for message streaming
- Containerized deployment with Docker
- Health monitoring and automatic recovery
- Swagger for API documentation and testing
- Air for docker automtic build

## 🏗️ Architecture

The system consists of two main microservices:

1. **WebSocket Producer Service**
   - Handles incoming WebSocket connections
   - Processes and validates messages
   - Publishes messages to NATS

2. **WebSocket Consumer Service**
   - Subscribes to NATS topics
   - Processes and stores messages in PostgreSQL
   - Manages message delivery to connected clients

### Infrastructure Components

- **PostgreSQL**: Persistent message storage
- **NATS**: Message streaming and pub/sub
- **Docker**: Containerization and orchestration

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
- Producer Service: `http://localhost:${PRODUCER_HTTP_PORT}`
- Consumer Service: `http://localhost:${CONSUMER_HTTP_PORT}`
- PostgreSQL: `localhost:${POSTGRES_PORT}`
- NATS: `localhost:${NATS_PORT}`

## 📦 Service Details

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
# Build producer service
docker-compose build ws-producer-app

# Build consumer service
docker-compose build ws-consumer-app
```

### Viewing Logs

```bash
# View all logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f ws-producer-app
docker-compose logs -f ws-consumer-app
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
