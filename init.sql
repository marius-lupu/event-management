DROP DATABASE IF EXISTS event_db;
CREATE DATABASE event_db;  
\c event_db;

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    ts TIMESTAMP NOT NULL,
    customer_id VARCHAR(14) NOT NULL,
    system_name VARCHAR(100),
    billed_amount INTEGER,
    msg VARCHAR(100) NOT NULL
);