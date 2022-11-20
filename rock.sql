use test

drop table if exists `account`;
create table account (
    id bigint(20) unsigned not null auto_increment,
    username varchar(20) not null,
    password varchar(64) not null,
    status tinyint(4) default 1,
    create_time datetime not null ,
    update_time datetime not null,
    last_login_time datetime default null,
    primary key (`id`)
) engine= innodb auto_increment= 1 default charset=utf8mb4 comment = 'account info';

insert into account values (1, 'admin', 'admin', 1, '2022-11-01 00:00:00', '2022-11-01 00:00:00', '2022-11-01 00:00:00')