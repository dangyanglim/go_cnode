DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(40) NOT NULL DEFAULT '',
  `last_name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `property`;
CREATE TABLE `property` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `password` varchar(40) NOT NULL DEFAULT '',
  `name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `propertytoken`;
CREATE TABLE `propertytoken` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `token` varchar(40) NOT NULL DEFAULT '',
  `tokenExptime` varchar(80) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `propertycheckcode`;
CREATE TABLE `propertycheckcode` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `checkcode` varchar(40) NOT NULL DEFAULT '',
  `tokenExptime` varchar(80) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `usercode`;
CREATE TABLE `usercode` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `token` varchar(40) NOT NULL DEFAULT '',
  `tokenExptime` varchar(80) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `usertoken`;
CREATE TABLE `usertoken` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `token` varchar(40) NOT NULL DEFAULT '',
  `tokenExptime` varchar(80) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(40) NOT NULL DEFAULT '',
  `pwd` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `garden`;
CREATE TABLE `garden` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `garden` varchar(40) NOT NULL DEFAULT '',
  `area` varchar(40) NOT NULL DEFAULT '',
  `city` varchar(40) NOT NULL DEFAULT '',
  `province` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `building`;
CREATE TABLE `building` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `building` varchar(40) NOT NULL DEFAULT '',
  `gardenid` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `floor`;
CREATE TABLE `floor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `floor` varchar(40) NOT NULL DEFAULT '',
  `buildingid` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `room` varchar(40) NOT NULL DEFAULT '',
  `floorid` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `renter`;
CREATE TABLE `renter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `renter` varchar(40) NOT NULL DEFAULT '',
  `name` varchar(40) NOT NULL DEFAULT '',
  `timeBegin` varchar(40) NOT NULL DEFAULT '',
  `timeEnd` varchar(40) NOT NULL DEFAULT '',
  `roomid` varchar(40) NOT NULL DEFAULT '',
  `status` varchar(40) NOT NULL DEFAULT '',
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
DROP TABLE IF EXISTS `ammeter`;
CREATE TABLE `ammeter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spacetype` varchar(40) NOT NULL DEFAULT '',
  `spaceid` varchar(40) NOT NULL DEFAULT '',
  `voltage` varchar(40) NOT NULL DEFAULT '',
  `ammeter_addr` varchar(40) NOT NULL DEFAULT '',
  `current` varchar(40) NOT NULL DEFAULT '',
  `energy` varchar(40) NOT NULL DEFAULT '',
  `gateway` varchar(40) NOT NULL DEFAULT '',
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `state` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ammeter_addr` (`ammeter_addr`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `smartlock`;
CREATE TABLE `smartlock` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spacetype` varchar(40) NOT NULL DEFAULT '',
  `spaceid` varchar(40) NOT NULL DEFAULT '',
  `lock_addr` varchar(40) NOT NULL DEFAULT '',
  `gateway` varchar(40) NOT NULL DEFAULT '',
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `state` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `lock_addr` (`lock_addr`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `water`;
CREATE TABLE `water` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spacetype` varchar(40) NOT NULL DEFAULT '',
  `spaceid` varchar(40) NOT NULL DEFAULT '',
  `water_addr` varchar(40) NOT NULL DEFAULT '',
  `gateway` varchar(40) NOT NULL DEFAULT '',
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `state` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `water_addr` (`water_addr`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `gateway`;
CREATE TABLE `gateway` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spacetype` varchar(40) NOT NULL DEFAULT '',
  `spaceid` varchar(40) NOT NULL DEFAULT '',
  `gateway` varchar(40) NOT NULL DEFAULT '',
  `propertyid` varchar(40) NOT NULL DEFAULT '',
  `state` varchar(40) NOT NULL DEFAULT '',
  `name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `gateway` (`gateway`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;