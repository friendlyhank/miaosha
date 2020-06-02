/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50730
Source Host           : 127.0.0.1:3306
Source Database       : miaosha

Target Server Type    : MYSQL
Target Server Version : 50730
File Encoding         : 65001

Date: 2020-06-02 18:11:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `oid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品oid',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `price` bigint(20) DEFAULT '0' COMMENT '价格',
  `stocknum` int(11) DEFAULT '0' COMMENT '商品库存',
  `description` varchar(255) DEFAULT NULL COMMENT '商品描述',
  `status` tinyint(4) DEFAULT '1' COMMENT '订单状态',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单创建时间',
  PRIMARY KEY (`oid`)
) ENGINE=InnoDB AUTO_INCREMENT=100002 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES ('100000', '华为 HUAWEI nova 7', '294900', '1000', '华为 HUAWEI nova 7 5G 6400万后置四摄 5G SoC芯片 OLED极点全面屏 8GB+128GB 蜜语红·星耀版', '1', '2020-06-02 14:11:18');
INSERT INTO `goods` VALUES ('100001', '华为 HUAWEI P40', '415800', '1000', '华为 HUAWEI P40 麒麟990 5G SoC芯片 5000万超感知徕卡三摄 30倍数字变焦 6GB+128GB晨曦金全网通5G手机', '1', '2020-06-02 14:14:16');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `orderid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单号',
  `orderno` varchar(64) DEFAULT NULL COMMENT '订单编号',
  `soid` bigint(20) DEFAULT '0' COMMENT '店铺soid',
  `oid` bigint(20) DEFAULT '0' COMMENT '商品oid',
  `price` bigint(20) DEFAULT '0' COMMENT '价格',
  `num` int(11) DEFAULT '0' COMMENT '购买数量',
  `status` tinyint(4) DEFAULT '0' COMMENT '订单状态',
  `payuid` bigint(20) DEFAULT '0' COMMENT '支付用户uid',
  `paystatus` tinyint(4) DEFAULT '0' COMMENT '支付状态',
  `paytime` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '支付时间',
  `mobile` varchar(24) DEFAULT NULL COMMENT '手机号',
  `address` varchar(255) DEFAULT NULL COMMENT '收获地址',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单创建时间',
  `updatetime` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`orderid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of order
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `uid` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户uid',
  `nick` varchar(255) DEFAULT NULL COMMENT '用户昵称',
  `password` varchar(32) DEFAULT NULL COMMENT '用户密码',
  `logo` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `logintime` timestamp NULL DEFAULT NULL COMMENT '最后一次登录的时间戳',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
