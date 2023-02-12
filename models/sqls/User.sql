DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `userid` int(11) NOT NULL AUTO_INCREMENT,        -- 自增主键id
                        `username` varchar(20) NOT NULL,             -- 用户名
                        `password` varchar(40) NOT NULL,                 -- 用户密码
                        `email` varchar(50) NOT NULL ,                    -- 用户邮箱
                        `pic` blob, -- 用户头像
                        PRIMARY KEY (`userid`),
                        UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;