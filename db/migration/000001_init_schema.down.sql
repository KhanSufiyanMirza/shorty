-- Drop the index created in the up migration
DROP INDEX IF EXISTS idx_custom_shorts_custom_short_url;
DROP INDEX IF EXISTS idx_rate_limits_ip_address;

-- Drop the table
DROP TABLE custom_shorts;
DROP TABLE rate_limits;
