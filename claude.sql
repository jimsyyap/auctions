-- Auction Website Database Schema

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create ENUMs
CREATE TYPE account_status_enum AS ENUM ('active', 'suspended', 'deactivated');
CREATE TYPE verification_type_enum AS ENUM ('seller', 'identity', 'address');
CREATE TYPE verification_status_enum AS ENUM ('pending', 'approved', 'rejected');
CREATE TYPE address_type_enum AS ENUM ('billing', 'shipping', 'both');
CREATE TYPE attribute_type_enum AS ENUM ('text', 'number', 'date', 'boolean', 'select');
CREATE TYPE item_condition_enum AS ENUM ('new', 'like_new', 'good', 'fair', 'poor');
CREATE TYPE auction_type_enum AS ENUM ('timed', 'reserve', 'buy_now');
CREATE TYPE listing_status_enum AS ENUM ('draft', 'active', 'ended', 'sold', 'cancelled');
CREATE TYPE shipping_option_enum AS ENUM ('seller_ships', 'local_pickup', 'both');
CREATE TYPE bid_status_enum AS ENUM ('active', 'outbid', 'won', 'cancelled');
CREATE TYPE transaction_type_enum AS ENUM ('auction_win', 'buy_now');
CREATE TYPE payment_status_enum AS ENUM ('pending', 'paid', 'refunded', 'disputed');
CREATE TYPE fulfillment_status_enum AS ENUM ('pending', 'shipped', 'delivered', 'cancelled');
CREATE TYPE payment_method_enum AS ENUM ('credit_card', 'paypal', 'bank_transfer', 'crypto');
CREATE TYPE payment_processor_status_enum AS ENUM ('pending', 'completed', 'failed', 'refunded');
CREATE TYPE notification_type_enum AS ENUM ('outbid', 'ending_soon', 'price_drop', 'message', 'feedback', 'payment', 'shipping');

-- Users and Authentication

-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    profile_image_url VARCHAR(255),
    phone_number VARCHAR(20),
    is_email_verified BOOLEAN DEFAULT FALSE,
    is_phone_verified BOOLEAN DEFAULT FALSE,
    is_seller_verified BOOLEAN DEFAULT FALSE,
    seller_rating DECIMAL(3,2),
    buyer_rating DECIMAL(3,2),
    account_status account_status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- User Verification Table
CREATE TABLE user_verification (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    verification_type verification_type_enum NOT NULL,
    document_type VARCHAR(50),
    document_url VARCHAR(255),
    status verification_status_enum DEFAULT 'pending',
    reviewed_by UUID REFERENCES users(id),
    review_notes TEXT,
    submitted_at TIMESTAMP DEFAULT NOW(),
    reviewed_at TIMESTAMP,
    CONSTRAINT valid_review CHECK ((status = 'pending' AND reviewed_by IS NULL AND reviewed_at IS NULL) OR 
                                  ((status = 'approved' OR status = 'rejected') AND reviewed_by IS NOT NULL AND reviewed_at IS NOT NULL))
);

-- User Addresses Table
CREATE TABLE user_addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    address_type address_type_enum NOT NULL,
    is_default BOOLEAN DEFAULT FALSE,
    street_address1 VARCHAR(255) NOT NULL,
    street_address2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Item Listings and Categories

-- Categories Table
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_id UUID REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon_url VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    display_order INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Category Attributes Table
CREATE TABLE category_attributes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    attribute_type attribute_type_enum NOT NULL,
    is_required BOOLEAN DEFAULT FALSE,
    options JSONB,
    display_order INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Listings Table
CREATE TABLE listings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID NOT NULL REFERENCES users(id),
    category_id UUID NOT NULL REFERENCES categories(id),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    condition item_condition_enum NOT NULL,
    starting_price DECIMAL(12,2) NOT NULL CHECK (starting_price >= 0),
    reserve_price DECIMAL(12,2) CHECK (reserve_price IS NULL OR reserve_price >= starting_price),
    buy_now_price DECIMAL(12,2) CHECK (buy_now_price IS NULL OR buy_now_price > starting_price),
    current_price DECIMAL(12,2) CHECK (current_price IS NULL OR current_price >= starting_price),
    current_winner_id UUID REFERENCES users(id),
    bid_count INTEGER DEFAULT 0,
    auction_type auction_type_enum DEFAULT 'timed',
    status listing_status_enum NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    actual_end_time TIMESTAMP,
    shipping_option shipping_option_enum NOT NULL,
    shipping_price DECIMAL(8,2) CHECK (shipping_price IS NULL OR shipping_price >= 0),
    item_location VARCHAR(255),
    views_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT valid_dates CHECK (end_time > start_time),
    CONSTRAINT valid_actual_end_time CHECK (actual_end_time IS NULL OR actual_end_time >= start_time)
);

-- Listing Attributes Table
CREATE TABLE listing_attributes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id UUID NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    attribute_id UUID NOT NULL REFERENCES category_attributes(id),
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Listing Images Table
CREATE TABLE listing_images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id UUID NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    image_url VARCHAR(255) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    display_order INTEGER,
    alt_text VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Auction and Bidding

-- Bids Table
CREATE TABLE bids (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id UUID NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    bidder_id UUID NOT NULL REFERENCES users(id),
    amount DECIMAL(12,2) NOT NULL CHECK (amount > 0),
    max_amount DECIMAL(12,2) CHECK (max_amount IS NULL OR max_amount >= amount),
    bid_time TIMESTAMP DEFAULT NOW(),
    status bid_status_enum DEFAULT 'active',
    is_auto_bid BOOLEAN DEFAULT FALSE,
    ip_address VARCHAR(45),
    user_agent VARCHAR(255)
);

-- Watchlist Table
CREATE TABLE watchlist (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    listing_id UUID NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    added_at TIMESTAMP DEFAULT NOW(),
    notifications_enabled BOOLEAN DEFAULT TRUE,
    UNIQUE(user_id, listing_id)
);

-- Payments and Transactions

-- Transactions Table
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id UUID NOT NULL REFERENCES listings(id),
    seller_id UUID NOT NULL REFERENCES users(id),
    buyer_id UUID NOT NULL REFERENCES users(id),
    transaction_type transaction_type_enum NOT NULL,
    amount DECIMAL(12,2) NOT NULL CHECK (amount > 0),
    item_price DECIMAL(12,2) NOT NULL CHECK (item_price > 0),
    shipping_cost DECIMAL(8,2) DEFAULT 0 CHECK (shipping_cost >= 0),
    platform_fee DECIMAL(8,2) DEFAULT 0 CHECK (platform_fee >= 0),
    tax_amount DECIMAL(8,2) DEFAULT 0 CHECK (tax_amount >= 0),
    currency CHAR(3) DEFAULT 'USD',
    payment_status payment_status_enum DEFAULT 'pending',
    fulfillment_status fulfillment_status_enum DEFAULT 'pending',
    tracking_number VARCHAR(100),
    tracking_url VARCHAR(255),
    buyer_address_id UUID REFERENCES user_addresses(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT amount_equals_sum CHECK (
        amount = item_price + shipping_cost + platform_fee + tax_amount
    )
);

-- Payments Table
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    payment_method payment_method_enum NOT NULL,
    payment_amount DECIMAL(12,2) NOT NULL CHECK (payment_amount > 0),
    payment_fee DECIMAL(8,2) DEFAULT 0 CHECK (payment_fee >= 0),
    processor_reference VARCHAR(255),
    status payment_processor_status_enum NOT NULL,
    payment_date TIMESTAMP,
    refund_date TIMESTAMP,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT valid_refund_date CHECK (refund_date IS NULL OR refund_date > payment_date)
);

-- Communication and Feedback

-- Messages Table
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    listing_id UUID REFERENCES listings(id) ON DELETE SET NULL,
    sender_id UUID NOT NULL REFERENCES users(id),
    recipient_id UUID NOT NULL REFERENCES users(id),
    subject VARCHAR(255),
    body TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    parent_id UUID REFERENCES messages(id),
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT different_users CHECK (sender_id != recipient_id)
);

-- Feedback Table
CREATE TABLE feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    reviewer_id UUID NOT NULL REFERENCES users(id),
    reviewee_id UUID NOT NULL REFERENCES users(id),
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comments TEXT,
    as_seller BOOLEAN NOT NULL,
    response TEXT,
    is_visible BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT different_users CHECK (reviewer_id != reviewee_id),
    UNIQUE(transaction_id, reviewer_id, as_seller)
);

-- Notifications

-- Notification Preferences Table
CREATE TABLE notification_preferences (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    notification_type notification_type_enum NOT NULL,
    email_enabled BOOLEAN DEFAULT TRUE,
    push_enabled BOOLEAN DEFAULT TRUE,
    sms_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, notification_type)
);

-- Notifications Table
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type notification_type_enum NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    related_id UUID,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- System and Administration

-- System Settings Table
CREATE TABLE system_settings (
    key VARCHAR(100) PRIMARY KEY,
    value TEXT NOT NULL,
    description TEXT,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by UUID REFERENCES users(id)
);

-- Admin Logs Table
CREATE TABLE admin_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    admin_id UUID NOT NULL REFERENCES users(id),
    action VARCHAR(255) NOT NULL,
    entity_type VARCHAR(50),
    entity_id UUID,
    details JSONB,
    ip_address VARCHAR(45),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create Indexes

-- Users Indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_verification_status ON users(is_email_verified, is_seller_verified);

-- Listings Indexes
CREATE INDEX idx_listings_seller ON listings(seller_id);
CREATE INDEX idx_listings_category ON listings(category_id);
CREATE INDEX idx_listings_status_end_time ON listings(status, end_time);
CREATE INDEX idx_listings_price ON listings(current_price);
CREATE INDEX idx_listings_status_date ON listings(status, start_time, end_time);

-- Bids Indexes
CREATE INDEX idx_bids_listing ON bids(listing_id);
CREATE INDEX idx_bids_bidder ON bids(bidder_id);
CREATE INDEX idx_bids_listing_amount ON bids(listing_id, amount);
CREATE INDEX idx_bids_status ON bids(status);

-- Categories Indexes
CREATE INDEX idx_categories_parent ON categories(parent_id);
CREATE INDEX idx_categories_slug ON categories(slug);

-- Transactions Indexes
CREATE INDEX idx_transactions_seller ON transactions(seller_id);
CREATE INDEX idx_transactions_buyer ON transactions(buyer_id);
CREATE INDEX idx_transactions_listing ON transactions(listing_id);
CREATE INDEX idx_transactions_payment_status ON transactions(payment_status);
CREATE INDEX idx_transactions_fulfillment_status ON transactions(fulfillment_status);

-- Messages Indexes
CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_recipient ON messages(recipient_id);
CREATE INDEX idx_messages_listing ON messages(listing_id);
CREATE INDEX idx_messages_thread ON messages(parent_id);
CREATE INDEX idx_messages_read_status ON messages(recipient_id, is_read);

-- Notifications Indexes
CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_user_read ON notifications(user_id, is_read);
CREATE INDEX idx_notifications_type ON notifications(type);

-- Feedback Indexes
CREATE INDEX idx_feedback_transaction ON feedback(transaction_id);
CREATE INDEX idx_feedback_reviewer ON feedback(reviewer_id);
CREATE INDEX idx_feedback_reviewee ON feedback(reviewee_id);
CREATE INDEX idx_feedback_rating ON feedback(rating);

-- Create Triggers for Updated_at

-- Function to update timestamps
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Add updated_at triggers for all tables with that column
CREATE TRIGGER update_users_modtime
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_user_addresses_modtime
    BEFORE UPDATE ON user_addresses
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_categories_modtime
    BEFORE UPDATE ON categories
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_category_attributes_modtime
    BEFORE UPDATE ON category_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_listings_modtime
    BEFORE UPDATE ON listings
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_listing_attributes_modtime
    BEFORE UPDATE ON listing_attributes
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_transactions_modtime
    BEFORE UPDATE ON transactions
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_payments_modtime
    BEFORE UPDATE ON payments
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_feedback_modtime
    BEFORE UPDATE ON feedback
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_notification_preferences_modtime
    BEFORE UPDATE ON notification_preferences
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- Create view for active listings
CREATE VIEW active_listings AS
SELECT l.*, u.username as seller_username, c.name as category_name,
       (SELECT COUNT(*) FROM watchlist w WHERE w.listing_id = l.id) as watch_count
FROM listings l
JOIN users u ON l.seller_id = u.id
JOIN categories c ON l.category_id = c.id
WHERE l.status = 'active' AND l.start_time <= NOW() AND l.end_time > NOW();

-- Create view for user statistics
CREATE VIEW user_statistics AS
SELECT 
    u.id,
    u.username,
    COUNT(DISTINCT CASE WHEN l.seller_id = u.id THEN l.id END) as listings_count,
    COUNT(DISTINCT CASE WHEN t.seller_id = u.id THEN t.id END) as sales_count,
    COUNT(DISTINCT CASE WHEN t.buyer_id = u.id THEN t.id END) as purchases_count,
    AVG(CASE WHEN f.reviewee_id = u.id AND f.as_seller = TRUE THEN f.rating END) as avg_seller_rating,
    COUNT(CASE WHEN f.reviewee_id = u.id AND f.as_seller = TRUE THEN 1 END) as seller_reviews_count,
    AVG(CASE WHEN f.reviewee_id = u.id AND f.as_seller = FALSE THEN f.rating END) as avg_buyer_rating,
    COUNT(CASE WHEN f.reviewee_id = u.id AND f.as_seller = FALSE THEN 1 END) as buyer_reviews_count
FROM users u
LEFT JOIN listings l ON u.id = l.seller_id
LEFT JOIN transactions t ON u.id = t.seller_id OR u.id = t.buyer_id
LEFT JOIN feedback f ON u.id = f.reviewee_id
GROUP BY u.id, u.username;
