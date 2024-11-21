-- migrate:up
CREATE TABLE rides (
    id BIGSERIAL PRIMARY KEY,
    rider_id BIGINT NOT NULL,
    driver_id BIGINT,
    pickup_location POINT NOT NULL,
    destination POINT NOT NULL,
    status INT NOT NULL, -- enum: 1 - 'WAITING_FOR_DRIVER'; 2 - 'CANCELLED'; 3 - 'WAITING_FOR_PICKUP'; 4 - 'IN_PROGRESS'; 5 - 'FINISHED'
    distance DECIMAL(6, 2),
    fare DECIMAL(10, 2),
    final_price DECIMAL(10, 2),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
drop table rides;
