DROP TABLE IF EXISTS `topos`;
DROP TABLE IF EXISTS `vnfs`;
DROP TABLE IF EXISTS `cpes`;
DROP TABLE IF EXISTS `users`;

--
-- Table structure for table `user`
--

CREATE TABLE `users` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `level` int(8) NOT NULL DEFAULT '0',
  `username` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Table structure for table `cpe`
--

CREATE TABLE `cpes` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `new` varchar(5) NOT NULL,
  `type` int NOT NULL,
  `userid` int(64) NOT NULL,
  `topoid` int(64) NOT NULL,
  `macaddr` varchar(255) NOT NULL UNIQUE,
  `mgmtip` varchar(255) NOT NULL UNIQUE,
  `mgmtgw` varchar(255) NOT NULL,
  `publicip` varchar(255) NOT NULL UNIQUE,
  `publicgw` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `location` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

--
-- Table structure for table `vnf`
--

CREATE TABLE `vnfs` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(255) NOT NULL,
  `vnftype` varchar(255) NOT NULL,
  `flavor` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `version` varchar(255) NOT NULL,
  `template` varchar(255) NOT NULL,
  `distro` varchar(255) NOT NULL,
  `release` varchar(255) NOT NULL,
  `arch` varchar(255) NOT NULL,
  `startcmd` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

--
-- Table structure for table `topology`
--

CREATE TABLE `topos` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `cpeid` int(64) NOT NULL UNIQUE,
  `userid` int(64) NOT NULL,
  `name` varchar(255) NOT NULL,
  `topogui` text NOT NULL,
  `topoorch` text NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`cpeid`) REFERENCES `cpes`(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;




INSERT INTO `users` (`level`,`username`,`password`) VALUES (1,"steven","password");
INSERT INTO `users` (`level`,`username`,`password`) VALUES (1,"patrick","password");
INSERT INTO `users` (`level`,`username`,`password`) VALUES (1,"enterprise","password");


INSERT INTO `cpes` (`created_at`,`new`,`type`,`userid`,`topoid`,`mgmtip`,`macaddr`,`publicip`,`publicgw`,`name`,`location`) VALUES (NOW(),"TRUE",0,3,0,"10.250.3.62/16","bc:30:5b:f9:37:50","192.168.60.1/24","192.168.60.2","PACPE","PaloAlto");
INSERT INTO `cpes` (`created_at`,`new`,`type`,`userid`,`topoid`,`mgmtip`,`macaddr`,`publicip`,`publicgw`,`name`,`location`) VALUES (NOW(),"TRUE",0,3,0,"10.250.4.63/16","bc:30:5b:f9:37:51","192.168.70.1/24","192.168.70.2","SFCPE","SanFrancisco");
INSERT INTO `cpes` (`created_at`,`new`,`type`,`userid`,`topoid`,`mgmtip`,`macaddr`,`publicip`,`publicgw`,`name`,`location`) VALUES (NOW(),"TRUE",0,3,0,"10.250.4.64/16","bc:30:5b:f9:37:52","192.168.80.1/24","192.168.80.2","NYCPE","NewYork");


INSERT INTO `vnfs` (`name`,`vnftype`,`flavor`,`image`,`version`,`template`,`distro`,`release`,`arch`,`startcmd`) VALUES ("Linux Router","router","docker","Ubuntu","14.04","","","","","/bin/bash -c \"while true; do sleep 1; done\"");
INSERT INTO `vnfs` (`name`,`vnftype`,`flavor`,`image`,`version`,`template`,`distro`,`release`,`arch`,`startcmd`) VALUES ("Linux Firewall","firewall","docker","Ubuntu","14.04","","","","","/bin/bash -c \"while true; do sleep 1; done\"");
INSERT INTO `vnfs` (`name`,`vnftype`,`flavor`,`image`,`version`,`template`,`distro`,`release`,`arch`,`startcmd`) VALUES ("Apache Traffic Server","cache","docker","Ubuntu","14.04","","","","","");

