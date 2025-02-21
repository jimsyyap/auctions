-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each user
    username VARCHAR(50) UNIQUE NOT NULL,   -- Unique username
    email VARCHAR(100) UNIQUE NOT NULL,     -- Unique email address
    password_hash VARCHAR(255) NOT NULL,    -- Hashed password
    role VARCHAR(20) NOT NULL,              -- User role (Admin, Seller, Buyer)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the user was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Timestamp when the user was last updated
);

-- Create products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each product
    seller_id INT NOT NULL,                 -- Foreign key to the user who listed the product
    name VARCHAR(255) NOT NULL,             -- Name of the product
    description TEXT,                       -- Description of the product
    starting_price DECIMAL(10, 2) NOT NULL, -- Starting price for the auction
    reserve_price DECIMAL(10, 2),           -- Reserve price (minimum bid amount)
    auction_type VARCHAR(50) NOT NULL,      -- Type of auction (English, Dutch, Sealed-Bid)
    status VARCHAR(50) DEFAULT 'Active',    -- Status of the auction (Active, Ended, Sold)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the product was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the product was last updated
    CONSTRAINT fk_seller
        FOREIGN KEY (seller_id)
        REFERENCES users(id)
        ON DELETE CASCADE                   -- Delete products if the seller is deleted
);

-- Create auctions table
CREATE TABLE auctions (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each auction
    product_id INT NOT NULL,                -- Foreign key to the product being auctioned
    start_time TIMESTAMP NOT NULL,          -- Start time of the auction
    end_time TIMESTAMP NOT NULL,            -- End time of the auction
    current_price DECIMAL(10, 2) NOT NULL,  -- Current highest bid price
    winner_id INT,                          -- Foreign key to the winning user (nullable)
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE,                  -- Delete auctions if the product is deleted
    CONSTRAINT fk_winner
        FOREIGN KEY (winner_id)
        REFERENCES users(id)
        ON DELETE SET NULL                 -- Set winner_id to NULL if the user is deleted
);

-- Create bids table
CREATE TABLE bids (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each bid
    auction_id INT NOT NULL,                -- Foreign key to the auction
    bidder_id INT NOT NULL,                 -- Foreign key to the user who placed the bid
    amount DECIMAL(10, 2) NOT NULL,         -- Amount of the bid
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the bid was placed
    CONSTRAINT fk_auction
        FOREIGN KEY (auction_id)
        REFERENCES auctions(id)
        ON DELETE CASCADE,                  -- Delete bids if the auction is deleted
    CONSTRAINT fk_bidder
        FOREIGN KEY (bidder_id)
        REFERENCES users(id)
        ON DELETE CASCADE                   -- Delete bids if the user is deleted
);

-- Create payments table
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each payment
    auction_id INT NOT NULL,                -- Foreign key to the auction
    buyer_id INT NOT NULL,                  -- Foreign key to the user who made the payment
    amount DECIMAL(10, 2) NOT NULL,         -- Amount paid
    payment_status VARCHAR(50) NOT NULL,    -- Payment status (Pending, Completed, Refunded)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the payment was made
    CONSTRAINT fk_auction
        FOREIGN KEY (auction_id)
        REFERENCES auctions(id)
        ON DELETE CASCADE,                  -- Delete payments if the auction is deleted
    CONSTRAINT fk_buyer
        FOREIGN KEY (buyer_id)
        REFERENCES users(id)
        ON DELETE CASCADE                   -- Delete payments if the user is deleted
);

-- Create notifications table
CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each notification
    user_id INT NOT NULL,                   -- Foreign key to the user who receives the notification
    message TEXT NOT NULL,                  -- Notification message
    status VARCHAR(50) DEFAULT 'Unread',   -- Notification status (Read, Unread)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the notification was created
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE                   -- Delete notifications if the user is deleted
);
