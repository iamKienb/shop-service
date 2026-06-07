DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'shop_addresses' AND column_name = 'city_id'
    ) THEN
        ALTER TABLE shop_addresses RENAME COLUMN city_id TO province_id;
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'shop_addresses' AND column_name = 'district_id'
    ) THEN
        ALTER TABLE shop_addresses DROP COLUMN district_id;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'shop_addresses' AND column_name = 'country_name'
    ) THEN
        ALTER TABLE shop_addresses ADD COLUMN country_name TEXT NOT NULL DEFAULT '';
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'shop_addresses' AND column_name = 'province_name'
    ) THEN
        ALTER TABLE shop_addresses ADD COLUMN province_name TEXT NOT NULL DEFAULT '';
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'shop_addresses' AND column_name = 'ward_name'
    ) THEN
        ALTER TABLE shop_addresses ADD COLUMN ward_name TEXT NOT NULL DEFAULT '';
    END IF;
END $$;

ALTER TABLE shop_addresses
    ALTER COLUMN country_id TYPE TEXT USING country_id::TEXT,
    ALTER COLUMN province_id TYPE TEXT USING province_id::TEXT,
    ALTER COLUMN ward_id TYPE TEXT USING ward_id::TEXT;

DROP INDEX IF EXISTS idx_shop_addresses_city_id;
DROP INDEX IF EXISTS idx_shop_addresses_district_id;
CREATE INDEX IF NOT EXISTS idx_shop_addresses_province_id ON shop_addresses(province_id);
