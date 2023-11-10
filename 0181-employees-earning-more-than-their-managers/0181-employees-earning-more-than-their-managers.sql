select B.name as Employee from Employee as A  join Employee as B on A.id=B.managerId
where A.salary < B.salary