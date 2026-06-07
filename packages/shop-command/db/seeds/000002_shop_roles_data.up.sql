INSERT INTO shop_roles (id, code, name) VALUES
    (1, 'OWNER', 'Store Owner'),
    (2, 'MANAGER', 'Store Manager'),
    (3, 'CASHIER', 'Cashier'),
    (4, 'WAREHOUSE', 'Warehouse Staff'),
    (5, 'MARKETING', 'Marketing Staff')
ON CONFLICT (id) DO UPDATE
SET
    code = EXCLUDED.code,
    name = EXCLUDED.name;
