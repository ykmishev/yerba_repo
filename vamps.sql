create table Customer(
    id int NOT NULL,
    name VARCHAR(50),
    referee_id int);

insert into Customer(id, name, referee_id) values (1,'Will', NULL),
(2, 'Jane', NULL),
(3, 'Alex', 2),
(4, 'Bill', NULL),
(5, 'Zack', 1),
(6, 'Mark', 2);