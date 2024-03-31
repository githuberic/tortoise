CREATE TABLE `user`(
    `userId`   bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username` varchar(100) NOT NULL DEFAULT '' COMMENT 'username',
    `password` varchar(100) NOT NULL DEFAULT '' COMMENT 'password',
    `addTime`  datetime     NOT NULL DEFAULT '2020-12-15 00:00:00' COMMENT '注册时间',
    PRIMARY KEY (`userId`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4  COMMENT='用户表'

INSERT INTO `user` (`userId`, `username`, `password`, `addTime`) VALUES
(1, 'lhd', '$2a$10$adv/A7yGhOOTUUIt4vkpIufuIkABiCKeRxxKD0do1dgnqJqSFvrNq', '2020-12-15 00:00:00');

INSERT INTO `user` (`userId`, `username`, `password`, `addTime`) VALUES
    (2, 'lgq', '$2a$10$u99/hKzdWwwvJlt2H/.47.E3OT5SqyMEzjXuzwcIBW.774NXFspyS', '2022-06-08 23:35:12');

/*密码:123,是bcrypt加密后的*/