SHOW DATABASES;
SHOW CREATE TABLE todolist;
use gorestfulapi_exercise;
select * from todolist;
desc todolist;

rename table `todolistcontent` to `todolist`;
alter table `todolist` rename column `namatodolist` to `todolisttitle`;
alter table `todolist` add column `todolistcontent` text after `todolisttitle`;
alter table `todolist` add column `todolisttabledate` DATETIME after `id`;
alter table `todolist` add column `todolistcontentchecked` bool after `todolistcontent`;
alter table `todolist` rename column `todolistcontentchecked` to `checked`;
alter table `todolist` add column `todolistsubcontent` text after `todolistcontent`;
alter table `todolist` MODIFY column `checked` boolean; 
alter table `todolist` modify column `datetime` timestamp not null;

delete from todolist where id="";
insert into `todolist` (`todolisttitle`, `todolistcontent`, `todolistsubcontent`, `checked`)
values
('to do list 1', 'contoh isi main to do list 2', 'sub kategori to do list', true),
('to do list 2', 'contoh isi main to do list 1', 'sub kategori to do list', true),
('to do list 2', 'contoh isi main to do list 2', 'sub kategori to do list', true),
('to do list 3', 'contoh isi main to do list 1', 'sub kategori to do list', true);

alter table `todolist` drop COLUMN `todolistsubcontent`;

/*-------------------------------------------------------------------------------------*/
create database `db_todolist`;
use `db_todolist`;
CREATE TABLE `todolist` (
  `id` int NOT NULL AUTO_INCREMENT,
  `datetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `todolisttitle` varchar(255) NOT NULL,
  `todolistcontent` text,
  `checked` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
select * from todolist;
desc todolist;

insert into `todolist` (`todolisttitle`, `todolistcontent`, `checked`)
values
('belanjaan', 'tomat x4', true),
('belanjaan', 'cabe 1Kg', true),
('belanjaan', 'brokoli 1Bonggol', true),
('belanjaan', 'tempe x4', false),
('workout', 'pushup x5', true),
('workout', 'running in place 2Min', true),
('workout', 'barbel curl x10', true),
('workout', 'situp x5', true),
('workout', 'jumping jack 2Min', true);
