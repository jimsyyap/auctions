CREATE TABLE Payments (
    id SERIAL PRIMARY KEY,
    auction_id INT NOT NULL,
    buyer_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_status VARCHAR(10) CHECK (payment_status IN ('Pending', 'Completed', 'Refunded')) NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (auction_id) REFERENCES Auctions(id) ON DELETE CASCADE,
    FOREIGN KEY (buyer_id) REFERENCES Users(id) ON DELETE CASCADE
);
