CREATE DATABASE IF NOT EXISTS sensors;

USE sensors;

CREATE TABLE IF NOT EXISTS sensors (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    sensor_value DOUBLE NOT NULL,
    sensor_type VARCHAR(50) NOT NULL,
    id1 CHAR(1) NOT NULL,
    id2 INT NOT NULL,
    timestamp DATETIME NOT NULL,
    INDEX idx_id1_id2 (id1, id2),
    INDEX idx_timestamp (timestamp),
    INDEX idx_id1_id2_ts (id1, id2, timestamp)
);