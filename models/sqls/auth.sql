CREATE TABLE `auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '用户名',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `PracticeWeb`.`auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');