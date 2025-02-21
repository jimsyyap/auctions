-- Insert a user
INSERT INTO users (username, email, password_hash, role)
VALUES ('john_doe', 'john@example.com', 'hashed_password_123', 'Seller');

-- Insert a product
INSERT INTO products (seller_id, name, description, starting_price, reserve_price, auction_type, status)
VALUES (1, 'Vintage Watch', 'A rare vintage watch from the 1920s.', 500.00, 1000.00, 'English', 'Active');

-- Insert an auction
INSERT INTO auctions (product_id, start_time, end_time, current_price)
VALUES (1, '2023-10-01 12:00:00', '2023-10-10 12:00:00', 500.00);

-- Insert a bid
INSERT INTO bids (auction_id, bidder_id, amount)
VALUES (1, 1, 550.00);

-- Insert a payment
INSERT INTO payments (auction_id, buyer_id, amount, payment_status)
VALUES (1, 1, 550.00, 'Completed');

-- Insert a notification
INSERT INTO notifications (user_id, message)
VALUES (1, 'Your bid of $550.00 has been placed.');
