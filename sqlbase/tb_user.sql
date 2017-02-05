/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50173
Source Host           : 192.168.2.221:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2017-02-03 20:20:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `salt` varchar(255) DEFAULT NULL COMMENT '盐',
  `updateTime` int(11) DEFAULT NULL COMMENT '更新时间',
  `createTime` int(11) DEFAULT NULL COMMENT '创建时间',
  `lastLoginTime` int(11) DEFAULT NULL COMMENT '最后登陆时间',
  `teamId` int(11) DEFAULT NULL COMMENT '战队id',
  `imageUrl` varchar(255) DEFAULT NULL COMMENT '头像url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
