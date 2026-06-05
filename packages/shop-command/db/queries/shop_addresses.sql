-- name: CreateShopAddress :exec
INSERT INTO shop_addresses (
    id,
    shop_id, 

    country_id,
    country_name,

    province_id,
    province_name,

    ward_id,
    ward_name,

    address_line,
    contact_name,
    phone_number,
    type,

    created_at,
    updated_at,

    created_by,
    updated_by
)
VALUES (
    @id::uuid,
    @shop_id::uuid,

    @country_id::text,
    @country_name::text,

    @province_id::text,
    @province_name::text,

    @ward_id::text,
    @ward_name::text,

    address_line::text,
    contact_name::text,
    phone_number::text,
    type::text,

    created_at::TIMESTAMPTZ,
    updated_at::TIMESTAMPTZ,

    created_by::uuid,
    updated_by::uuid
);


-- name: CheckRequiredAddresses :one
SELECT 
    EXISTS (
        SELECT 1 
        FROM shop_addresses sa_pickup 
        WHERE sa_pickup.shop_id = $1 AND sa_pickup.type = 'PICKUP'
    ) AS has_pickup,
    EXISTS (
        SELECT 1 
        FROM shop_addresses sa_return 
        WHERE sa_return.shop_id = $1 AND sa_return.type = 'RETURN'
    ) AS has_return;

