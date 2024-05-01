CREATE TABLE
  users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    mobile_number VARCHAR(255) NOT NULL,
    password_digest VARCHAR(255) NOT NULL,
    tenant_id BIGINT NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id)
  );