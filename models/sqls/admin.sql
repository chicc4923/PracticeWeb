DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
                          `adminid` int(11) NOT NULL AUTO_INCREMENT,
                          `adminname` varchar(30) NOT NULL,
                          `password` varchar(40) NOT NULL,
                          PRIMARY KEY (`adminid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
