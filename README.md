# Online Auction Website

An online auction platform built with **Golang** (backend), **PostgreSQL** (database), and **ReactJS** (frontend). This project supports real-time bidding, multiple auction types, and secure payment integration.

---

## Features

- **User Roles**:
  - Admin: Manages users, auctions, and platform settings.
  - Seller: Lists items for auction and manages listings.
  - Buyer: Bids on items and makes purchases.
- **Auction Types**:
  - English Auction: Incremental bidding with a timer.
  - Dutch Auction: Price decreases over time.
  - Sealed-Bid Auction: Hidden bids with the highest bid winning.
- **Real-Time Bidding**: Powered by WebSockets.
- **Payment Integration**: Secure payments via Stripe (supports AUD).
- **Notifications**: Email and in-app notifications for outbid alerts and auction updates.
- **Advanced Features**:
  - AI-based product recommendations.
  - Social sharing of auctions.
  - Escrow services for high-value items.

---

## Tech Stack

- **Backend**: Golang (Gin framework)
- **Database**: PostgreSQL
- **Frontend**: ReactJS (TailwindCSS for styling)
- **Real-Time Communication**: WebSockets (Gorilla WebSocket)
- **Authentication**: JWT (JSON Web Tokens)
- **Payment Gateway**: Stripe
- **Deployment**:
  - Backend: AWS/Google Cloud/DigitalOcean
  - Frontend: Vercel/Netlify
  - Database: AWS RDS/Supabase

---

## Project Structure

```
auction-website/
├── backend/               # Golang backend
│   ├── main.go            # Entry point for the backend
│   ├── handlers/          # Request handlers
│   ├── models/            # Database models
│   ├── routes/            # API routes
│   ├── utils/             # Utility functions (e.g., JWT, WebSockets)
│   └── config/            # Configuration files (e.g., database, environment variables)
├── frontend/              # ReactJS frontend
│   ├── src/               # React source code
│   │   ├── components/    # Reusable components
│   │   ├── pages/         # Page components
│   │   ├── App.js         # Main application component
│   │   ├── index.js       # Entry point for the frontend
│   │   └── styles/        # CSS/Tailwind styles
│   ├── public/            # Static assets
│   └── package.json       # Frontend dependencies
├── README.md              # Project documentation
└── .gitignore             # Files to ignore in Git
```

---

## Getting Started

### Prerequisites

- **Go** (v1.20 or higher): [Install Go](https://golang.org/doc/install)
- **PostgreSQL** (v14 or higher): [Install PostgreSQL](https://www.postgresql.org/download/)
- **Node.js** (v16 or higher): [Install Node.js](https://nodejs.org/)
- **Git**: [Install Git](https://git-scm.com/)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/auction-website.git
   cd auction-website
   ```

2. **Set Up the Backend**:
   - Navigate to the `backend` directory:
     ```bash
     cd backend
     ```
   - Install Go dependencies:
     ```bash
     go mod download
     ```
   - Set up environment variables:
     ```bash
     cp .env.example .env
     ```
     Update the `.env` file with your PostgreSQL credentials:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your-username
     DB_PASSWORD=your-password
     DB_NAME=auction_db
     JWT_SECRET=your-jwt-secret
     ```

3. **Set Up the Database**:
   - Create the database:
     ```bash
     createdb auction_db
     ```
   - Run the SQL scripts to create tables (see `backend/models/schema.sql`).

4. **Run the Backend**:
   ```bash
   go run main.go
   ```
   The backend will start at `http://localhost:8080`.

5. **Set Up the Frontend**:
   - Navigate to the `frontend` directory:
     ```bash
     cd ../frontend
     ```
   - Install dependencies:
     ```bash
     npm install
     ```
   - Run the frontend:
     ```bash
     npm start
     ```
   The frontend will start at `http://localhost:3000`.

---

## API Endpoints

| Method | Endpoint            | Description                        |
|--------|---------------------|------------------------------------|
| GET    | `/`                 | Welcome message                   |
| POST   | `/api/register`     | User registration                 |
| POST   | `/api/login`        | User login                        |
| GET    | `/api/auctions`     | List all auctions                 |
| POST   | `/api/auctions`     | Create a new auction              |
| GET    | `/api/auctions/:id` | Get details of a specific auction |
| POST   | `/api/bids`         | Place a bid                       |

---

## Contributing

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message here"
   ```
4. Push to the branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Gin Framework](https://gin-gonic.com/)
- [ReactJS](https://reactjs.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [Stripe](https://stripe.com/)
