DROP INDEX IF EXISTS idx_shop_addresses_province_id;

ALTER TABLE shop_addresses
    ALTER COLUMN country_id TYPE INT USING NULLIF(country_id, '')::INT,
    ALTER COLUMN province_id TYPE INT USING NULLIF(province_id, '')::INT,
    ALTER COLUMN ward_id TYPE INT USING NULLIF(ward_id, '')::INT;

ALTER TABLE shop_addresses RENAME COLUMN province_id TO city_id;
ALTER TABLE shop_addresses ADD COLUMN IF NOT EXISTS district_id INT NOT NULL DEFAULT 0;

ALTER TABLE shop_addresses DROP COLUMN IF EXISTS country_name;
ALTER TABLE shop_addresses DROP COLUMN IF EXISTS province_name;
ALTER TABLE shop_addresses DROP COLUMN IF EXISTS ward_name;

CREATE INDEX IF NOT EXISTS idx_shop_addresses_city_id ON shop_addresses(city_id);
CREATE INDEX IF NOT EXISTS idx_shop_addresses_district_id ON shop_addresses(district_id);
