/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : 127.0.0.1:3306
 Source Schema         : go_test

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 10/08/2020 20:22:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cron
-- ----------------------------
DROP TABLE IF EXISTS `cron`;
CREATE TABLE `cron`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `exp` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `command` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `is_del` tinyint(1) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cron
-- ----------------------------
INSERT INTO `cron` VALUES (9, 'LOL', '每8秒LOL一次', '*/8 * * * * * *', 'sleep 1', '2020-08-10 19:37:19', NULL, 0);
INSERT INTO `cron` VALUES (8, '拳击', '每5秒拳击一次', '*/5 * * * * * *', 'sleep 1', '2020-08-10 19:36:10', NULL, 0);
INSERT INTO `cron` VALUES (7, '睡眠', '每10秒睡眠一次', '*/10 * * * * * *', 'slep 1', '2020-08-10 14:07:23', NULL, 0);
INSERT INTO `cron` VALUES (6, '打印', '每四秒打印一次', '*/4 * * * * * *', 'echo 1', '2020-08-10 14:06:23', NULL, 0);

SET FOREIGN_KEY_CHECKS = 1;
