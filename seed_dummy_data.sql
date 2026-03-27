-- Dummy Data for 30 Days Simulation (5 Users, 5 Businesses)
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE transaction_items;
TRUNCATE TABLE transactions;
TRUNCATE TABLE expenses;
TRUNCATE TABLE products;
TRUNCATE TABLE businesses;
TRUNCATE TABLE users;
SET FOREIGN_KEY_CHECKS = 1;

-- 1. USERS (Password: "password")
-- Hash: $2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L
INSERT INTO users (id, created_at, updated_at, name, email, password, provider) VALUES 
(1, NOW(), NOW(), 'John Doe', 'john@example.com', '$2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L', 'manual'),
(2, NOW(), NOW(), 'Jane Smith', 'jane@example.com', '$2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L', 'manual'),
(3, NOW(), NOW(), 'Bob Wilson', 'bob@example.com', '$2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L', 'manual'),
(4, NOW(), NOW(), 'Alice Brown', 'alice@example.com', '$2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L', 'manual'),
(5, NOW(), NOW(), 'Charlie Davis', 'charlie@example.com', '$2a$10$6mI8L2F2X.oR8k6R.J8k.u8L.J8k.u8L.J8k.u8L.J8k.u8L', 'manual');

-- 2. BUSINESSES
INSERT INTO businesses (id, created_at, updated_at, user_id, name, description, category, address) VALUES 
(1, NOW(), NOW(), 1, 'Sync Coffee', 'Kedai kopi modern dengan biji kopi pilihan.', 'F&B', 'Jl. Sudirman No. 123, Jakarta'),
(2, NOW(), NOW(), 2, 'Clean & Fresh', 'Layanan laundry ekspres dan kiloan.', 'Service', 'Jl. Thamrin No. 45, Jakarta'),
(3, NOW(), NOW(), 3, 'Tech Hub', 'Toko aksesoris gadget dan servis elektronik.', 'Retail', 'Jl. Gatot Subroto No. 67, Jakarta'),
(4, NOW(), NOW(), 4, 'Sharp Cuts', 'Barbershop premium dengan gaya rambut terkini.', 'Grooming', 'Jl. Rasuna Said No. 89, Jakarta'),
(5, NOW(), NOW(), 5, 'Tasty Bites', 'Restoran penyedia masakan nusantara.', 'F&B', 'Jl. Kemang Raya No. 12, Jakarta');

-- 3. PRODUCTS
-- Business 1: Sync Coffee
INSERT INTO products (id, created_at, updated_at, business_id, name, price, stock, min_stock) VALUES 
(1, NOW(), NOW(), 1, 'Espresso', 25000, 100, 10),
(2, NOW(), NOW(), 1, 'Latte', 35000, 100, 10),
(3, NOW(), NOW(), 1, 'Cappuccino', 35000, 100, 10),
(4, NOW(), NOW(), 1, 'Croissant', 20000, 100, 10);
-- Business 2: Clean & Fresh
INSERT INTO products (id, created_at, updated_at, business_id, name, price, stock, min_stock) VALUES 
(5, NOW(), NOW(), 2, 'Cuci Kiloan', 10000, 1000, 50),
(6, NOW(), NOW(), 2, 'Cuci Satuan', 15000, 500, 20),
(7, NOW(), NOW(), 2, 'Dry Clean', 50000, 200, 10);
-- Business 3: Tech Hub
INSERT INTO products (id, created_at, updated_at, business_id, name, price, stock, min_stock) VALUES 
(8, NOW(), NOW(), 3, 'Charging Cable', 50000, 50, 5),
(9, NOW(), NOW(), 3, 'Powerbank', 250000, 20, 2),
(10, NOW(), NOW(), 3, 'Screen Protector', 35000, 100, 10);
-- Business 4: Sharp Cuts
INSERT INTO products (id, created_at, updated_at, business_id, name, price, stock, min_stock) VALUES 
(11, NOW(), NOW(), 4, 'Haircut Premium', 75000, 1000, 0),
(12, NOW(), NOW(), 4, 'Shaving', 30000, 1000, 0),
(13, NOW(), NOW(), 4, 'Hair Vitamin', 15000, 50, 5);
-- Business 5: Tasty Bites
INSERT INTO products (id, created_at, updated_at, business_id, name, price, stock, min_stock) VALUES 
(14, NOW(), NOW(), 5, 'Nasi Goreng Special', 35000, 100, 5),
(15, NOW(), NOW(), 5, 'Ayam Bakar', 40000, 100, 5),
(16, NOW(), NOW(), 5, 'Es Teh Manis', 8000, 200, 10);

-- 4. TRANSACTIONS & TRANSACTION ITEMS (Simulasi 30 hari: 2026-02-25 s/d 2026-03-27)
-- Sync Coffee (Business 1) - Beberapa sampel
INSERT INTO transactions (id, created_at, updated_at, business_id, total_amount, payment_method, status, transaction_date) VALUES 
(1, '2026-02-26 09:00:00', '2026-02-26 09:00:00', 1, 60000, 'Cash', 'completed', '2026-02-26 09:00:00'),
(2, '2026-03-05 14:00:00', '2026-03-05 14:00:00', 1, 90000, 'QRIS', 'completed', '2026-03-05 14:00:00'),
(3, '2026-03-15 11:00:00', '2026-03-15 11:00:00', 1, 105000, 'Cash', 'completed', '2026-03-15 11:00:00');
INSERT INTO transaction_items (transaction_id, product_id, quantity, price) VALUES 
(1, 1, 1, 25000), (1, 2, 1, 35000),
(2, 2, 2, 35000), (2, 4, 1, 20000),
(3, 3, 3, 35000);

-- Clean & Fresh (Business 2)
INSERT INTO transactions (id, created_at, updated_at, business_id, total_amount, payment_method, status, transaction_date) VALUES 
(4, '2026-03-01 10:00:00', '2026-03-01 10:00:00', 2, 50000, 'Cash', 'completed', '2026-03-01 10:00:00'),
(5, '2026-03-10 16:00:00', '2026-03-10 16:00:00', 2, 100000, 'QRIS', 'completed', '2026-03-10 16:00:00');
INSERT INTO transaction_items (transaction_id, product_id, quantity, price) VALUES 
(4, 5, 5, 10000),
(5, 7, 2, 50000);

-- Tech Hub (Business 3)
INSERT INTO transactions (id, created_at, updated_at, business_id, total_amount, payment_method, status, transaction_date) VALUES 
(6, '2026-03-12 13:00:00', '2026-03-12 13:00:00', 3, 300000, 'Transfer', 'completed', '2026-03-12 13:00:00'),
(7, '2026-03-20 15:00:00', '2026-03-20 15:00:00', 3, 85000, 'QRIS', 'completed', '2026-03-20 15:00:00');
INSERT INTO transaction_items (transaction_id, product_id, quantity, price) VALUES 
(6, 9, 1, 250000), (6, 8, 1, 50000),
(7, 10, 1, 35000), (7, 8, 1, 50000);

-- Sharp Cuts (Business 4)
INSERT INTO transactions (id, created_at, updated_at, business_id, total_amount, payment_method, status, transaction_date) VALUES 
(8, '2026-03-05 18:00:00', '2026-03-05 18:00:00', 4, 75000, 'Cash', 'completed', '2026-03-05 18:00:00'),
(9, '2026-03-25 10:00:00', '2026-03-25 10:00:00', 4, 120000, 'QRIS', 'completed', '2026-03-25 10:00:00');
INSERT INTO transaction_items (transaction_id, product_id, quantity, price) VALUES 
(8, 11, 1, 75000),
(9, 11, 1, 75000), (9, 12, 1, 30000), (9, 13, 1, 15000);

-- Tasty Bites (Business 5)
INSERT INTO transactions (id, created_at, updated_at, business_id, total_amount, payment_method, status, transaction_date) VALUES 
(10, '2026-03-02 12:00:00', '2026-03-02 12:00:00', 5, 43000, 'Cash', 'completed', '2026-03-02 12:00:00'),
(11, '2026-03-26 19:00:00', '2026-03-26 19:00:00', 5, 88000, 'QRIS', 'completed', '2026-03-26 19:00:00');
INSERT INTO transaction_items (transaction_id, product_id, quantity, price) VALUES 
(10, 14, 1, 35000), (10, 16, 1, 8000),
(11, 15, 2, 40000), (11, 16, 1, 8000);

-- 5. EXPENSES
INSERT INTO expenses (business_id, title, amount, category, created_at, updated_at) VALUES 
(1, 'Biji Kopi Arabica', 1500000, 'Supplies', '2026-03-01 08:00:00', '2026-03-01 08:00:00'),
(1, 'Susu Segar', 500000, 'Supplies', '2026-03-10 08:00:00', '2026-03-10 08:00:00'),
(2, 'Detergen & Softener', 300000, 'Supplies', '2026-03-05 08:00:00', '2026-03-05 08:00:00'),
(2, 'Listrik Bulanan', 800000, 'Utility', '2026-03-20 08:00:00', '2026-03-20 08:00:00'),
(3, 'Restock Kabel Data', 2000000, 'Purchase', '2026-03-02 08:00:00', '2026-03-02 08:00:00'),
(4, 'Sewa Kursi Barber', 1000000, 'Rent', '2026-02-28 08:00:00', '2026-02-28 08:00:00'),
(5, 'Bahan Sembako', 2500000, 'Supplies', '2026-03-01 08:00:00', '2026-03-01 08:00:00'),
(5, 'Gaji Tukang Masak', 3000000, 'Sallary', '2026-03-25 08:00:00', '2026-03-25 08:00:00');
