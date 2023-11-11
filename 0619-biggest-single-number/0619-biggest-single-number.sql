SELECT MAX(num) AS num
FROM (
    SELECT num
    FROM MyNumbers
    GROUP BY num
    HAVING COUNT(*) = 1
) AS SingleNumbers
UNION ALL
SELECT NULL
WHERE NOT EXISTS (
    SELECT 1
    FROM MyNumbers
    GROUP BY num
    HAVING COUNT(*) != 1
);
