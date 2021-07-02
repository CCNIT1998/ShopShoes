CREATE TABLE `products` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `image` varchar(255) NOT NULL,
  `price_old` float,
  `price_current` float,
  `sale` int,
  `madein` varchar(50),
  `category_id` int,
  `historical_sold` int,
  `rating_star` int,
  `quantity` int
);

CREATE TABLE `categories` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);

CREATE TABLE `countries` (
  `code` varchar(5) PRIMARY KEY NOT NULL,
  `name` varchar(255) NOT NULL
);
ALTER TABLE
  `products`
ADD
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`ID`);
