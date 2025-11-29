# E-commerce Application: Read/Write Pattern Matrix

## Overview
This matrix outlines the patterns, scenarios, lookup types, and read/write intensity for the e-commerce application entities.

---

## 1. Users Entity

| Pattern | Scenario | Read/Write | Lookup Type | Read-Heavy / Write-Heavy | RPS Estimate | Notes |
|---------|----------|------------|-------------|-------------------------|--------------|-------|
| **Authentication** | User login | Read | Email/Mobile (Index) | Read-Heavy | High (1000-5000 RPS) | Frequent lookups by email/mobile |
| **Registration** | New user signup | Write | Email/Mobile (Unique Index) | Write-Heavy | Medium (100-500 RPS) | Requires uniqueness check |
| **Profile View** | View user profile | Read | User ID (Primary Key) | Read-Heavy | Medium (500-2000 RPS) | Direct PK lookup |
| **Profile Update** | Update user details | Write | User ID (Primary Key) | Write-Heavy | Low (50-200 RPS) | Less frequent updates |
| **Address Management** | Add/Update address | Write | User ID (Primary Key) | Write-Heavy | Low (50-200 RPS) | Array field updates |
| **User Search** | Admin search users | Read | Name/Email (Full-text/Index) | Read-Heavy | Low (10-50 RPS) | Admin operations |

**Summary for Users**: **Read-Heavy** (Login and profile views dominate)

---

## 2. Products Entity

| Pattern | Scenario | Read/Write | Lookup Type | Read-Heavy / Write-Heavy | RPS Estimate | Notes |
|---------|----------|------------|-------------|-------------------------|--------------|-------|
| **Product Listing** | Browse products | Read | Category/Filter (Composite Index) | Read-Heavy | Very High (5000-20000 RPS) | Most frequent operation |
| **Product Search** | Search by name/description | Read | Full-text Search Index | Read-Heavy | High (2000-10000 RPS) | Search queries |
| **Product Details** | View product page | Read | Product ID (Primary Key) | Read-Heavy | Very High (5000-20000 RPS) | Direct PK lookup |
| **Price Filter** | Filter by price range | Read | Price (Range Index) | Read-Heavy | High (1000-5000 RPS) | Range queries |
| **Inventory Check** | Check product availability | Read | Product ID + Quantity (Index) | Read-Heavy | Very High (5000-20000 RPS) | Real-time inventory |
| **Product Creation** | Add new product | Write | Product ID (Primary Key) | Write-Heavy | Low (10-100 RPS) | Admin operation |
| **Inventory Update** | Update stock quantity | Write | Product ID (Primary Key) | Write-Heavy | Medium (500-2000 RPS) | Frequent during orders |
| **Price Update** | Update product price | Write | Product ID (Primary Key) | Write-Heavy | Low (10-100 RPS) | Admin operation |
| **Review Aggregation** | Calculate average rating | Read/Write | Product ID (Primary Key) | Read-Heavy | Medium (500-2000 RPS) | Read reviews, write avg |

**Summary for Products**: **Read-Heavy** (Browsing, searching, viewing dominate)

---

## 3. Cart Entity

| Pattern | Scenario | Read/Write | Lookup Type | Read-Heavy / Write-Heavy | RPS Estimate | Notes |
|---------|----------|------------|-------------|-------------------------|--------------|-------|
| **Add to Cart** | Add product to cart | Write | User ID (Foreign Key Index) | Write-Heavy | High (2000-10000 RPS) | Frequent operation |
| **View Cart** | Display cart items | Read | User ID (Foreign Key Index) | Read-Heavy | High (2000-10000 RPS) | Frequent operation |
| **Update Quantity** | Change item quantity | Write | Cart ID + Product ID (Composite Index) | Write-Heavy | Medium (500-2000 RPS) | Update operation |
| **Remove Item** | Remove from cart | Write | Cart ID + Product ID (Composite Index) | Write-Heavy | Medium (500-2000 RPS) | Delete operation |
| **Cart Total** | Calculate cart total | Read | User ID (Foreign Key Index) | Read-Heavy | High (2000-10000 RPS) | Aggregation query |
| **Clear Cart** | Empty cart after checkout | Write | User ID (Foreign Key Index) | Write-Heavy | Medium (500-2000 RPS) | Bulk delete |
| **Cart Abandonment** | Track abandoned carts | Read | User ID + Status (Composite Index) | Read-Heavy | Low (10-50 RPS) | Analytics query |

**Summary for Cart**: **Balanced** (Similar read/write frequency, both high)

---

## 4. Orders Entity

| Pattern | Scenario | Read/Write | Lookup Type | Read-Heavy / Write-Heavy | RPS Estimate | Notes |
|---------|----------|------------|-------------|-------------------------|--------------|-------|
| **Place Order** | Create new order | Write | Order ID (Primary Key) | Write-Heavy | Medium (500-2000 RPS) | Transaction creation |
| **Order History** | View user orders | Read | User ID (Foreign Key Index) | Read-Heavy | High (1000-5000 RPS) | List query with FK |
| **Order Details** | View specific order | Read | Order ID (Primary Key) | Read-Heavy | Medium (500-2000 RPS) | Direct PK lookup |
| **Order Status Update** | Update order status | Write | Order ID (Primary Key) | Write-Heavy | Medium (500-2000 RPS) | Status transitions |
| **Order Tracking** | Track order by ID | Read | Order ID (Primary Key) | Read-Heavy | High (1000-5000 RPS) | Frequent lookups |
| **Order Search** | Search orders (Admin) | Read | Order ID/User ID/Status (Composite Index) | Read-Heavy | Low (10-100 RPS) | Admin operations |
| **Order Analytics** | Sales reports | Read | Date Range + Status (Composite Index) | Read-Heavy | Low (1-10 RPS) | Aggregation queries |
| **Order Cancellation** | Cancel order | Write | Order ID (Primary Key) | Write-Heavy | Low (50-200 RPS) | Status update |

**Summary for Orders**: **Read-Heavy** (Order history and tracking dominate)

---

## RPS Estimation Methodology

### How RPS Estimates Were Determined

The RPS (Requests Per Second) estimates in this matrix are based on the following factors:

#### 1. **User Behavior Patterns**
- **Browsing vs Purchasing Ratio**: Industry standard shows ~2-5% conversion rate
  - For every 100 product views, ~2-5 purchases occur
  - This means browsing/reading operations are 20-50x more frequent than writes
- **Session Patterns**: Average user session includes:
  - Multiple product views (10-50 per session)
  - Several cart operations (add, view, update)
  - One order placement (if converted)

#### 2. **Operation Frequency Classification**
- **Very High (5,000-20,000 RPS)**: 
  - Core user-facing operations that happen on every page load
  - Examples: Product listing, product details, inventory checks
  - Assumes: 100K-500K daily active users with multiple page views
  
- **High (1,000-10,000 RPS)**:
  - Frequent operations but not on every page load
  - Examples: Login, cart operations, order tracking
  - Assumes: 50K-200K daily active users with regular interactions
  
- **Medium (500-2,000 RPS)**:
  - Regular operations but less frequent
  - Examples: Order placement, profile updates, inventory updates
  - Assumes: 20K-100K daily active users
  
- **Low (10-200 RPS)**:
  - Infrequent operations, admin tasks, or one-time actions
  - Examples: Registration, product creation, analytics
  - Assumes: Small percentage of users or admin-only operations

#### 3. **Industry Benchmarks**
Based on typical e-commerce patterns:
- **Small-Medium E-commerce**: 1,000-10,000 RPS peak
- **Medium-Large E-commerce**: 10,000-50,000 RPS peak
- **Large E-commerce (Amazon-scale)**: 100,000+ RPS peak

These estimates assume a **medium-scale e-commerce** application.

#### 4. **Time-Based Considerations**
- **Peak Hours**: RPS can be 3-5x higher during:
  - Sale events (Black Friday, etc.)
  - Evening hours (6 PM - 10 PM)
  - Weekends
- **Off-Peak Hours**: RPS can drop to 10-20% of peak
- Estimates provided are for **average peak hours**, not absolute maximums

#### 5. **Operation-Specific Reasoning**

**Product Operations (Very High RPS)**:
- Product listing/details: Every user browsing session generates multiple requests
- Inventory checks: Happens on every product view and before cart operations
- Rationale: These are the "hot path" operations in e-commerce

**Cart Operations (High RPS)**:
- Add to cart: Happens frequently but not as often as browsing
- View cart: Checked multiple times before checkout
- Rationale: High engagement but lower than browsing

**Order Operations (Medium RPS)**:
- Order placement: Only happens when user converts (2-5% of sessions)
- Order tracking: Users check 2-5 times per order
- Rationale: Lower frequency but still significant volume

**User Operations (Variable RPS)**:
- Login: High frequency (daily or per session)
- Registration: Low frequency (one-time per user)
- Profile updates: Infrequent (monthly or less)

#### 6. **Scaling Factors**

These estimates can be adjusted based on:
- **User Base Size**: Scale proportionally
  - 10K users → Divide by 10
  - 1M users → Multiply by 10
- **Business Model**: 
  - B2B: Lower RPS, higher transaction value
  - B2C: Higher RPS, lower transaction value
- **Geographic Distribution**: 
  - Single region: Concentrated peak hours
  - Global: Distributed load across time zones

#### 7. **Calculation Example**

For **Product Listing (5,000-20,000 RPS)**:
- Assumptions:
  - 100,000 daily active users
  - Average 20 product page views per user per day
  - Peak hours = 4 hours (16% of day) with 40% of traffic
  - Peak hour requests = 100,000 users × 20 views × 40% = 800,000 requests
  - Peak hour RPS = 800,000 / 3,600 seconds = ~222 RPS
  - With 5x peak multiplier (sale events) = ~1,100 RPS
  - Multiple concurrent users browsing = 5,000-20,000 RPS range

### Important Notes

⚠️ **These are estimates, not guarantees**
- Actual RPS depends on: user base size, business model, marketing campaigns, seasonality
- Monitor actual traffic patterns and adjust estimates accordingly
- Use these for capacity planning and index design discussions
- Consider implementing monitoring and auto-scaling based on real metrics

---

## Overall Application Summary

### Read vs Write Ratio
- **Read Operations**: ~70-80% of total operations
- **Write Operations**: ~20-30% of total operations

### System Classification
**Overall: READ-HEAVY**

### High RPS Scenarios (Top 5)
1. **Product Listing/Browsing** - 5,000-20,000 RPS (Read)
2. **Product Details View** - 5,000-20,000 RPS (Read)
3. **Inventory Check** - 5,000-20,000 RPS (Read)
4. **Add to Cart** - 2,000-10,000 RPS (Write)
5. **View Cart** - 2,000-10,000 RPS (Read)

### Critical Indexes Required

1. **Users**
   - Primary Key: `id`
   - Unique Index: `email`, `mobile`
   - Index: `email`, `mobile` (for login)

2. **Products**
   - Primary Key: `id` (or product identifier)
   - Full-text Index: `name`, `description`
   - Index: `price` (for range queries)
   - Index: `category` (if exists)
   - Index: `quantity` (for inventory checks)

3. **Cart**
   - Primary Key: `id`
   - Foreign Key Index: `user` (User ID)
   - Composite Index: `(user, status)` (for active carts)
   - Index on items array (if supported)

4. **Orders**
   - Primary Key: `orderID`
   - Foreign Key Index: `user` (User ID)
   - Composite Index: `(user, status)` (for order history)
   - Index: `status` (for filtering)
   - Index: `date` (for time-based queries)

### Database Design Recommendations

1. **Read Optimization**
   - Implement read replicas for product browsing
   - Use caching (Redis) for product listings and details
   - Cache cart data in memory
   - Use CDN for static product images

2. **Write Optimization**
   - Batch inventory updates
   - Use connection pooling
   - Implement write-behind caching for cart updates
   - Use optimistic locking for inventory updates

3. **Index Strategy**
   - Prioritize indexes on frequently queried fields
   - Monitor index usage and remove unused indexes
   - Consider partial indexes for filtered queries
   - Use composite indexes for multi-column queries
