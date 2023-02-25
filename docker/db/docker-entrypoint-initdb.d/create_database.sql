create database ec_web;

create user ec_web identified by '';
grant all on *.* to 'ec_web'@'%';

flush privileges;
