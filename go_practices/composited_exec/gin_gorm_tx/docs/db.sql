create
database gorm_exec_db default character set utf8mb4 collate utf8mb4_bin;

CREATE TABLE `m_goods`
(
    `goodsId`   bigint UNSIGNED NOT NULL COMMENT '商品id',
    `goodsName` varchar(200) NOT NULL DEFAULT '' COMMENT '商品名称',
    `stock`     int UNSIGNED NOT NULL DEFAULT '0' COMMENT '库存数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

INSERT INTO `m_goods` (`goodsId`, `goodsName`, `stock`)
VALUES (1, '施丹兰蜂蜜牛奶手工皂', 10);

ALTER TABLE `m_goods` ADD PRIMARY KEY (`goodsId`);


CREATE TABLE `m_order`
(
    `orderId`   bigint UNSIGNED NOT NULL COMMENT '订单id',
    `userId`    bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
    `salePrice` decimal(10, 0) NOT NULL DEFAULT '0' COMMENT '订单金额'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';
ALTER TABLE `m_order`
    ADD PRIMARY KEY (`orderId`),
    ADD KEY `userId` (`userId`);

CREATE TABLE `m_order_goods`
(
    `ogId`    bigint UNSIGNED NOT NULL COMMENT 'id',
    `orderId` bigint UNSIGNED NOT NULL COMMENT '订单id',
    `goodsId` bigint UNSIGNED NOT NULL COMMENT '商品id',
    `buyNum`  int UNSIGNED NOT NULL COMMENT '购买数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单商品表';
ALTER TABLE `m_order_goods`
    ADD PRIMARY KEY (`ogId`),
    ADD UNIQUE KEY `orderId_2` (`orderId`,`goodsId`);

ALTER TABLE `m_goods` MODIFY `goodsId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id', AUTO_INCREMENT=2;

ALTER TABLE `m_order` MODIFY `orderId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单id';

ALTER TABLE `m_order_goods` MODIFY `ogId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id';