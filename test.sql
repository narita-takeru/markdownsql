DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL primary key comment 'user identifier',
  `name` varchar(255) NOT NULL comment 'user name',
  `created_at` datetime NOT NULL DEFAULT current_timestamp comment 'record created.',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp on update current_timestamp comment 'record updated.',
  INDEX(`name`)
) ENGINE = InnoDB DEFAULT CHARSET utf8;

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` bigint NOT NULL primary key comment 'product identifier',
  `category` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` integer NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp comment 'record created.',
  `updated_at` datetime NOT NULL DEFAULT current_timestamp on update current_timestamp comment 'record updated.',
  UNIQUE(`category`,`name`)
) ENGINE = InnoDB DEFAULT CHARSET utf8;


