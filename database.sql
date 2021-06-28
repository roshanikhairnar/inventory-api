create database coditation;
use coditation;
create table Category (categoryID int not null,categoryName varchar(255),PRIMARY KEY (categoryID));
create table Subcategory(subcategoryID int not null, categoryID int, subcategoryName varchar(255),primary key(subcategoryID),
FOREIGN KEY (categoryID) REFERENCES category(categoryID));
create table Products(productID int not null,subcategoryID int,productName varchar(255),productPrice int,primary key(productID),
foreign key(subcategoryID) references Subcategory(subcategoryID));

INSERT INTO `coditation`.`category`
(`categoryID`,
`categoryName`)
VALUES
(1,'electronics');
INSERT INTO `coditation`.`category`
(`categoryID`,
`categoryName`)
VALUES
(2,'home');
INSERT INTO `coditation`.`category`
(`categoryID`,
`categoryName`)
VALUES
(3,'Fashion');
INSERT INTO `coditation`.`category`
(`categoryID`,
`categoryName`)
VALUES
(4,'Books');
INSERT INTO `coditation`.`category`
(`categoryID`,
`categoryName`)
VALUES
(5,'Books');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(1,1,'camera');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(2,1,'TV');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(3,2,'Sofa');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(4,2,'Table');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(5,3,'Tshirt');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(6,3,'kurta');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(7,4,'fiction');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(8,4,'biography');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(9,5,'vegetable');
INSERT INTO `coditation`.`subcategory`
(`subcategoryID`,
`categoryID`,
`subcategoryName`)
VALUES
(10,5,'snacks');

INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('1', '1', 'Nikon', '20000');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('2', '2', 'LG tv', '20000');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('3', '3', 'pepperfry sofa', '20000');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('4', '4', 'pepperfry table', '10000');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('5', '5', 'Levis Tshirt', '500');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('6', '6', 'vaani kurta', '700');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('7', '7', 'one day in london book', '300');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('8', '8', 'MK gandhi', '600');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('9', '9', 'coriander', '30');
INSERT INTO `coditation`.`products` (`productID`, `subcategoryID`, `productName`, `productPrice`) VALUES ('10', '10', 'balaji chips', '30');




