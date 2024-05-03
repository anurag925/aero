CREATE TABLE
  users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    mobile_number VARCHAR(255),
    status TINYINT NOT NULL DEFAULT 0,
    verified BOOLEAN NOT NULL DEFAULT false,
    iter INTEGER,
    salt VARCHAR(255),
    password_digest VARCHAR(255),
    tenant_id BIGINT NOT NULL,
    UNIQUE(email),
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    INDEX(email)
  );