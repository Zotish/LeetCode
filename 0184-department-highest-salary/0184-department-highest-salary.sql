select d.name as Department, e.name as Employee, e.salary as Salary 
from Employee e join Department d on e.departmentId = d.id
where 1 > (select count(salary) 
            from Employee e1 
            where e1.salary > e.salary and e1.departmentId = e.departmentId);