
DROP TABLE IF EXISTS branch CASCADE;

ALTER TABLE car
  DROP branch_id UUID NOT NULL REFERENCES  branch(id);

ALTER TABLE car
  DROP  branch_percentage NUMERIC NOT NULL; 