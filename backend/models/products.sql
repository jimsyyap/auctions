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
