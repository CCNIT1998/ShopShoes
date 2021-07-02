ALTER TABLE
  `products`
ADD
  FOREIGN KEY (`madein`) REFERENCES `countries` (`code`);
