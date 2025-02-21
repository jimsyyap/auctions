CREATE TABLE Auctions (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    current_price DECIMAL(10,2) NOT NULL,
    winner_id INT NULL,
    FOREIGN KEY (product_id) REFERENCES Products(id) ON DELETE CASCADE,
    FOREIGN KEY (winner_id) REFERENCES Users(id) ON DELETE SET NULL
);
