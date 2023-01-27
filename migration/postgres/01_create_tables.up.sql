CREATE TABLE investor (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE car (
    id UUID PRIMARY KEY, 
    state_number VARCHAR NOT NULL,
    model  VARCHAR NOT NULL,
    status VARCHAR DEFAULT 'in_stock',
    price NUMERIC NOT NULL,
    daily_limit INT NOT NULL,
    over_limit INT NOT NULL,
    investor_percentage NUMERIC NOT NULL,
    investor_id UUID NOT NULL REFERENCES investor(id),
    km INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE client (
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE "order" (
    id UUID PRIMARY KEY,
    car_id UUID NOT NULL REFERENCES car(id),
    client_id UUID NOT NULL REFERENCES client(id),
    total_price NUMERIC NOT NULL,
    paid_price NUMERIC DEFAULT 0,
    day_count INT NOT NULL,
    give_km NUMERIC,
    receive_km NUMERIC,
    status VARCHAR NOT NULL DEFAULT 'new',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE debtors (
    price NUMERIC NOT NULL,
    car_id UUID NOT NULL REFERENCES car(id),
    client_id UUID NOT NULL REFERENCES client(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

