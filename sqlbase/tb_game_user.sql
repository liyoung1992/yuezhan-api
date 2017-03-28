/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50505
Source Host           : 192.168.80.221:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-03-28 13:07:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_game_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_game_user`;
CREATE TABLE `tb_game_user` (
  `id` int(11) NOT NULL,
  `gameId` int(11) NOT NULL,
  `userId` int(11) NOT NULL,
  PRIMARY KEY (`id`,`gameId`,`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='比赛和用户的关联表';
