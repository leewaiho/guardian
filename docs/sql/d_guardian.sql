/*
 Navicat Premium Data Transfer

 Source Server         : ECS
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 10.10.0.1:3306
 Source Schema         : d_guardian

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 09/02/2020 14:02:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for network
-- ----------------------------
DROP TABLE IF EXISTS `network`;
CREATE TABLE `network`
(
    `id`           int(11)     NOT NULL AUTO_INCREMENT,
    `name`         varchar(50) NOT NULL COMMENT '网络名称; 仅做展示用',
    `created_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb4 COMMENT 'VPN网络表';

-- ----------------------------
-- Table structure for node
-- ----------------------------
DROP TABLE IF EXISTS `node`;
CREATE TABLE `node`
(
    `id`                   int(10) unsigned NOT NULL AUTO_INCREMENT,
    `network_id`           int(10) unsigned NOT NULL,
    `hostname`             varchar(40)      NOT NULL COMMENT '主机标识; 仅做展示用',
    `vpn_ip`               varchar(18)      NOT NULL COMMENT 'VPNIP 使用CIDR表示; 例如: 10.10.0.1/32',
    `type`                 tinyint(4)       NOT NULL DEFAULT '2' COMMENT '节点类型; 1:网关 2:普通网络节点',
    `listen_ip`            varchar(15)      NOT NULL DEFAULT '' COMMENT '网关节点的公网IP',
    `listen_port`          varchar(20)      NOT NULL DEFAULT '' COMMENT '网关节点wireguard服务的公网端口',
    `dns`                  varchar(15)      NOT NULL DEFAULT '' COMMENT '节点使用的DNS服务器地址',
    `private_key`          varchar(80)      NOT NULL COMMENT 'wireguard 私钥',
    `public_key`           varchar(80)      NOT NULL COMMENT 'wireguard 公钥',
    `allowed_ips`          varchar(200)              DEFAULT '' COMMENT '转发至当前节点的CIDR',
    `persistent_keepalive` int(11)          NOT NULL DEFAULT '25' COMMENT '心跳包间隔',
    `created_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_disabled`          tinyint(4)       NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 12
  DEFAULT CHARSET = utf8mb4 COMMENT ='节点表';

SET FOREIGN_KEY_CHECKS = 1;
