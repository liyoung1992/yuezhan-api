/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50505
Source Host           : 192.168.80.221:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-03-28 13:07:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_game
-- ----------------------------
DROP TABLE IF EXISTS `tb_game`;
CREATE TABLE `tb_game` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `team1` int(11) NOT NULL,
  `team2` int(11) NOT NULL,
  `userId` int(11) DEFAULT NULL COMMENT '创建人',
  `createTime` int(11) DEFAULT NULL,
  `imageUrl` varchar(128) DEFAULT NULL COMMENT '比赛图片',
  `updateTime` int(11) DEFAULT NULL COMMENT '更新时间',
  `startTime` int(11) DEFAULT NULL COMMENT '比赛开始时间',
  `regionId` int(11) DEFAULT NULL COMMENT '大区id(可选)',
  `endTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`,`team1`,`team2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
