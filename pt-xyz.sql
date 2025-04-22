CREATE DATABASE `pt-xyz`

CREATE TABLE admin (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    user_name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);

CREATE TABLE consumer (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    KTP VARCHAR(20) NOT NULL UNIQUE,
    user_name VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    legal_name VARCHAR(255),
    born_location VARCHAR(100),
    born_date DATE,
    photo_KTP TEXT,
    selfie_photo TEXT,
    salary DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);


CREATE TABLE loan_limit (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    consumer_id CHAR(36) NOT NULL UNIQUE,
    limit_loan DECIMAL(15, 2) NOT NULL,
    limit_used DECIMAL(15, 2) NOT NULL,
    tenor_amount INT NOT NULL,
    FOREIGN KEY (consumer_id) REFERENCES consumer(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);



CREATE TABLE transaction_table (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    consumer_id CHAR(36) NOT NULL,
    company_id INT NOT NULL,
    external_transaction_id VARCHAR(255),
    company_name VARCHAR(255),
    company_category VARCHAR(255),
    contact_number VARCHAR(20),
    admin_fee DECIMAL(15, 2),
    total_price DECIMAL(15, 2),
    FOREIGN KEY (consumer_id) REFERENCES consumer(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);

CREATE TABLE loan_installment (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    transaction_id CHAR(36) NOT NULL,
    consumer_id CHAR(36) NOT NULL,
    installment_amount DECIMAL(15, 2),
    tenor DATE,
    interest_rate DECIMAL(5, 2),
    FOREIGN KEY (transaction_id) REFERENCES transaction_table(id),
    FOREIGN KEY (consumer_id) REFERENCES consumer(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);


CREATE TABLE transaction_product (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    transaction_id CHAR(36) NOT NULL,
    company_id INT,
    product_company_id CHAR(36),
    otr DECIMAL(15, 2),
    asset_name VARCHAR(255),
    price DECIMAL(15, 2),
    FOREIGN KEY (transaction_id) REFERENCES transaction_table(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);


CREATE TABLE master_product_pt_xyz (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    company_name VARCHAR(255),
    company_category VARCHAR(255),
    otr DECIMAL(15, 2),
    admin_fee DECIMAL(15, 2),
    asset_name VARCHAR(255),
    price DECIMAL(15, 2),
    stock INT,
    contact_number VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);


INSERT INTO master_product_pt_xyz (company_name, company_category, otr, admin_fee, asset_name, price, stock, contact_number)
VALUES
('PT XYZ Motors', 'Internal', 250000000, 5000000, 'Car Model A', 250000000, 10, '08123456789'),
('PT XYZ Motors', 'Internal', 150000000, 3000000, 'Car Model B', 150000000, 15, '08123456789'),
('PT XYZ Electronics', 'Internal', 5000000, 500000, 'Laptop Model X', 5000000, 20, '08234567890'),
('PT XYZ Electronics', 'Internal', 8000000, 800000, 'Smartphone Model Y', 8000000, 30, '08234567890');



