# Auction Website Database Schema

## Users and Authentication

### `users` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier for each user |
| email | VARCHAR(255) | UNIQUE, NOT NULL | User's email address |
| password_hash | VARCHAR(255) | NOT NULL | Hashed password |
| username | VARCHAR(50) | UNIQUE, NOT NULL | Display name for the user |
| first_name | VARCHAR(100) | | User's first name |
| last_name | VARCHAR(100) | | User's last name |
| profile_image_url | VARCHAR(255) | | URL to profile picture |
| phone_number | VARCHAR(20) | | Contact phone number |
| is_email_verified | BOOLEAN | DEFAULT FALSE | Whether email has been verified |
| is_phone_verified | BOOLEAN | DEFAULT FALSE | Whether phone has been verified |
| is_seller_verified | BOOLEAN | DEFAULT FALSE | Whether user is verified as a seller |
| seller_rating | DECIMAL(3,2) | | Average rating as a seller (1-5) |
| buyer_rating | DECIMAL(3,2) | | Average rating as a buyer (1-5) |
| account_status | ENUM | DEFAULT 'active' | Status: active, suspended, deactivated |
| created_at | TIMESTAMP | DEFAULT NOW() | Account creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Last update timestamp |

### `user_verification` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| user_id | UUID | FOREIGN KEY (users.id) | Reference to user |
| verification_type | ENUM | NOT NULL | Type: seller, identity, address |
| document_type | VARCHAR(50) | | Type of document submitted |
| document_url | VARCHAR(255) | | URL to verification document |
| status | ENUM | DEFAULT 'pending' | Status: pending, approved, rejected |
| reviewed_by | UUID | FOREIGN KEY (users.id) | Admin who reviewed |
| review_notes | TEXT | | Notes on verification review |
| submitted_at | TIMESTAMP | DEFAULT NOW() | Submission timestamp |
| reviewed_at | TIMESTAMP | | Review timestamp |

### `user_addresses` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| user_id | UUID | FOREIGN KEY (users.id) | Reference to user |
| address_type | ENUM | NOT NULL | Type: billing, shipping, both |
| is_default | BOOLEAN | DEFAULT FALSE | Whether this is default address |
| street_address1 | VARCHAR(255) | NOT NULL | Street address line 1 |
| street_address2 | VARCHAR(255) | | Street address line 2 |
| city | VARCHAR(100) | NOT NULL | City |
| state | VARCHAR(100) | | State/province |
| postal_code | VARCHAR(20) | NOT NULL | Postal/ZIP code |
| country | VARCHAR(100) | NOT NULL | Country |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

## Item Listings and Categories

### `categories` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| parent_id | UUID | FOREIGN KEY (categories.id) | Parent category (NULL for root) |
| name | VARCHAR(100) | NOT NULL | Category name |
| slug | VARCHAR(100) | UNIQUE, NOT NULL | URL-friendly name |
| description | TEXT | | Category description |
| icon_url | VARCHAR(255) | | URL to category icon |
| is_active | BOOLEAN | DEFAULT TRUE | Whether category is active |
| display_order | INTEGER | | Order for display purposes |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `category_attributes` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| category_id | UUID | FOREIGN KEY (categories.id) | Reference to category |
| name | VARCHAR(100) | NOT NULL | Attribute name |
| attribute_type | ENUM | NOT NULL | Type: text, number, date, boolean, select |
| is_required | BOOLEAN | DEFAULT FALSE | Whether required for listings |
| options | JSONB | | Options for select-type attributes |
| display_order | INTEGER | | Order for display purposes |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `listings` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| seller_id | UUID | FOREIGN KEY (users.id) | Reference to seller |
| category_id | UUID | FOREIGN KEY (categories.id) | Reference to category |
| title | VARCHAR(255) | NOT NULL | Listing title |
| description | TEXT | NOT NULL | Detailed description |
| condition | ENUM | NOT NULL | Condition: new, like_new, good, fair, poor |
| starting_price | DECIMAL(12,2) | NOT NULL | Starting bid amount |
| reserve_price | DECIMAL(12,2) | | Minimum acceptable price |
| buy_now_price | DECIMAL(12,2) | | Immediate purchase price |
| current_price | DECIMAL(12,2) | | Current highest bid |
| current_winner_id | UUID | FOREIGN KEY (users.id) | Current highest bidder |
| bid_count | INTEGER | DEFAULT 0 | Number of bids placed |
| auction_type | ENUM | DEFAULT 'timed' | Type: timed, reserve, buy_now |
| status | ENUM | NOT NULL | Status: draft, active, ended, sold, cancelled |
| start_time | TIMESTAMP | NOT NULL | Auction start time |
| end_time | TIMESTAMP | NOT NULL | Scheduled end time |
| actual_end_time | TIMESTAMP | | Actual end time (with extensions) |
| shipping_option | ENUM | NOT NULL | Options: seller_ships, local_pickup, both |
| shipping_price | DECIMAL(8,2) | | Cost of shipping |
| item_location | VARCHAR(255) | | Physical location of item |
| views_count | INTEGER | DEFAULT 0 | Number of listing views |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `listing_attributes` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing |
| attribute_id | UUID | FOREIGN KEY (category_attributes.id) | Reference to attribute |
| value | TEXT | NOT NULL | Attribute value |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `listing_images` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing |
| image_url | VARCHAR(255) | NOT NULL | URL to image |
| is_primary | BOOLEAN | DEFAULT FALSE | Whether this is main image |
| display_order | INTEGER | | Order for display purposes |
| alt_text | VARCHAR(255) | | Alternative text for image |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |

## Auction and Bidding 

### `bids` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing |
| bidder_id | UUID | FOREIGN KEY (users.id) | Reference to bidder |
| amount | DECIMAL(12,2) | NOT NULL | Bid amount |
| max_amount | DECIMAL(12,2) | | Maximum autobid amount |
| bid_time | TIMESTAMP | DEFAULT NOW() | Time bid was placed |
| status | ENUM | DEFAULT 'active' | Status: active, outbid, won, cancelled |
| is_auto_bid | BOOLEAN | DEFAULT FALSE | Whether placed by autobid system |
| ip_address | VARCHAR(45) | | IP address of bidder |
| user_agent | VARCHAR(255) | | Browser/device info |

### `watchlist` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| user_id | UUID | FOREIGN KEY (users.id) | Reference to user |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing |
| added_at | TIMESTAMP | DEFAULT NOW() | When item was added to watchlist |
| notifications_enabled | BOOLEAN | DEFAULT TRUE | Whether to send notifications |

## Payments and Transactions

### `transactions` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing |
| seller_id | UUID | FOREIGN KEY (users.id) | Reference to seller |
| buyer_id | UUID | FOREIGN KEY (users.id) | Reference to buyer |
| transaction_type | ENUM | NOT NULL | Type: auction_win, buy_now |
| amount | DECIMAL(12,2) | NOT NULL | Total amount |
| item_price | DECIMAL(12,2) | NOT NULL | Price of item |
| shipping_cost | DECIMAL(8,2) | DEFAULT 0 | Cost of shipping |
| platform_fee | DECIMAL(8,2) | DEFAULT 0 | Fees charged by platform |
| tax_amount | DECIMAL(8,2) | DEFAULT 0 | Tax collected |
| currency | CHAR(3) | DEFAULT 'USD' | Currency code |
| payment_status | ENUM | DEFAULT 'pending' | Status: pending, paid, refunded, disputed |
| fulfillment_status | ENUM | DEFAULT 'pending' | Status: pending, shipped, delivered, cancelled |
| tracking_number | VARCHAR(100) | | Shipping tracking number |
| tracking_url | VARCHAR(255) | | URL for tracking |
| buyer_address_id | UUID | FOREIGN KEY (user_addresses.id) | Shipping address |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `payments` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| transaction_id | UUID | FOREIGN KEY (transactions.id) | Reference to transaction |
| payment_method | ENUM | NOT NULL | Method: credit_card, paypal, etc. |
| payment_amount | DECIMAL(12,2) | NOT NULL | Amount paid |
| payment_fee | DECIMAL(8,2) | DEFAULT 0 | Payment processor fee |
| processor_reference | VARCHAR(255) | | Reference from payment processor |
| status | ENUM | NOT NULL | Status: pending, completed, failed, refunded |
| payment_date | TIMESTAMP | | When payment was made |
| refund_date | TIMESTAMP | | When refund was issued (if any) |
| metadata | JSONB | | Additional payment data |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

## Communication and Feedback

### `messages` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| listing_id | UUID | FOREIGN KEY (listings.id) | Reference to listing (optional) |
| sender_id | UUID | FOREIGN KEY (users.id) | Reference to sender |
| recipient_id | UUID | FOREIGN KEY (users.id) | Reference to recipient |
| subject | VARCHAR(255) | | Message subject |
| body | TEXT | NOT NULL | Message content |
| is_read | BOOLEAN | DEFAULT FALSE | Whether message has been read |
| read_at | TIMESTAMP | | When message was read |
| parent_id | UUID | FOREIGN KEY (messages.id) | Parent message in thread |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |

### `feedback` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| transaction_id | UUID | FOREIGN KEY (transactions.id) | Reference to transaction |
| reviewer_id | UUID | FOREIGN KEY (users.id) | User giving feedback |
| reviewee_id | UUID | FOREIGN KEY (users.id) | User receiving feedback |
| rating | INTEGER | NOT NULL | Rating (1-5) |
| comments | TEXT | | Feedback comments |
| as_seller | BOOLEAN | NOT NULL | Whether feedback about selling |
| response | TEXT | | Response to feedback |
| is_visible | BOOLEAN | DEFAULT TRUE | Whether visible to others |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

## Notifications

### `notification_preferences` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| user_id | UUID | FOREIGN KEY (users.id) | Reference to user |
| notification_type | ENUM | NOT NULL | Type: outbid, ending_soon, etc. |
| email_enabled | BOOLEAN | DEFAULT TRUE | Whether to send emails |
| push_enabled | BOOLEAN | DEFAULT TRUE | Whether to send push notifications |
| sms_enabled | BOOLEAN | DEFAULT FALSE | Whether to send SMS |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |
| updated_at | TIMESTAMP | DEFAULT NOW() | Update timestamp |

### `notifications` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| user_id | UUID | FOREIGN KEY (users.id) | Reference to user |
| type | ENUM | NOT NULL | Type: outbid, ending_soon, etc. |
| title | VARCHAR(255) | NOT NULL | Notification title |
| content | TEXT | NOT NULL | Notification content |
| related_id | UUID | | Reference to related entity |
| is_read | BOOLEAN | DEFAULT FALSE | Whether notification is read |
| read_at | TIMESTAMP | | When notification was read |
| created_at | TIMESTAMP | DEFAULT NOW() | Creation timestamp |

## System and Administration

### `system_settings` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| key | VARCHAR(100) | PRIMARY KEY | Setting identifier |
| value | TEXT | NOT NULL | Setting value |
| description | TEXT | | Description of setting |
| updated_at | TIMESTAMP | DEFAULT NOW() | Last update timestamp |
| updated_by | UUID | FOREIGN KEY (users.id) | Admin who updated setting |

### `admin_logs` Table
| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | UUID | PRIMARY KEY | Unique identifier |
| admin_id | UUID | FOREIGN KEY (users.id) | Reference to admin user |
| action | VARCHAR(255) | NOT NULL | Action performed |
| entity_type | VARCHAR(50) | | Type of entity affected |
| entity_id | UUID | | ID of entity affected |
| details | JSONB | | Additional action details |
| ip_address | VARCHAR(45) | | IP address of admin |
| created_at | TIMESTAMP | DEFAULT NOW() | Action timestamp |

## Indexes and Constraints

### Key Indexes
- `users` table:
  - Indexes on email, username for faster lookups
  - Index on email_verified, seller_verified for filtering

- `listings` table:
  - Indexes on seller_id, category_id for relationship queries
  - Composite index on status, end_time for active auction queries
  - Index on current_price for sorting by price

- `bids` table:
  - Indexes on listing_id, bidder_id for relationship queries
  - Composite index on listing_id, amount for determining highest bid

- `categories` table:
  - Index on parent_id for hierarchical queries
  - Index on slug for URL lookups

- `transactions` table:
  - Indexes on seller_id, buyer_id for user transaction history
  - Index on payment_status for filtering unpaid transactions

### Constraints
- Foreign key constraints for all relationships
- Check constraints for:
  - Valid price ranges (positive values)
  - Valid rating values (1-5)
  - Date constraints (end_time > start_time)
