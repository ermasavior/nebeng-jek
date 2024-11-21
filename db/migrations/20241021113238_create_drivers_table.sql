-- migrate:up
CREATE TABLE drivers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    vehicle_type INT, -- vehicle type enum: 1 - CAR; 2 - MOTORCYCLE
    vehicle_plate VARCHAR(20) NOT NULL UNIQUE,
    status INT DEFAULT 0 NOT NULL, -- status type enum: 0 - OFF; 1 - AVAILABLE
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO drivers (name, phone_number, vehicle_type, vehicle_plate) 
VALUES ('Agus', '08111901887', 1, 'B0110R'),('Bagas', '08111901888', 2, 'B0120R'), ('Fafa', '08111901889', 1, 'B0130R');

-- migrate:down
DROP TABLE drivers;