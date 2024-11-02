-- migrate:up
CREATE TABLE ride_commissions (
    id BIGSERIAL PRIMARY KEY,
    ride_id BIGINT NOT NULL,
    commission DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE ride_commissions;