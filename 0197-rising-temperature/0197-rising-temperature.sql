select A.id as Id from Weather  A join Weather  B on  DATEDIFF(day, A.recordDate, B.recordDate) = -1 and A.temperature>B.temperature
