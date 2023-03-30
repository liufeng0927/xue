/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 80001
 Source Host           : localhost:3306
 Source Schema         : xueden_student_mangement_system

 Target Server Type    : MySQL
 Target Server Version : 80001
 File Encoding         : 65001

 Date: 12/01/2023 14:51:36
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for s_course
-- ----------------------------
DROP TABLE IF EXISTS `s_course`;
CREATE TABLE `s_course`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`   bigint(20) NOT NULL DEFAULT 0,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`   bigint(20) NOT NULL DEFAULT 0,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `course_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `course_no`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '课程信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of s_course
-- ----------------------------
INSERT INTO `s_course`
VALUES (1, 1, '2022-12-14 10:59:59', '', 1, '2022-12-14 10:59:59', '线性代数', 'c001');
INSERT INTO `s_course`
VALUES (2, 1, '2022-12-14 11:00:31', '高等数学', 1, '2022-12-14 11:00:31', '高等数学', 'c002');
INSERT INTO `s_course`
VALUES (3, 1, '2022-12-14 11:00:51', '大学英语', 1, '2022-12-14 11:00:51', '大学英语', 'c003');
INSERT INTO `s_course`
VALUES (4, 1, '2022-12-14 11:01:17', 'C语言程序员设计', 1, '2022-12-14 11:01:17', 'C语言程序员设计', 'c004');
INSERT INTO `s_course`
VALUES (5, 1, '2022-12-14 11:01:38', '计算机组成原理', 1, '2022-12-14 11:01:38', '计算机组成原理', 'c005');
INSERT INTO `s_course`
VALUES (6, 1, '2022-12-14 11:01:58', '计算机网络原理', 1, '2022-12-14 11:01:58', '计算机网络原理', 'c006');
INSERT INTO `s_course`
VALUES (7, 1, '2022-12-14 11:02:11', '编译原理', 1, '2022-12-14 11:02:11', '编译原理', 'c007');
INSERT INTO `s_course`
VALUES (8, 1, '2022-12-14 11:02:35', '大学物理', 1, '2022-12-14 11:02:35', '大学物理', 'c008');
INSERT INTO `s_course`
VALUES (9, 1, '2022-12-14 11:03:02', 'JAVA程序员设计', 1, '2022-12-14 11:03:02', 'JAVA程序员设计', 'c009');
INSERT INTO `s_course`
VALUES (10, 1, '2022-12-14 11:04:31', '数字逻辑设计', 1, '2022-12-14 11:04:31', '数字逻辑设计', 'c010');
INSERT INTO `s_course`
VALUES (11, 1, '2022-12-14 11:04:54', '数据结构与算法', 1, '2022-12-14 11:04:54', '数据结构与算法', 'c011');
INSERT INTO `s_course`
VALUES (12, 1, '2022-12-14 11:05:08', '软件工程', 1, '2022-12-14 11:05:08', '软件工程', 'c012');
INSERT INTO `s_course`
VALUES (13, 1, '2022-12-14 11:05:24', '数据库系统', 1, '2022-12-14 11:05:24', '数据库系统', 'c013');
INSERT INTO `s_course`
VALUES (14, 1, '2022-12-14 11:06:00', '操作系统', 1, '2022-12-14 11:06:00', '操作系统', 'c014');

-- ----------------------------
-- Table structure for s_grade_class
-- ----------------------------
DROP TABLE IF EXISTS `s_grade_class`;
CREATE TABLE `s_grade_class`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`   bigint(20) NOT NULL DEFAULT 0,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`   bigint(20) NOT NULL DEFAULT 0,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `clazz`       int(11) NOT NULL DEFAULT 0,
    `code`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `grade`       int(11) NOT NULL DEFAULT 0,
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '班级信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of s_grade_class
-- ----------------------------
INSERT INTO `s_grade_class`
VALUES (1, 1, '2022-12-10 17:44:50', '21级1班', 1, '2022-12-10 17:44:53', 1, '202101', 2021, '21级1班');
INSERT INTO `s_grade_class`
VALUES (2, 1, '2022-12-10 18:02:08', '21级2班', 1, '2022-12-10 18:02:11', 2, '202102', 2021, '21级2班');
INSERT INTO `s_grade_class`
VALUES (3, 1, '2022-12-10 11:32:27', '21级3班', 1, '2022-12-10 11:32:27', 3, '202103', 2021, '21级3班');
INSERT INTO `s_grade_class`
VALUES (4, 1, '2022-12-10 11:34:53', '21级4班', 1, '2022-12-10 11:34:53', 4, '202104', 2021, '21级4班');
INSERT INTO `s_grade_class`
VALUES (5, 1, '2022-12-10 11:35:23', '22级1班', 1, '2022-12-10 11:35:23', 1, '202201', 2022, '22级1班');
INSERT INTO `s_grade_class`
VALUES (6, 1, '2022-12-10 11:42:30', '22级2班', 1, '2022-12-10 11:42:30', 2, '202202', 2022, '22级2班');
INSERT INTO `s_grade_class`
VALUES (8, 1, '2022-12-10 11:50:07', '22级3班', 1, '2022-12-10 11:50:07', 3, '202203', 2022, '22级3班');
INSERT INTO `s_grade_class`
VALUES (9, 1, '2022-12-10 12:29:47', '22级4班', 1, '2022-12-10 12:32:28', 4, '202204', 2022, '22级4班');

-- ----------------------------
-- Table structure for s_student
-- ----------------------------
DROP TABLE IF EXISTS `s_student`;
CREATE TABLE `s_student`
(
    `id`             bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`      bigint(20) NOT NULL DEFAULT 0,
    `create_time`    datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`      bigint(20) NOT NULL DEFAULT 0,
    `update_time`    datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `name`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `phone`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `qq`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `sex`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `stuno`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `grade_class_id` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX            `FK8mac4aaiivjtr0kd4kpyt4vm5`(`grade_class_id` ASC) USING BTREE,
    CONSTRAINT `s_student_ibfk_1` FOREIGN KEY (`grade_class_id`) REFERENCES `s_grade_class` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '学生信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of s_student
-- ----------------------------
INSERT INTO `s_student`
VALUES (1, 1, '2022-12-11 20:39:17', '韩立', 1, '2022-12-11 20:39:23', '韩立', '13800138001', '11184629', '男',
        '202101001', 1);
INSERT INTO `s_student`
VALUES (3, 1, '2022-12-12 12:27:33', '南宫婉', 1, '2022-12-12 12:27:33', '南宫婉', '13800138001', '1165498745', '女',
        '202201001', 5);
INSERT INTO `s_student`
VALUES (4, 1, '2022-12-17 12:33:12',
        '韩立好友，与韩立在七玄门相识，两人之间的情谊难以言表，以至于厉家与韩家都结成了世交，延续多年，代代都交好如初', 1,
        '2022-12-17 12:33:12', '厉飞雨', '13800138002', '11184629', '男', '202101002', 1);
INSERT INTO `s_student`
VALUES (5, 1, '2022-12-17 12:34:28', '真名汪凝，原乱星海妙音门门主之女，韩立平生所遇第一美女、红颜知己。', 1,
        '2022-12-17 12:34:28', '紫灵仙子', '13800138003', '123456987', '女', '202201002', 5);
INSERT INTO `s_student`
VALUES (6, 1, '2022-12-17 12:35:48',
        '魁星岛附近的修仙小门派女修，对师姐妍丽重情重义，为已成鬼魂的师姐不死重生付出极大代价，与师姐修炼阴阳轮回诀被未知通道吸入灵界，后被姜老怪/青元子收为徒。',
        1, '2022-12-17 12:35:48', '元瑶', '13800138004', '236598541', '女', '202201003', 5);
INSERT INTO `s_student`
VALUES (7, 1, '2022-12-17 12:38:30', '本是灵界银月妖狼玲珑仙子两魂之一', 1, '2022-12-17 12:38:30', '银月',
        '13800138005', '8965212', '女', '202202001', 6);
INSERT INTO `s_student`
VALUES (8, 1, '2022-12-17 12:39:48', '落云宗女修，负责管理药园，韩立入落云宗修炼，归其管辖，被家族逼婚，一直推脱不肯完婚。',
        1, '2022-12-17 12:39:48', '慕沛灵', '13800138006', '5969854', '女', '202202002', 6);
INSERT INTO `s_student`
VALUES (9, 1, '2022-12-17 12:40:43', '韩立黄枫谷师姐', 1, '2022-12-17 12:40:43', '陈巧倩', '13800138007', '85969585',
        '女', '202203001', 8);
INSERT INTO `s_student`
VALUES (10, 1, '2022-12-17 12:41:30',
        '韩立旧识。为人高傲，为了驻颜修炼化春决，喜欢男人为她争风吃醋，其师傅撮合她和韩立双修，两人相互不对眼，红拂派其和韩立一起参加燕家夺宝大会',
        1, '2022-12-17 12:41:30', '董萱儿', '13800138008', '88888888', '女', '202203002', 8);
INSERT INTO `s_student`
VALUES (11, 1, '2022-12-17 12:42:14',
        '御灵宗女修，亲切可人，早年曾在太南小会上卖给韩立金竺符笔，多年后和柳眉等奉命追查至木灵婴下落，被韩立发现，韩立追上她时，发现是多年前旧识只将其打晕。',
        1, '2022-12-17 12:46:06', '菡云芝', '13800138009', '99999999', '女', '202204001', 9);
INSERT INTO `s_student`
VALUES (12, 1, '2022-12-17 12:46:47',
        '与韩立同时参加七玄门试炼，又同时拜入墨大夫门下。学习《长春功》不成，为入七玄门内门习《象甲功》。', 1,
        '2022-12-17 12:46:47', '张铁', '13800138010', '77777777', '男', '202204002', 9);
INSERT INTO `s_student`
VALUES (13, 1, '2022-12-21 06:51:12', '云狐族，没落的狐族分支', 1, '2022-12-21 06:51:12', '柳乐儿', '13800138000',
        '11184629', '女', '202101003', 1);
INSERT INTO `s_student`
VALUES (14, 1, '2022-12-21 06:52:16', '北寒石矶殿大长老弟子，真仙巅峰，于北寒仙宫任职', 1, '2022-12-21 06:52:16', '高升',
        '13800138000', '111846219', '男', '202101004', 1);
INSERT INTO `s_student`
VALUES (15, 1, '2022-12-21 06:53:21', '本体是噬金仙，爱吃仙材仙器，韩立的灵宠，与韩立一起飞升仙界。', 1,
        '2022-12-21 06:53:21', '金童', '13800138000', '11184629', '男', '202101005', 1);
INSERT INTO `s_student`
VALUES (16, 1, '2022-12-21 06:54:00', '原为丰国宰相府七小姐，被灵寰界三大宗门之一的冷焰宗看中并收为内门弟子。', 1,
        '2022-12-21 06:54:00', '余梦寒', '13800138000', '11184629', '女', '202101006', 1);
INSERT INTO `s_student`
VALUES (17, 1, '2022-12-21 06:54:42', '冷焰宗内门弟子，元婴中期修为，余梦寒的师尊。', 1, '2022-12-21 06:54:42', '古韵月',
        '13800138000', '11184629', '女', '202101007', 1);
INSERT INTO `s_student`
VALUES (18, 1, '2022-12-21 06:55:31', '天鬼宗大乘期太上长老。', 1, '2022-12-21 06:55:31', '段人离', '13800138000',
        '11184629', '男', '202101008', 1);
INSERT INTO `s_student`
VALUES (19, 1, '2022-12-21 06:56:22', '天鬼宗大乘期太上大长老。', 1, '2022-12-21 06:56:22', '童人垩', '13800138000',
        '11184629', '男', '202101009', 1);
INSERT INTO `s_student`
VALUES (20, 1, '2022-12-21 06:57:14', '伏凌宗大长老封天都之徒，真仙中期修士，于仙宫供职', 1, '2022-12-21 06:57:14',
        '方磐', '13800138000', '11184629', '男', '202101010', 1);
INSERT INTO `s_student`
VALUES (21, 1, '2022-12-21 06:58:31', '韩立烛龙道的真仙仆从。为韩立饲养念羽并且外出寻找灵药及种子。', 1,
        '2022-12-21 06:58:31', '梦浅浅', '13800138000', '11184629', '女', '202101011', 5);
INSERT INTO `s_student`
VALUES (22, 1, '2022-12-21 07:00:12', '月华仙体，先天七个仙窍觉醒。', 1, '2022-12-21 07:00:12', '白素媛', '13800138000',
        '11184629', '女', '202201004', 5);
INSERT INTO `s_student`
VALUES (23, 1, '2022-12-21 07:00:57', '烛龙道十三道主之一，金仙境', 1, '2022-12-21 07:00:57', '云霓', '13800138000',
        '11184629', '女', '202201005', 5);
INSERT INTO `s_student`
VALUES (24, 1, '2022-12-21 07:03:38', '妙音门左使', 1, '2022-12-21 07:03:38', '范静梅', '13800138000', '11184629', '女',
        '202201006', 5);
INSERT INTO `s_student`
VALUES (25, 1, '2022-12-21 07:04:26', '妙音门右使', 1, '2022-12-21 07:04:26', '卓如婷', '13800138000', '11184629', '女',
        '202201007', 5);
INSERT INTO `s_student`
VALUES (26, 1, '2022-12-21 07:06:20', '落云宗第一美女，人称白凤仙，身具天灵根。', 1, '2022-12-21 07:06:20', '宋玉',
        '13800138000', '11184629', '女', '202201008', 5);
INSERT INTO `s_student`
VALUES (27, 1, '2022-12-21 07:07:06', '落云宗大长老。', 1, '2022-12-21 07:07:06', '程天坤', '13800138000', '11184629',
        '男', '202201009', 5);
INSERT INTO `s_student`
VALUES (28, 1, '2022-12-21 07:07:48', '落云宗长老', 1, '2022-12-21 07:07:48', '吕洛', '13800138000', '11184629', '男',
        '202201010', 5);
INSERT INTO `s_student`
VALUES (29, 1, '2022-12-21 07:08:37', '元婴中期，苍坤上人后人', 1, '2022-12-21 07:08:37', '南陇侯', '13800138000',
        '11184629', '男', '202201011', 5);

-- ----------------------------
-- Table structure for s_student_score
-- ----------------------------
DROP TABLE IF EXISTS `s_student_score`;
CREATE TABLE `s_student_score`
(
    `id`             bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`      bigint(20) NOT NULL DEFAULT 0,
    `create_time`    datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`      bigint(20) NOT NULL DEFAULT 0,
    `update_time`    datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `score`          float                                                         NOT NULL DEFAULT 0,
    `type`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `course_id`      bigint(20) NOT NULL DEFAULT 0,
    `student_id`     bigint(20) NOT NULL DEFAULT 0,
    `grade_class_id` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX            `FKrm1fvg0dausplxq9nphxb47j9`(`course_id` ASC) USING BTREE,
    INDEX            `FKfb1pcum4u0xlbod236kmngx4f`(`student_id` ASC) USING BTREE,
    INDEX            `FK8mac4aaiivjtr0kd4kpyt4vm5`(`grade_class_id` ASC) USING BTREE,
    CONSTRAINT `s_student_score_ibfk_1` FOREIGN KEY (`grade_class_id`) REFERENCES `s_grade_class` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `s_student_score_ibfk_2` FOREIGN KEY (`student_id`) REFERENCES `s_student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `s_student_score_ibfk_3` FOREIGN KEY (`course_id`) REFERENCES `s_course` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 82 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '学生成绩表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of s_student_score
-- ----------------------------
INSERT INTO `s_student_score`
VALUES (33, 1, '2022-12-20 10:07:33', '初始成绩', 1, '2022-12-20 10:08:03', 89, '已批改', 2, 1, 1);
INSERT INTO `s_student_score`
VALUES (34, 1, '2022-12-20 10:07:33', '初始成绩', 1, '2022-12-20 10:07:56', 82, '已批改', 2, 4, 1);
INSERT INTO `s_student_score`
VALUES (35, 1, '2022-12-20 10:08:13', '初始成绩', 1, '2022-12-20 10:08:36', 100, '已批改', 1, 3, 5);
INSERT INTO `s_student_score`
VALUES (36, 1, '2022-12-20 10:08:13', '初始成绩', 1, '2022-12-20 10:08:30', 98, '已批改', 1, 5, 5);
INSERT INTO `s_student_score`
VALUES (37, 1, '2022-12-20 10:08:13', '初始成绩', 1, '2022-12-21 03:29:29', 68, '已批改', 1, 6, 5);
INSERT INTO `s_student_score`
VALUES (38, 1, '2022-12-20 10:10:05', '初始成绩', 1, '2022-12-20 10:10:18', 98, '已批改', 1, 1, 1);
INSERT INTO `s_student_score`
VALUES (39, 1, '2022-12-20 10:10:05', '初始成绩', 1, '2022-12-21 01:29:28', 59, '已批改', 1, 4, 1);
INSERT INTO `s_student_score`
VALUES (40, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:10:56', 52, '已批改', 1, 13, 1);
INSERT INTO `s_student_score`
VALUES (41, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:05', 66, '已批改', 1, 14, 1);
INSERT INTO `s_student_score`
VALUES (42, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:10', 77, '已批改', 1, 15, 1);
INSERT INTO `s_student_score`
VALUES (43, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:18', 88, '已批改', 1, 16, 1);
INSERT INTO `s_student_score`
VALUES (44, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:27', 55, '已批改', 1, 17, 1);
INSERT INTO `s_student_score`
VALUES (45, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:33', 68, '已批改', 1, 18, 1);
INSERT INTO `s_student_score`
VALUES (46, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:42', 84, '已批改', 1, 19, 1);
INSERT INTO `s_student_score`
VALUES (47, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:51', 58, '已批改', 1, 20, 1);
INSERT INTO `s_student_score`
VALUES (48, 1, '2022-12-21 07:10:47', '初始成绩', 1, '2022-12-21 07:11:59', 100, '已批改', 1, 21, 1);
INSERT INTO `s_student_score`
VALUES (49, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:12:15', 88, '已批改', 3, 1, 1);
INSERT INTO `s_student_score`
VALUES (50, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:12:07', 94, '已批改', 3, 4, 1);
INSERT INTO `s_student_score`
VALUES (51, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:12:01', 81, '已批改', 3, 13, 1);
INSERT INTO `s_student_score`
VALUES (52, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:54', 43, '已批改', 3, 14, 1);
INSERT INTO `s_student_score`
VALUES (53, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:48', 68, '已批改', 3, 15, 1);
INSERT INTO `s_student_score`
VALUES (54, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:39', 90, '已批改', 3, 16, 1);
INSERT INTO `s_student_score`
VALUES (55, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:32', 75, '已批改', 3, 17, 1);
INSERT INTO `s_student_score`
VALUES (56, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:26', 85, '已批改', 3, 18, 1);
INSERT INTO `s_student_score`
VALUES (57, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:19', 56, '已批改', 3, 19, 1);
INSERT INTO `s_student_score`
VALUES (58, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:12', 77, '已批改', 3, 20, 1);
INSERT INTO `s_student_score`
VALUES (59, 1, '2022-12-22 05:10:53', '初始成绩', 1, '2022-12-22 05:11:07', 69, '已批改', 3, 21, 1);
INSERT INTO `s_student_score`
VALUES (60, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:14:13', 99, '已批改', 3, 3, 5);
INSERT INTO `s_student_score`
VALUES (61, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:14:06', 92, '已批改', 3, 5, 5);
INSERT INTO `s_student_score`
VALUES (62, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:14:01', 88, '已批改', 3, 6, 5);
INSERT INTO `s_student_score`
VALUES (63, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:52', 71, '已批改', 3, 22, 5);
INSERT INTO `s_student_score`
VALUES (64, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:46', 64, '已批改', 3, 23, 5);
INSERT INTO `s_student_score`
VALUES (65, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:39', 60, '已批改', 3, 24, 5);
INSERT INTO `s_student_score`
VALUES (66, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:35', 59, '已批改', 3, 25, 5);
INSERT INTO `s_student_score`
VALUES (67, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:30', 58, '已批改', 3, 26, 5);
INSERT INTO `s_student_score`
VALUES (68, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:25', 67, '已批改', 3, 27, 5);
INSERT INTO `s_student_score`
VALUES (69, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:20', 69, '已批改', 3, 28, 5);
INSERT INTO `s_student_score`
VALUES (70, 1, '2022-12-22 05:13:00', '初始成绩', 1, '2022-12-22 05:13:14', 77, '已批改', 3, 29, 5);
INSERT INTO `s_student_score`
VALUES (71, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:16:17', 77, '已批改', 4, 1, 1);
INSERT INTO `s_student_score`
VALUES (72, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:16:10', 94, '已批改', 4, 4, 1);
INSERT INTO `s_student_score`
VALUES (73, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:16:03', 73, '已批改', 4, 13, 1);
INSERT INTO `s_student_score`
VALUES (74, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:56', 86, '已批改', 4, 14, 1);
INSERT INTO `s_student_score`
VALUES (75, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:51', 58, '已批改', 4, 15, 1);
INSERT INTO `s_student_score`
VALUES (76, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:46', 66, '已批改', 4, 16, 1);
INSERT INTO `s_student_score`
VALUES (77, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:42', 53, '已批改', 4, 17, 1);
INSERT INTO `s_student_score`
VALUES (78, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:37', 87, '已批改', 4, 18, 1);
INSERT INTO `s_student_score`
VALUES (79, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:33', 96, '已批改', 4, 19, 1);
INSERT INTO `s_student_score`
VALUES (80, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:27', 75, '已批改', 4, 20, 1);
INSERT INTO `s_student_score`
VALUES (81, 1, '2022-12-22 05:15:19', '初始成绩', 1, '2022-12-22 05:15:23', 44, '已批改', 4, 21, 1);

-- ----------------------------
-- Table structure for s_teacher
-- ----------------------------
DROP TABLE IF EXISTS `s_teacher`;
CREATE TABLE `s_teacher`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`   bigint(20) NOT NULL DEFAULT 0,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`   bigint(20) NOT NULL DEFAULT 0,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `phone`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `qq`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `sex`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `teach_no`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `course_id`   bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `FKgpacv4uc6gmdaridy3afaf5co`(`course_id` ASC) USING BTREE,
    CONSTRAINT `s_teacher_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `s_course` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '教师信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of s_teacher
-- ----------------------------
INSERT INTO `s_teacher`
VALUES (1, 1, '2022-12-16 12:16:53',
        '岚州枭雄，精通医术，见识过仙人交战后被仙术所震撼，从此一心想踏入修仙长生之途而隐居于七玄门，化名为墨大夫。\n', 1,
        '2022-12-16 12:16:53', '墨居仁', '13800138001', '11184629', '男', 't001', 1);
INSERT INTO `s_teacher`
VALUES (2, 1, '2022-12-16 12:18:30', '出身人界的玄剑门，自创《青元剑诀》功法', 1, '2022-12-16 12:18:30', '青元子',
        '13800138002', '11184629', '男', 't002', 2);
INSERT INTO `s_teacher`
VALUES (3, 1, '2022-12-16 12:20:36', '黄枫谷结丹长老之一', 1, '2022-12-16 12:20:36', '李化元', '13800138003',
        '11184629', '男', 't003', 3);
INSERT INTO `s_teacher`
VALUES (4, 1, '2022-12-16 12:21:57', '乱星海极阴岛岛主，姓乌，玄骨上人的徒弟。', 1, '2022-12-16 12:21:57', '极阴祖师',
        '13800138004', '11184629', '男', 't004', 4);
INSERT INTO `s_teacher`
VALUES (5, 1, '2022-12-16 12:22:56',
        '弥罗老祖是仙界真言门覆灭前最后一任宗主，修炼时间法则至大罗后期，坐下有五大亲传弟子；', 1, '2022-12-16 12:22:56',
        '弥罗老祖', '13800138005', '11184629', '男', 't005', 5);
INSERT INTO `s_teacher`
VALUES (6, 1, '2022-12-16 12:25:15',
        '本体是只黄金巨蟹，前世为魔域魔君石空解。自以为前主人未知原因陨落、自己掉落魔界魔源海苦灵岛附近百万余年，隐秘收集仙灵力修复自身。',
        1, '2022-12-16 12:25:15', '蟹道人', '13800138006', '11184629', '男', 't006', 6);
INSERT INTO `s_teacher`
VALUES (7, 1, '2022-12-16 12:26:17', '降世真仙，受上界监察仙使指派，和另一名巡查使者共同下界铲除螟虫之母。', 1,
        '2022-12-16 12:26:17', '何康老鬼', '13800138007', '11184629', '男', 't007', 7);
INSERT INTO `s_teacher`
VALUES (8, 1, '2022-12-16 12:27:26',
        '天才人物，千竹教教主，自创《大衍决》、《傀儡术》。万年前自知无望突破化神，把元神寄付傀儡之上存于世间。', 1,
        '2022-12-16 12:27:26', '大衍神君', '13800138008', '11184629', '男', 't008', 8);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`   bigint(20) NOT NULL DEFAULT 0,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`   bigint(20) NOT NULL DEFAULT 0,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `code`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role`
VALUES (1, 1, '2022-12-08 19:12:03', '系统管理员', 1, '2022-12-08 19:11:59', 'ROLE_ADMIN', '系统管理员');
INSERT INTO `sys_role`
VALUES (2, 1, '2022-12-08 19:12:15', '普通用户', 1, '2022-12-08 19:12:30', 'ROLE_USER', '普通用户');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `create_by`   bigint(20) NOT NULL DEFAULT 0,
    `create_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `remarks`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `update_by`   bigint(20) NOT NULL DEFAULT 0,
    `update_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `password`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `realname`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `sex`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `status`      int(11) NOT NULL DEFAULT 0,
    `username`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `email`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `user_icon`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
    `role_id`     bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `FK4dm5kxn3potpfgdigehw7pdyu`(`role_id` ASC) USING BTREE,
    CONSTRAINT `sys_user_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user`
VALUES (1, 1, '2022-12-02 12:55:05', '我是一名管理员', 1, '2022-12-27 03:05:28', 'E10ADC3949BA59ABBE56E057F20F883E',
        '墨鱼老师', '女', 1, 'admin', '11184629@qq.com', '27a0cd9a-581c-44bf-acba-5de5cab72b0a.png', 1);
INSERT INTO `sys_user`
VALUES (2, 1, '2022-12-05 13:06:48',
        '韩立红颜知己，名义上的妾侍。落云宗女修，负责管理药园，韩立入落云宗修炼，归其管辖，被家族逼婚，一直推脱不肯完婚。后韩立结婴后，慕沛灵对外谎称已被韩立收为妾侍，韩立为突破瓶颈，假戏真唱收其为妾侍，改修颠凤培元功，后为突破元婴期未果身陨。',
        1, '2022-12-06 03:36:52', 'E10ADC3949BA59ABBE56E057F20F883E', '慕沛灵', '女', 1, 'mupeiling',
        'mupeiling@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (3, 1, '2022-12-05 13:09:44', '韩立红颜知己。又名雪玲，本是灵界银月妖狼玲珑仙子两魂之一。', 1,
        '2022-12-06 03:30:06', 'E10ADC3949BA59ABBE56E057F20F883E', '银月', '女', 1, 'yinyue', 'yinyue@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (4, 1, '2022-12-05 13:10:43', '这是一个凡人修仙者', 1, '2022-12-05 13:10:43', 'E10ADC3949BA59ABBE56E057F20F883E',
        '韩立', '男', 1, 'hanli', '22@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (5, 1, '2022-12-05 13:14:47', '韩立红颜知己。真名汪凝，原乱星海妙音门门主之女，韩立平生所遇第一美女、红颜知己。', 1,
        '2022-12-06 03:37:29', 'E10ADC3949BA59ABBE56E057F20F883E', '紫灵', '女', 1, 'zilin', '33@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (6, 1, '2022-12-05 13:36:25', '这是一个备注', 1, '2022-12-05 13:36:25', 'E10ADC3949BA59ABBE56E057F20F883E',
        '厉飞雨', '男', 1, 'lifeiyu', '22@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (7, 1, '2022-12-05 13:38:54',
        '韩立红颜知己。魁星岛附近的修仙小门派女修，对师姐妍丽重情重义，为已成鬼魂的师姐不死重生付出极大代价，与师姐修炼阴阳轮回诀被未知通道吸入灵界，后被姜老怪/青元子收为徒。韩立于其帮衬良多，对韩立情素暗生。',
        1, '2022-12-05 13:38:54', 'E10ADC3949BA59ABBE56E057F20F883E', '元瑶', '女', 1, 'yuanyao', '33@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (10, 1, '2022-12-06 01:13:57', '韩立道侣。掩月宗女修，本命法宝朱雀环，修炼素女轮回功 。', 1, '2022-12-06 01:13:57',
        'E10ADC3949BA59ABBE56E057F20F883E', '南宫婉', '男', 1, 'nangongwan', '333@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (11, 1, '2022-12-06 03:46:19',
        '韩立黄枫谷师姐。经过一些事与非后，爱上韩立，曾经向韩立告白，被一心求大道的韩立拒绝，为等韩立未婚不嫁，道消身死。后韩立在灵界碰到与巧倩长相极其相似之人，讶然不已。',
        1, '2022-12-06 03:46:19', 'E10ADC3949BA59ABBE56E057F20F883E', '陈巧倩', '女', 1, 'chenqiaoqian',
        'chenqiaoqian@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (12, 1, '2022-12-06 03:47:02',
        '韩立旧识。为人高傲，为了驻颜修炼化春决，喜欢男人为她争风吃醋，其师傅撮合她和韩立双修，两人相互不对眼，红拂派其和韩立一起参加燕家夺宝大会。',
        1, '2022-12-06 03:47:02', 'E10ADC3949BA59ABBE56E057F20F883E', '董萱儿', '女', 1, 'dongxuaner',
        'dongxuaner@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (13, 1, '2022-12-06 03:47:40',
        '韩立旧识。御灵宗女修，亲切可人，早年曾在太南小会上卖给韩立金竺符笔，多年后和柳眉等奉命追查至木灵婴下落，被韩立发现，韩立追上她时，发现是多年前旧识只将其打晕。',
        1, '2022-12-06 03:47:40', 'E10ADC3949BA59ABBE56E057F20F883E', '菡云芝', '女', 1, 'hanyunzhi',
        'hanyunzhi@qq.com', '', 2);
INSERT INTO `sys_user`
VALUES (14, 1, '2022-12-06 03:48:41',
        '韩立旧识文樯之女。妙音门女修，外出执行危险任务，父女被韩立所救。韩立助其脱离妙音门文思月打算以身相许，奈何韩立不解风情，将她安置在一个荒岛上自行离去。',
        1, '2022-12-06 03:48:41', 'E10ADC3949BA59ABBE56E057F20F883E', '文思月', '男', 1, 'wensiyue', 'wensiyue@qq.com',
        '', 2);

SET
FOREIGN_KEY_CHECKS = 1;
