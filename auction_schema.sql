-- Create Users and Authentication Tables

CREATE TABLE users (
    id UUID PRIMARY KEY,
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
    account_status ENUM('active', 'suspended', 'deactivated') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_verification (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    verification_type ENUM('seller', 'identity', 'address') NOT NULL,
    document_type VARCHAR(50),
    document_url VARCHAR(255),
    status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
    reviewed_by UUID REFERENCES users(id),
    review_notes TEXT,
    submitted_at TIMESTAMP DEFAULT NOW(),
    reviewed_at TIMESTAMP
);

CREATE TABLE user_addresses (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    address_type ENUM('billing', 'shipping', 'both') NOT NULL,
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

-- Create Item Listings and Categories Tables

CREATE TABLE categories (
    id UUID PRIMARY KEY,
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

CREATE TABLE category_attributes (
    id UUID PRIMARY KEY,
    category_id UUID REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    attribute_type ENUM('text', 'number', 'date', 'boolean', 'select') NOT NULL,
    is_required BOOLEAN DEFAULT FALSE,
    options JSONB,
    display_order INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE listings (
    id UUID PRIMARY KEY,
    seller_id UUID REFERENCES users(id),
    category_id UUID REFERENCES categories(id),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    condition ENUM('new', 'like_new', 'good', 'fair', 'poor') NOT NULL,
    starting_price DECIMAL(12,2) NOT NULL,
    reserve_price DECIMAL(12,2),
    buy_now_price DECIMAL(12,2),
    current_price DECIMAL(12,2),
    current_winner_id UUID REFERENCES users(id),
    bid_count INTEGER DEFAULT 0,
    auction_type ENUM('timed', 'reserve', 'buy_now') DEFAULT 'timed',
    status ENUM('draft', 'active', 'ended', 'sold', 'cancelled') NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    actual_end_time TIMESTAMP,
    shipping_option ENUM('seller_ships', 'local_pickup', 'both') NOT NULL,
    shipping_price DECIMAL(8,2),
    item_location VARCHAR(255),
    views_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE listing_attributes (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id),
    attribute_id UUID REFERENCES category_attributes(id),
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE listing_images (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id),
    image_url VARCHAR(255) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    display_order INTEGER,
    alt_text VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create Auction and Bidding Tables

CREATE TABLE bids (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id),
    bidder_id UUID REFERENCES users(id),
    amount DECIMAL(12,2) NOT NULL,
    max_amount DECIMAL(12,2),
    bid_time TIMESTAMP DEFAULT NOW(),
    status ENUM('active', 'outbid', 'won', 'cancelled') DEFAULT 'active',
    is_auto_bid BOOLEAN DEFAULT FALSE,
    ip_address VARCHAR(45),
    user_agent VARCHAR(255)
);

CREATE TABLE watchlist (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    listing_id UUID REFERENCES listings(id),
    added_at TIMESTAMP DEFAULT NOW(),
    notifications_enabled BOOLEAN DEFAULT TRUE
);

-- Create Payments and Transactions Tables

CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id),
    seller_id UUID REFERENCES users(id),
    buyer_id UUID REFERENCES users(id),
    transaction_type ENUM('auction_win', 'buy_now') NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    item_price DECIMAL(12,2) NOT NULL,
    shipping_cost DECIMAL(8,2) DEFAULT 0,
    platform_fee DECIMAL(8,2) DEFAULT 0,
    tax_amount DECIMAL(8,2) DEFAULT 0,
    currency CHAR(3) DEFAULT 'USD',
    payment_status ENUM('pending', 'paid', 'refunded', 'disputed') DEFAULT 'pending',
    fulfillment_status ENUM('pending', 'shipped', 'delivered', 'cancelled') DEFAULT 'pending',
    tracking_number VARCHAR(100),
    tracking_url VARCHAR(255),
    buyer_address_id UUID REFERENCES user_addresses(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE payments (
    id UUID PRIMARY KEY,
    transaction_id UUID REFERENCES transactions(id),
    payment_method ENUM('credit_card', 'paypal') NOT NULL,
    payment_amount DECIMAL(12,2) NOT NULL,
    payment_fee DECIMAL(8,2) DEFAULT 0,
    processor_reference VARCHAR(255),
    status ENUM('pending', 'completed', 'failed', 'refunded') NOT NULL,
    payment_date TIMESTAMP,
    refund_date TIMESTAMP,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create Communication and Feedback Tables

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id),
    sender_id UUID REFERENCES users(id),
    recipient_id UUID REFERENCES users(id),
    subject VARCHAR(255),
    body TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    parent_id UUID REFERENCES messages(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE feedback (
    id UUID PRIMARY KEY,
    transaction_id UUID REFERENCES transactions(id),
    reviewer_id UUID REFERENCES users(id),
    reviewee_id UUID REFERENCES users(id),
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comments TEXT,
    as_seller BOOLEAN NOT NULL,
    response TEXT,
    is_visible BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create Notifications Tables

CREATE TABLE notification_preferences (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    notification_type ENUM('outbid', 'ending_soon') NOT NULL,
    email_enabled BOOLEAN DEFAULT TRUE,
    push_enabled BOOLEAN DEFAULT TRUE,
    sms_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    type ENUM('outbid', 'ending_soon') NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    related_id UUID,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create System and Administration Tables

CREATE TABLE system_settings (
    key VARCHAR(100) PRIMARY KEY,
    value TEXT NOT NULL,
    description TEXT,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by UUID REFERENCES users(id)
);

CREATE TABLE admin_logs (
    id UUID PRIMARY KEY,
    admin_id UUID REFERENCES users(id),
    action VARCHAR(255) NOT NULL,
    entity_type VARCHAR(50),
    entity_id UUID,
    details JSONB,
    ip_address VARCHAR(45),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Create Indexes

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email_verified ON users(is_email_verified);
CREATE INDEX idx_users_seller_verified ON users(is_seller_verified);

CREATE INDEX idx_listings_seller_id ON listings(seller_id);
CREATE INDEX idx_listings_category_id ON listings(category_id);
CREATE INDEX idx_listings_status_end_time ON listings(status, end_time);
CREATE INDEX idx_listings_current_price ON listings(current_price);

CREATE INDEX idx_bids_listing_id ON bids(listing_id);
CREATE INDEX idx_bids_bidder_id ON bids(bidder_id);
CREATE INDEX idx_bids_listing_id_amount ON bids(listing_id, amount);

CREATE INDEX idx_categories_parent_id ON categories(parent_id);
CREATE INDEX idx_categories_slug ON categories(slug);

CREATE INDEX idx_transactions_seller_id ON transactions(seller_id);
CREATE INDEX idx_transactions_buyer_id ON transactions(buyer_id);
CREATE INDEX idx_transactions_payment_status ON transactions(payment_status);

-- Add Constraints

ALTER TABLE listings ADD CONSTRAINT chk_start_end_time CHECK (end_time > start_time);
ALTER TABLE feedback ADD CONSTRAINT chk_rating_range CHECK (rating >= 1 AND rating <= 5);
ALTER TABLE bids ADD CONSTRAINT chk_positive_amount CHECK (amount > 0);
ALTER TABLE listings ADD CONSTRAINT chk_positive_price CHECK (starting_price > 0 AND (reserve_price IS NULL OR reserve_price > 0) AND (buy_now_price IS NULL OR buy_now_price > 0));
