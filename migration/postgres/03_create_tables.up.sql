
CREATE TABLE branch (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
   

ALTER TABLE car
   ADD branch_id UUID NOT NULL REFERENCES  branch(id);


ALTER TABLE "order"
   ADD branch_id UUID NOT NULL REFERENCES  branch(id);

