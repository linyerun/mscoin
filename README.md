# 介绍

# 数据库表
```sql
CREATE TABLE `member`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `ali_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `qr_code_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `appeal_success_times` int(0) NOT NULL,
  `appeal_times` int(0) NOT NULL,
  `application_time` bigint(0) NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `bank` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `branch` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `card_no` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `certified_business_apply_time` bigint(0) NOT NULL,
  `certified_business_check_time` bigint(0) NOT NULL,
  `certified_business_status` int(0) NOT NULL,
  `channel_id` int(0) NOT NULL DEFAULT 0,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `first_level` int(0) NOT NULL,
  `google_date` bigint(0) NOT NULL,
  `google_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `google_state` int(0) NOT NULL DEFAULT 0,
  `id_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `inviter_id` bigint(0) NOT NULL,
  `is_channel` int(0) NOT NULL DEFAULT 0,
  `jy_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `last_login_time` bigint(0) NOT NULL,
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `country` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `district` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `login_count` int(0) NOT NULL,
  `login_lock` int(0) NOT NULL,
  `margin` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `member_level` int(0) NOT NULL,
  `mobile_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `promotion_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `publish_advertise` int(0) NOT NULL,
  `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `real_name_status` int(0) NOT NULL,
  `registration_time` bigint(0) NOT NULL,
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `second_level` int(0) NOT NULL,
  `sign_in_ability` tinyint(4) NOT NULL DEFAULT b'1',
  `status` int(0) NOT NULL,
  `third_level` int(0) NOT NULL,
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `token_expire_time` bigint(0) NOT NULL,
  `transaction_status` int(0) NOT NULL,
  `transaction_time` bigint(0) NOT NULL,
  `transactions` int(0) NOT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `qr_we_code_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `wechat` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `local` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `integration` bigint(0) NOT NULL DEFAULT 0,
  `member_grade_id` bigint(0) NOT NULL DEFAULT 1 COMMENT '等级id',
  `kyc_status` int(0) NOT NULL DEFAULT 0 COMMENT 'kyc等级',
  `generalize_total` bigint(0) NOT NULL DEFAULT 0 COMMENT '注册赠送积分',
  `inviter_parent_id` bigint(0) NOT NULL DEFAULT 0,
  `super_partner` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `kick_fee` decimal(19, 2) NOT NULL,
  `power` decimal(8, 4) NOT NULL DEFAULT 0.0000 COMMENT '个人矿机算力(每日维护)',
  `team_level` int(0) NOT NULL DEFAULT 0 COMMENT '团队人数(每日维护)',
  `team_power` decimal(8, 4) NOT NULL DEFAULT 0.0000 COMMENT '团队矿机算力(每日维护)',
  `member_level_id` bigint(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UK_gc3jmn7c2abyo3wf6syln5t2i`(`username`) USING BTREE,
  UNIQUE INDEX `UK_10ixebfiyeqolglpuye0qb49u`(`mobile_phone`) USING BTREE,
  INDEX `FKbt72vgf5myy3uhygc90xna65j`(`local`) USING BTREE,
  INDEX `FK8jlqfg5xqj5epm9fpke6iotfw`(`member_level_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;
```

# gorm
- `db.Table().Updates()`时，不会自行更新`updated_at`字段，而使用`db.Model`时，它会更新字段。
- gorm.DB是否并发安全。https://juejin.cn/post/7134002645651439630，gorm.Open获得的DB正常用是并发安全的


# commit message规范
格式: `<type>(<scope>): <subject>`
## type(必须)
用于说明git commit的类别，只允许使用下面的标识。

feat：新功能（feature）。

fix/to：修复bug，可以是QA发现的BUG，也可以是研发自己发现的BUG。
- fix：产生diff并自动修复此问题。适合于一次提交直接修复问题
- to：只产生diff不自动修复此问题。适合于多次提交。最终修复问题提交时使用fix

docs：文档（documentation）。

style：格式（不影响代码运行的变动）。

refactor：重构（即不是新增功能，也不是修改bug的代码变动）。

perf：优化相关，比如提升性能、体验。

test：增加测试。

chore：构建过程或辅助工具的变动。

revert：回滚到上一个版本。

merge：代码合并。

sync：同步主线或分支的Bug。

## scope(可选)
scope用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

例如在Angular，可以是location，browser，compile，compile，rootScope， ngHref，ngClick，ngView等。如果你的修改影响了不止一个scope，你可以使用*代替。

## subject(必须)
subject是commit目的的简短描述，不超过50个字符。

建议使用中文（感觉中国人用中文描述问题能更清楚一些）。

## 例子
- fix(DAO):用户查询缺少username属性
- feat(Controller):用户查询接口开发