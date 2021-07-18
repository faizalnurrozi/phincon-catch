/*
 Navicat Premium Data Transfer

 Source Server         : MySQL Development (Container)
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : 208.87.132.133:3388
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 18/07/2021 18:13:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pokemon
-- ----------------------------
DROP TABLE IF EXISTS `pokemon`;
CREATE TABLE `pokemon` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pokemon_id` int(11) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) TYPE=InnoDB AUTO_INCREMENT=18;

-- ----------------------------
-- Records of pokemon
-- ----------------------------
BEGIN;
INSERT INTO `pokemon` VALUES (14, 3, 'Venusaur\'s Faizal', '2021-07-18 08:51:54', '2021-07-18 11:04:29', NULL);
INSERT INTO `pokemon` VALUES (15, 2, 'Ivysaur\'s Faizal', '2021-07-18 08:51:54', '2021-07-18 11:08:54', '2021-07-18 11:08:54');
INSERT INTO `pokemon` VALUES (16, 9, 'Blastosie\'s Faizal', '2021-07-18 10:33:56', '2021-07-18 11:04:58', NULL);
INSERT INTO `pokemon` VALUES (17, 6, 'Charizard\'s Faizal', '2021-07-18 11:09:30', '2021-07-18 11:09:30', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
