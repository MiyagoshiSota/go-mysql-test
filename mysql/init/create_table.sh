#!/bin/sh

CMD_MYSQL="mysql -u root -p${MYSQL_ROOT_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table Player (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    name varchar(50) NOT NULL
    );"
$CMD_MYSQL -e "create table History (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    player_id int(10) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME
)"

$CMD_MYSQL -e  "insert into Player values (1,'miyagoshi');"
$CMD_MYSQL -e  "insert into History values (1, 1,cast('2009-08-03 23:58:01' as datetime),cast('2009-08-03 23:59:01' as datetime));"