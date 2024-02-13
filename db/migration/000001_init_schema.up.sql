CREATE TABLE custom_shorts (
    id SERIAL PRIMARY KEY,
    custom_short_url TEXT UNIQUE NOT NULL,
    actual_url TEXT NOT NULL,
    ttl INTERVAL
);
CREATE INDEX idx_custom_shorts_custom_short_url ON custom_shorts (custom_short_url);

CREATE TABLE rate_limits (
    ip_address TEXT UNIQUE NOT NULL,
    remaining_rate_limit INT NOT NULL,
    ttl INTERVAL
);
CREATE INDEX idx_rate_limits_ip_address ON rate_limits (ip_address);
