# Write your MySQL query statement below
select firstName,lastName,city,state from Person Left Join Address On Person.personId = Address.personId