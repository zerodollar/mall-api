-- http://blog.csdn.net/mu_tou/article/details/3961977
create database mall DEFAULT CHARACTER SET utf8;
use mall;
drop table if EXISTS SYS_MENU_INF;

CREATE TABLE IF NOT EXISTS  SYS_MENU_INF
(
  CATEGORY  VARCHAR(20) not null, -- 菜单分类， 
  ID      VARCHAR(20) not null,   -- id
  NAME VARCHAR(200) not null,     -- 名字 
  REMARK  VARCHAR(200),           -- 备注
  PARENT_ID VARCHAR(20),          -- 上级菜单id，1级菜单是本身id
  LEVEL   VARCHAR(2) not null,    -- 菜单级别 1/2/3
  POS     int,                    -- 同级菜单位置，相同以name字母序
  URL     VARCHAR(200),           -- 菜单对应的链接
  ICON    VARCHAR(100),           -- 菜单的图标
  STATUS       VARCHAR(4) default 'Y', -- 菜单是否生效
  CREATE_DATE  DATETIME default NOW(), -- 创建日期
  CREATE_EMP   VARCHAR(10) ,           -- 创建人
  MODIFY_DATE  DATETIME,                -- 修改日期
  MODIFY_EMP   VARCHAR(10)             -- 修改人
);
ALTER TABLE SYS_MENU_INF ADD PRIMARY KEY ( ID ) ;
insert into SYS_MENU_INF values ('NAV', '1','衣服','REMARK','1','01',1,'url','heart text-red','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '2','家电','REMARK','2','01',2,'url','heart text-red','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '3','珠宝','REMARK','3','01',3,'url','heart text-red','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '4','办公','REMARK','4','01',5,'url','heart text-red','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '5','学习','REMARK','5','01',4,'url','heart text-red','Y',NOW(),'admin',null,null);


insert into SYS_MENU_INF values ('NAV', '101','男装','REMARK','1','02',1,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '102','女装','REMARK','1','02',2,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '103','内衣','REMARK','1','02',3,'url','circle-o','Y',NOW(),'admin',null,null);

insert into SYS_MENU_INF values ('NAV', '10101','衬衫','REMARK','101','03',1,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '10102','裤子','REMARK','101','03',2,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '10103','鞋子','REMARK','101','03',3,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '10104','袜子','REMARK','101','03',4,'url','circle-o','Y',NOW(),'admin',null,null);
insert into SYS_MENU_INF values ('NAV', '10105','领带','REMARK','101','03',5,'url','circle-o','Y',NOW(),'admin',null,null);
