SELECT
    b.id,
    b.name,
    SUM(o.paid_price / 100 * c.branch_percentage)
FROM "order" AS o
JOIN car AS c ON c.id = o.car_id
JOIN branch AS b ON b.id = c.branch_id
GROUP BY b.id, b.name
;