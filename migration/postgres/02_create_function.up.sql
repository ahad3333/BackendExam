
CREATE OR REPLACE FUNCTION order_status_trigger() RETURNS trigger
LANGUAGE PLPGSQL
AS
$$
    BEGIN

        IF new.status = 'new'
            THEN
                UPDATE car SET status = 'booked' WHERE id = new.car_id;
        ELSIF 
            old.status = 'new' AND 
            new.status = 'client_took'
                THEN
                    UPDATE car SET status = 'in_use' WHERE id = new.car_id;
                    UPDATE "order" SET give_km = (SELECT km FROM car WHERE id = new.car_id) 
                    WHERE id = new.id;
        ELSIF 
            old.status = 'client_took' AND 
            new.status = 'client_returned'
                THEN
                    UPDATE car SET status = 'in_stock' WHERE id = new.car_id;
        END IF;

        return new;
    END;
$$;

CREATE TRIGGER order_trigger
AFTER INSERT OR UPDATE ON "order"
FOR EACH ROW EXECUTE PROCEDURE order_status_trigger();


CREATE OR REPLACE FUNCTION pere_limit() RETURNS trigger LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        car_info RECORD;
        limit_calc NUMERIC;
    BEGIN

        IF old.status = 'client_took' AND
            new.status = 'client_returned' THEN
    
            SELECT * FROM car
            INTO car_info
            WHERE id = new.car_id;

            limit_calc = (new.receive_km - new.give_km) - (new.day_count * car_info.daily_limit);

            IF limit_calc > 0 THEN

                INSERT INTO debtors (price, car_id, client_id, updated_at) VALUES
                    (
                        limit_calc * car_info.over_limit,
                        new.car_id,
                        new.client_id,
                        now()
                    );

            END IF;

        END IF;

        return new;
    END;
$$;


CREATE TRIGGER order_trigger_pere_limit
AFTER UPDATE ON "order"
FOR EACH ROW EXECUTE PROCEDURE pere_limit();
