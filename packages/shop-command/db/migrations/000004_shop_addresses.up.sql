CREATE TABLE shop_addresses (
    id UUID PRIMARY KEY,
    shop_id UUID NOT NULL REFERENCES shops(id) ON DELETE CASCADE,

    country_id TEXT NOT NULL,
    country_id TEXT NOT NULL,

    province_id TEXT NOT NULL,
    province_name TEXT NOT NULL,

    ward_id TEXT NOT NULL,
    ward_id TEXT NOT NULL,


    address_line TEXT NOT NULL,
    contact_name  TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    type TEXT NOT NULL,

    created_by UUID NOT NULL,
    updated_by UUID,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT shop_address_type_check
        CHECK (type IN ('PICKUP', 'RETURN'))
);

CREATE INDEX idx_shop_addresses_shop_id ON shop_addresses(shop_id);
CREATE INDEX idx_shop_addresses_country_id ON shop_addresses(country_id);
CREATE INDEX idx_shop_addresses_city_id ON shop_addresses(province_id);
CREATE INDEX idx_shop_addresses_ward_id ON shop_addresses(ward_id);
