-- vehicle type enum: 1 - CAR; 2 - MOTORCYCLE
-- status type enum: 0 - OFF; 1 - AVAILABLE
CREATE TABLE drivers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    vehicle_type INT,
    vehicle_plate VARCHAR(20) NOT NULL UNIQUE,
    status INT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE riders (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- enum: 1 - 'WAITING_FOR_DRIVER'; 2 - 'CANCELLED'; 3 - 'WAITING_FOR_PICKUP'; 4 - 'IN_PROGRESS'; 5 - 'FINISHED'
CREATE TABLE rides (
    id BIGSERIAL PRIMARY KEY,
    rider_id BIGINT NOT NULL,
    driver_id BIGINT,
    pickup_location POINT NOT NULL,
    destination POINT NOT NULL,
    status INT NOT NULL,
    distance DECIMAL(6, 2),
    fare DECIMAL(10, 2),
    final_price DECIMAL(10, 2),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ride_commissions (
    id BIGSERIAL PRIMARY KEY,
    ride_id BIGINT NOT NULL,
    platform_fee DECIMAL(10, 2),
    driver_commission DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO drivers (name, phone_number, vehicle_type, vehicle_plate) 
VALUES ('Agus', '08111901887', 1, 'B0110R'),('Bagas', '08111901888', 2, 'B0120R'), ('Fafa', '08111901889', 1, 'B0130R');

INSERT INTO riders (name, phone_number) 
VALUES ('Melati', '08111901999');
