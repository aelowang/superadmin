CREATE TABLE `manage_site_deduction` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL UNIQUE,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `site_uuid` varchar(36) NOT NULL,
  `score` decimal(20,2) NOT NULL,
  `deduction_method` varchar(20) NOT NULL,
  `deduction_status` varchar(20) NOT NULL,
  `remark` text NOT NULL,
  `site_name` varchar(100) NOT NULL,
  `price_currency` varchar(20) NOT NULL,
  `operator_uuid` varchar(36) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
