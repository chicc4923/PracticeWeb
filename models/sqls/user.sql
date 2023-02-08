DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `userid` int(11) NOT NULL AUTO_INCREMENT,        -- 自增主键id
                        `username` varchar(20) DEFAULT NULL,             -- 用户名
                        `password` varchar(40) NOT NULL,                 -- 用户密码
                        `email` varchar(50) NOT NULL,                    -- 用户邮箱
                        `status` int(11) DEFAULT '0',                    -- 用户是否通过邮箱验证 , 0未通过, 1通过
                        `newdate` varchar(15) NOT NULL,                  -- 注册时间
                        PRIMARY KEY (`userid`),
                        UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;