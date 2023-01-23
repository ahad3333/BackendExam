CREATE TABLE Investor (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE Car(
    id UUID PRIMARY KEY,
    state_number VARCHAR NOT NULL,
    model VARCHAR NOT NULL,
    status VARCHAR DEFAULT 'in_stock',
    price NUMERIC NOT NULL,
    daily_limit NUMERIC NOT NULL,
    over_limit NUMERIC NOT NULL,
    investor_percentage VARCHAR DEFAULT '%70',
    investor_id UUID NOT NULL REFERENCES investor(id),
    km NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE Client (
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE "Order" (
    id UUID PRIMARY KEY,
    car_id UUID NOT NULL REFERENCES car(id),
    client_id UUID NOT NULL REFERENCES client(id),
    total_price NUMERIC,
    day_count NUMERIC,
    give_km NUMERIC,
    paid_price NUMERIC,
    receive_km NUMERIC,
    status text DEFAULT 'new',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE InvestorBenefit (
    name VARCHAR NOT NULL,
    price NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE Bebtors (
    name VARCHAR NOT NULL,
    Bebt NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
