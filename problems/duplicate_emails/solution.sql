# Write your MySQL query statement below
select distinct e1.Email
from Person e1 join Person e2 
on e1.Email = e2.Email
where e1.id <> e2.id