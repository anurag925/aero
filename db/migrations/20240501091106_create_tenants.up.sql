CREATE TABLE
  tenants (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255),
    status TINYINT NOT NULL DEFAULT 0,
    schema_name VARCHAR(255),
    data_source VARCHAR(255),
    settings JSON
  );


