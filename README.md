# AuctionHub - Online Auction Platform

A modern online auction platform inspired by TradeMe, built with Go, PostgreSQL, and React.

## 🚀 Features

- User accounts with seller verification
- Listing creation and management
- Real-time bidding system
- Category browsing and advanced search
- Watchlists and favorites
- Secure payment processing
- Messaging between buyers and sellers
- User ratings and feedback

## 🛠️ Technology Stack

### Backend
- **Go (Golang)**: High-performance web service
- **Gin/Echo**: Web framework
- **PostgreSQL**: Primary database
- **GORM**: ORM for database operations
- **Redis**: Caching and real-time features
- **JWT**: Authentication
- **WebSockets**: Real-time notifications

### Frontend
- **React**: UI library
- **TypeScript**: Type-safe JavaScript
- **Redux/Context**: State management
- **React Router**: Navigation
- **Styled Components/Tailwind**: Styling
- **Socket.io**: WebSocket client

## 🏗️ Project Structure

The project follows a clean architecture approach:

- `backend/`: Go backend service
  - `cmd/`: Application entry points
  - `internal/`: Core application code
  - `pkg/`: Reusable packages
  - `migrations/`: Database migrations

- `frontend/`: React frontend application
  - `src/components/`: Reusable UI components
  - `src/pages/`: Page components
  - `src/services/`: API service functions
  - `src/store/`: State management

## 🚦 Getting Started

### Prerequisites
- Go 1.20+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (optional)

### Backend Setup

```bash
cd backend
cp .env.example .env  # Configure your environment variables
go mod download
go run cmd/api/main.go
```

### Frontend Setup

```bash
cd frontend
npm install
npm start
```

### Docker Setup

```bash
docker-compose up -d
```

## 📝 API Documentation

API documentation is available at `/api/docs` when running the server.

## 🧪 Testing

### Backend Tests

```bash
cd backend
go test ./...
```

### Frontend Tests

```bash
cd frontend
npm test
```

## 🚀 Deployment

Detailed deployment instructions can be found in [DEPLOYMENT.md](./DEPLOYMENT.md).

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
