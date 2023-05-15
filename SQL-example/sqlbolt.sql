-- Website: https://sqlbolt.com/

-- Exercise 1

-- SQL Lesson 1: SELECT queries 101

/* 
1.Find the title of each film 
2.Find the director of each film 
3.Find the title and director of each film
4.Find the title and year of each film
5.Find all the information about each film 
*/

SELECT * FROM movies;

SELECT director FROM movies;

SELECT title, director FROM movies; 

SELECT title, year FROM movies; 

SELECT * FROM movies; 

-- Exercise 2

-- SQL Lesson 2: Queries with constraints (Pt. 1)

/* 
1.Find the movie with a row id of 6
2.Find the movies released in the years between 2000 and 2010
3.Find the movies not released in the years between 2000 and 2010
4.Find the first 5 Pixar movies and their release year
 */

SELECT id, title FROM movies 
WHERE id = 6;

SELECT title, year FROM movies
WHERE year BETWEEN 2000 AND 2010;

SELECT title, year FROM movies
WHERE year < 2000 OR year > 2010;

SELECT title, year FROM movies
WHERE year <= 2003;


-- Exercise 3

-- SQL Lesson 3: Queries with constraints (Pt. 2)

/* 
1.Find all the Toy Story movies
2.Find all the movies directed by John Lasseter
3.Find all the movies (and director) not directed by John Lasseter
4.Find all the WALL-* movies
 */

SELECT * FROM movies
WHERE Title LIKE "Toy Story%";

SELECT * FROM movies
WHERE Director LIKE "John Lasseter";-- or WHERE Director = "John Lasseter"

SELECT * FROM movies
WHERE Director NOT LIKE "John Lasseter";

--  % 是匹配一个或多个字符; _ 是匹配一个字符  AN_ 只匹配AND 不匹配AN
SELECT * FROM movies
WHERE Title  LIKE "WALL-%";

-- Exercise 4

-- SQL Lesson 4: Filtering and sorting Query results

/* 
1.List the last four Pixar movies released (ordered from most recent to least)
2.List the first five Pixar movies sorted alphabetically
3.List the next five Pixar movies sorted alphabetically
4.List all directors of Pixar movies (alphabetically), without duplicates
 */

-- ** DISTINCT 关键词用来去除重复行 放在SELECT 后

-- ASC 正序 DESC倒序

-- LIMIT  num_limit 选择要返回的行数;  OFFSET num_offset 指定从哪里开始计算行数

SELECT DISTINCT director FROM movies
ORDER BY director ASC;

SELECT  * FROM movies
ORDER BY Year DESC
Limit  4;

SELECT  * FROM movies
ORDER BY Title ASC
Limit  5;


-- Exercise 5

-- SQL Review: Simple SELECT Queries

/* 
1.List all the Canadian cities and their populations
2.Order all the cities in the United States by their latitude from north to south
3.List all the cities west of Chicago, ordered from west to east
4.List the two largest cities in Mexico (by population)
5.List the third and fourth largest cities (by population) in the United States and their population
 */

SELECT City, Population FROM north_american_cities
WHERE Country="Canada";

SELECT city FROM north_american_cities
WHERE Country="United States"
Order by Latitude DESC ;

SELECT city, longitude FROM north_american_cities
WHERE longitude < -87.629798
ORDER BY longitude ASC;

SELECT city, population FROM north_american_cities
WHERE country LIKE "Mexico"
ORDER BY population DESC
LIMIT 2;

SELECT city, population FROM north_american_cities
WHERE country LIKE "United States"
ORDER BY population DESC
LIMIT 2 OFFSET 2;

-- Exercise 6

-- SQL Lesson 6: Multi-table queries with JOINs

/* 
1.Find the domestic and international sales for each movie
2.Show the sales numbers for each movie that did better internationally rather than domestically
3.List all the movies by their ratings in descending order
 */

--INNER JOIN = JOIN
SELECT title, domestic_sales, international_sales 
FROM movies
  JOIN boxoffice
    ON movies.id = boxoffice.movie_id;

SELECT title, domestic_sales, international_sales
FROM movies
  JOIN boxoffice
    ON movies.id = boxoffice.movie_id
WHERE international_sales > domestic_sales;

SELECT title, rating
FROM movies
  JOIN boxoffice
    ON movies.id = boxoffice.movie_id
ORDER BY rating DESC;

-- Exercise 7

-- SQL Lesson 7: OUTER JOINs

/* 
1.Find the list of all buildings that have employees
2.Find the list of all buildings and their capacity
3.List all buildings and the distinct employee roles in each building (including empty buildings)
 */

SELECT DISTINCT building FROM employees;

SELECT * FROM buildings;

--LEFT JOIN 需要加强 理解
SELECT DISTINCT building_name, role 
FROM buildings 
  LEFT JOIN employees
    ON building_name = building;


-- Exercise 8

-- SQL Lesson 8: A short note on NULLs

/* 
1.Find the name and role of all employees who have not been assigned to a building
2.Find the names of the buildings that hold no employees
 */

SELECT name, role FROM employees
WHERE building IS NULL;

SELECT DISTINCT building_name
FROM buildings 
  LEFT JOIN employees
    ON building_name = building
WHERE role IS NULL;-- WHERE Building IS NULL

-- Exercise 9

-- SQL Lesson 9: Queries with expressions

/* 
1.List all movies and their combined sales in millions of dollars
2.List all movies and their ratings in percent
3.List all movies that were released on even number years
 */

SELECT title, (domestic_sales + international_sales) / 1000000 AS gross_sales_millions
FROM movies
  JOIN boxoffice
    ON movies.id = boxoffice.movie_id;

SELECT title, rating * 10 AS rating_percent
FROM movies
  JOIN boxoffice
    ON movies.id = boxoffice.movie_id;

SELECT title, year
FROM movies
WHERE year % 2 = 0;

-- Exercise 10

-- Queries with aggregates (Pt. 1)

/* 
1.Find the longest time that an employee has been at the studio
2.For each role, find the average number of years employed by employees in that role
3.Find the total number of employee years worked in each building
 */

SELECT MAX(years_employed) as Max_years_employed
FROM employees;

SELECT role, AVG(years_employed) as Average_years_employed
FROM employees
GROUP BY role;

SELECT building, SUM(years_employed) as Total_years_employed
FROM employees
GROUP BY building;

-- Exercise 11

-- SQL Lesson 11: Queries with aggregates (Pt. 2)

/* 
1.Find the number of Artists in the studio (without a HAVING clause)
2.Find the number of Employees of each role in the studio
3.Find the total number of years employed by all Engineers
 */

SELECT role, COUNT(*) as Number_of_artists
FROM employees
WHERE role = "Artist";

SELECT role, COUNT(*)
FROM employees
GROUP BY role;

SELECT role, SUM(years_employed)
FROM employees
GROUP BY role
HAVING role = "Engineer";

-- Exercise 12

-- 

/* 

 */


-- Exercise 13

-- 

/* 

 */