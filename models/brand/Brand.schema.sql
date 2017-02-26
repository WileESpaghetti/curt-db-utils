Brand | CREATE TABLE `Brand` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `logoAlt` varchar(255) DEFAULT NULL,
  `formalName` varchar(255) DEFAULT NULL,
  `longName` varchar(255) DEFAULT NULL,
  `primaryColor` varchar(10) DEFAULT NULL,
  `autocareID` varchar(4) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT