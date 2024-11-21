-- migrate:up
CREATE TABLE riders (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO riders (name, phone_number) 
VALUES ('Melati', '08111901999');

-- migrate:down
DROP TABLE riders;
