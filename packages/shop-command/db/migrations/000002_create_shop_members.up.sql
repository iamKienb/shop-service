CREATE TABLE shop_members (
    shop_id UUID REFERENCES shops(id) ON DELETE CASCADE,
    member_id UUID NOT NULL,
    joined_at TIMESTAMPTZ DEFAULT now(),
    added_by  UUID NOT NULL,
    PRIMARY KEY (shop_id, member_id)
);
