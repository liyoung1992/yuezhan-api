/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50505
Source Host           : 192.168.80.221:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-03-28 13:07:33
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `qq` varchar(48) NOT NULL COMMENT 'qq',
  `phone` varchar(36) NOT NULL COMMENT '电话号码',
  `userName` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `salt` varchar(255) DEFAULT NULL COMMENT '盐',
  `updateTime` int(11) NOT NULL COMMENT '更新时间',
  `createTime` int(11) DEFAULT NULL COMMENT '创建时间',
  `lastLoginTime` int(11) DEFAULT NULL COMMENT '最后登陆时间',
  `teamId` int(11) DEFAULT NULL COMMENT '战队id',
  `lastLoginIp` varchar(255) DEFAULT NULL,
  `imageUrl` varchar(255) DEFAULT NULL COMMENT '头像url',
  PRIMARY KEY (`id`,`email`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
