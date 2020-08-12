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

 Date: 12/08/2020 16:16:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for daemon_cron
-- ----------------------------
DROP TABLE IF EXISTS `daemon_cron`;
CREATE TABLE `daemon_cron`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `command` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `is_killer` tinyint(3) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of daemon_cron
-- ----------------------------
INSERT INTO `daemon_cron` VALUES (11, '输出2', '一直输出2', 'echo 2', '2020-08-11 16:21:10', '2020-08-11 16:21:36', 0);
INSERT INTO `daemon_cron` VALUES (12, '输出3', '一直输出3', 'echo 3', '2020-08-11 16:23:09', '2020-08-11 16:46:20', 0);
INSERT INTO `daemon_cron` VALUES (6, '打印', '每四秒打印一次', 'echo 1', '2020-08-10 14:06:23', NULL, 0);

SET FOREIGN_KEY_CHECKS = 1;


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
  `is_killer` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cron
-- ----------------------------
INSERT INTO `cron` VALUES (9, 'LOL', '每8秒LOL一次', '*/8 * * * * * *', 'sleep 1', '2020-08-10 19:37:19', NULL, 0);
INSERT INTO `cron` VALUES (8, '拳击', '每5秒拳击一次', '*/5 * * * * * *', 'sleep 1', '2020-08-10 19:36:10', NULL, 0);
INSERT INTO `cron` VALUES (7, '睡眠', '每10秒睡眠一次', '*/10 * * * * * *', 'slep 1', '2020-08-10 14:07:23', NULL, 0);
INSERT INTO `cron` VALUES (6, '打印', '每四秒打印一次', '*/4 * * * * * *', 'echo 1', '2020-08-10 14:06:23', NULL, 0);

SET FOREIGN_KEY_CHECKS = 1;


-- ----------------------------
-- Table structure for once_cron
-- ----------------------------
DROP TABLE IF EXISTS `once_cron`;
CREATE TABLE `once_cron`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `command` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `run_at` timestamp(0) NULL DEFAULT NULL COMMENT '运行时间',
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `is_killer` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of once_cron
-- ----------------------------
INSERT INTO `once_cron` VALUES (9, 'LOL', '每8秒LOL一次', 'sleep 1', '2020-08-12 11:41:13', '2020-08-10 19:37:19', '2020-08-12 11:41:18', 0);
INSERT INTO `once_cron` VALUES (8, '拳击', '每5秒拳击一次', 'sleep 1', '2020-08-13 11:41:20', '2020-08-10 19:36:10', '2020-08-12 11:41:23', 0);

SET FOREIGN_KEY_CHECKS = 1;