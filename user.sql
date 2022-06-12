create table user
(
    uid      bigint auto_increment
        primary key,
    username varchar(20)  default '' null,
    password varchar(100) default '' null
);

