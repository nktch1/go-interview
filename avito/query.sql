SELECT ReportsTo, COUNT(*) as Members, ROUND(AVG(Age)) AS 'Average Age'
FROM (
    SELECT ReportsTo, Age
    FROM maintable_PYS02
    GROUP BY FirstName, LastName
) t
WHERE ReportsTo != 'null'
GROUP BY ReportsTo