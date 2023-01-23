
-- select 
-- sum(o.day_count * c.price)
-- FROM "Order" as o
-- join car as c on c.id = o.car_id
-- where c.id = 'e7a00c90-63bb-4104-9130-2e432390faa5'

-- select 
--     c.price * o.day_count  as total_price,
--     c.price * o.day_count/100*70 as Investor
-- FROM car as c
-- join Client as cl on cl.id = '9b6ee155-5d0d-4bc1-afe1-910777e50053'
-- join "Order" as o on o.car_id = '105636c0-74a9-411f-98d1-ab3f3de72fc1'

--       select 
--                 i.name,
-- 				c.price * o.day_count,
-- 				c.price * o.day_count/100*70,
--                 cl.first_name 
--                 o.total_price - paid_price 
-- 			FROM "Order" as o
-- 			join Client as cl on cl.id = o.client_id
-- 			join car as c on o.car_id = c.id
--             join Investor as i on investor_id = c.investor_id
-- 			where c.id = '105636c0-74a9-411f-98d1-ab3f3de72fc1' 
--             and cl.id = '9b6ee155-5d0d-4bc1-afe1-910777e50053'
     
-- ALTER TABLE car  
-- DROP COLUMN description;

-- ALTER TABLE car 
--   ADD ;
-- UPDATE car 

-- SET    investor_percentage = InvestorBenefit.price
-- FROM   InvestorBenefit 
-- WHERE  car.id = 'fd448c28-0fcb-4bbd-8f5a-9364c4c366c2' 
-- AND   investor_percentage IS DISTINCT FROM InvestorBenefit.price