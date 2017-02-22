/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50173
Source Host           : 192.168.2.221:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2017-02-03 20:28:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_team
-- ----------------------------
DROP TABLE IF EXISTS `tb_team`;
CREATE TABLE `tb_team` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `teamName` varchar(255) DEFAULT NULL COMMENT '队伍名称',
  `regionId` int(11) DEFAULT NULL COMMENT '大区id(可选)-1表示未选',
  `userId` int(11) DEFAULT NULL COMMENT '创建人',
  `leaderId` int(11) DEFAULT NULL COMMENT '队长',
  `createTime` int(11) DEFAULT NULL COMMENT '创建时间',
  `imageUrl` varchar(255) DEFAULT NULL COMMENT '头像url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;