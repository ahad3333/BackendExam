
DROP TABLE IF EXISTS branch CASCADE;

ALTER TABLE car
  DROP branch_id ;

ALTER TABLE "order"
    DROP branch_id ;

  