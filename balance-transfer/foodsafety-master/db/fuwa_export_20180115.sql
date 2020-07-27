-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2018-01-15 12:23:25
-- 服务器版本： 5.7.20-0ubuntu0.16.04.1
-- PHP Version: 7.0.22-0ubuntu0.16.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `fuwa`
--
CREATE DATABASE IF NOT EXISTS `fuwa` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `fuwa`;

-- --------------------------------------------------------

--
-- 表的结构 `baking_quality_inspect`
--

CREATE TABLE `baking_quality_inspect` (
  `baking_inspectid` varchar(200) NOT NULL,
  `productid` varchar(200) NOT NULL,
  `weight` varchar(200) DEFAULT NULL,
  `temperature` varchar(200) DEFAULT NULL,
  `personid` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `cooling_quality_inspect`
--

CREATE TABLE `cooling_quality_inspect` (
  `cooling_inspectid` varchar(200) NOT NULL,
  `productid` varchar(200) NOT NULL,
  `date` date DEFAULT NULL,
  `temperature` varchar(200) DEFAULT NULL,
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `drug`
--

CREATE TABLE `drug` (
  `drugid` varchar(200) NOT NULL COMMENT '药品id',
  `name` varchar(200) DEFAULT NULL COMMENT '名称',
  `dosage` varchar(200) DEFAULT NULL COMMENT '剂型',
  `standard` varchar(200) DEFAULT NULL COMMENT '生产标准',
  `effect` varchar(200) DEFAULT NULL COMMENT '作用',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `drug`
--

INSERT INTO `drug` (`drugid`, `name`, `dosage`, `standard`, `effect`, `personid`) VALUES
('CCJHB0317012501', '除虫剂', '液态', '国标', '除虫', NULL),
('FHCBJ0317021001', '氟环唑', '悬浮剂', 'Q/SSCC101-2011', '杀菌', NULL),
('KJJBJ0317020101', '咯菌腈', '悬浮剂', 'Q/320583GQB', '杀菌', NULL),
('MXABJ0317021801', '咪鲜胺', '水乳剂', 'GB22625-2008', '杀菌', NULL),
('XDJBJ0317012001', '消毒剂', '液态', '国标', '消毒', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `dry`
--

CREATE TABLE `dry` (
  `dryid` int(11) NOT NULL COMMENT '干燥id',
  `productid` varchar(200) DEFAULT NULL COMMENT '产品id',
  `temp` int(11) DEFAULT NULL COMMENT '温度',
  `drytime` int(11) DEFAULT NULL COMMENT '烘烤时间（分钟）',
  `watercontent` varchar(200) DEFAULT NULL COMMENT '水分（%）',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `dry`
--

INSERT INTO `dry` (`dryid`, `productid`, `temp`, `drytime`, `watercontent`, `date`, `personid`) VALUES
(1, 'JC1CMGUO01170828ZS01', 115, 45, '1.96', '2017-08-28', NULL),
(2, 'JC2CMGUO01170828HG01', 115, 45, '1.90', '2017-08-28', NULL),
(3, 'JC3CMS01170830ZM01', 120, 15, '2.06', '2017-08-30', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `feed`
--

CREATE TABLE `feed` (
  `feedid` varchar(200) NOT NULL COMMENT '进料id',
  `weight` varchar(200) DEFAULT NULL COMMENT '稻谷重量（kg）',
  `watercontent` varchar(200) DEFAULT NULL COMMENT '大米水分含量（%）',
  `brokenrice` varchar(200) DEFAULT NULL COMMENT '碎米率（%）',
  `qingmilv` varchar(200) DEFAULT NULL COMMENT '青米率（%）',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `feed`
--

INSERT INTO `feed` (`feedid`, `weight`, `watercontent`, `brokenrice`, `qingmilv`, `date`, `personid`) VALUES
('CK5DMJG0117082801', '169440', '14.5', '3.5', '2.5', '2017-08-28', NULL),
('CK6DMJG0117082801', '97780', '14.1', '2', '3', '2017-08-28', NULL),
('CK7DMJG0117083001', '203330', '14.1', '2', '3', '2017-08-30', NULL),
('CK8DMJG0117083001', '92060', '14.5', '3', '2', '2017-08-30', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `feed2product`
--

CREATE TABLE `feed2product` (
  `feedid` varchar(200) NOT NULL COMMENT '进料id',
  `productid` varchar(200) NOT NULL COMMENT '产品id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `feed2product`
--

INSERT INTO `feed2product` (`feedid`, `productid`, `personid`) VALUES
('CK5DMJG0117082801', 'JC1CMGUO01170828ZS01', NULL),
('CK6DMJG0117082801', 'JC2CMGUO01170828HG01', NULL),
('CK7DMJG0117083001', 'JC3CMS01170830ZM01', NULL),
('CK8DMJG0117083001', 'JC4CMGEN01170830HZ01', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `injection`
--

CREATE TABLE `injection` (
  `injectionid` int(11) NOT NULL COMMENT '注模id',
  `productid` varchar(200) DEFAULT NULL COMMENT '产品id',
  `temp` int(11) DEFAULT NULL COMMENT '温度',
  `drytime` int(11) DEFAULT NULL COMMENT '烘烤时间（秒）',
  `weight` int(11) DEFAULT NULL COMMENT '10枚重（g）',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `injection`
--

INSERT INTO `injection` (`injectionid`, `productid`, `temp`, `drytime`, `weight`, `date`, `personid`) VALUES
(1, 'JC3CMS01170830ZM01', 190, 6, 160, '2017-08-30', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `input`
--

CREATE TABLE `input` (
  `inputid` varchar(200) NOT NULL COMMENT '入库id',
  `harvestdate` date DEFAULT NULL COMMENT '收割日期',
  `quantity` int(11) DEFAULT NULL COMMENT '入库量（kg）',
  `inputdate` date DEFAULT NULL COMMENT '入库日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `input`
--

INSERT INTO `input` (`inputid`, `harvestdate`, `quantity`, `inputdate`, `personid`) VALUES
('NT12017072243990601', '2017-07-22', 439906, '2017-07-26', NULL),
('NT22017072342178101', '2017-07-23', 421781, '2017-07-27', NULL),
('NT32017072141156501', '2017-07-21', 411565, '2017-07-25', NULL),
('NT4201707202627101', '2017-07-20', 26271, '2017-07-24', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `input2warehouse`
--

CREATE TABLE `input2warehouse` (
  `inputid` varchar(200) NOT NULL COMMENT 'rukuid',
  `warehouseid` varchar(200) NOT NULL COMMENT '仓库id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `input2warehouse`
--

INSERT INTO `input2warehouse` (`inputid`, `warehouseid`, `personid`) VALUES
('NT12017072243990601', 'CK6', NULL),
('NT22017072342178101', 'CK7', NULL),
('NT22017072342178101', 'CK8', NULL),
('NT32017072141156501', 'CK6', NULL),
('NT4201707202627101', 'CK5', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `input_quality_inspect`
--

CREATE TABLE `input_quality_inspect` (
  `input_inspectid` int(11) NOT NULL COMMENT '质检id',
  `inputid` varchar(200) NOT NULL COMMENT '入库id',
  `mildew` varchar(200) DEFAULT NULL COMMENT '霉变',
  `pest` varchar(200) DEFAULT NULL COMMENT '虫害',
  `watercontent` varchar(200) DEFAULT NULL COMMENT '水分含量（%）',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `input_quality_inspect`
--

INSERT INTO `input_quality_inspect` (`input_inspectid`, `inputid`, `mildew`, `pest`, `watercontent`, `personid`) VALUES
(1, 'NT4201707202627101', '无', '无', '13.1', NULL),
(2, 'NT32017072141156501', '无', '无', '13', NULL),
(3, 'NT12017072243990601', '无', '无', '13.1', NULL),
(4, 'NT22017072342178101', '无', '无', '13.1', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `material`
--

CREATE TABLE `material` (
  `materialid` varchar(200) NOT NULL COMMENT '原料id',
  `kind` varchar(200) DEFAULT NULL COMMENT '种类',
  `weight` varchar(200) DEFAULT NULL COMMENT '重量（t）',
  `source` varchar(200) DEFAULT NULL COMMENT '来源',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `material`
--

INSERT INTO `material` (`materialid`, `kind`, `weight`, `source`, `date`, `personid`) VALUES
('BSTMS0517082201', '白砂糖', '10', '四川眉山王五糖业', '2017-08-22', NULL),
('XMFSJZ0617082501', '小麦粉', '10', '河北石家庄晋务粉业', '2017-08-25', NULL),
('ZLYMAS0417082001', '棕榈油', '10', '安徽马鞍山金龙鱼有限公司', '2017-08-20', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `material2product`
--

CREATE TABLE `material2product` (
  `materialid` varchar(200) NOT NULL COMMENT '原料id',
  `productid` varchar(200) NOT NULL COMMENT '产品id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `material2product`
--

INSERT INTO `material2product` (`materialid`, `productid`, `personid`) VALUES
('ZLYMAS0417082001', 'JC1CMGUO01170828ZS01', NULL),
('XMFSJZ0617082501', 'JC2CMGUO01170828HG01', NULL),
('BSTMS0517082201', 'JC3CMS01170830ZM01', NULL),
('BSTMS0517082201', 'JC4CMGEN01170830HZ01', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `package_quality_inspect`
--

CREATE TABLE `package_quality_inspect` (
  `package_inspectid` varchar(200) NOT NULL,
  `productid` varchar(200) NOT NULL,
  `oxygencontent` varchar(200) DEFAULT NULL,
  `innercapsule` varchar(200) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `person`
--

CREATE TABLE `person` (
  `personid` varchar(200) NOT NULL COMMENT '用户id',
  `name` varchar(200) DEFAULT NULL COMMENT '姓名',
  `sex` varchar(200) DEFAULT NULL COMMENT '性别',
  `workplace` varchar(200) DEFAULT NULL COMMENT '工作单位',
  `job` varchar(200) DEFAULT NULL COMMENT '岗位',
  `password` varchar(200) DEFAULT NULL COMMENT '密码'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `person`
--

INSERT INTO `person` (`personid`, `name`, `sex`, `workplace`, `job`, `password`) VALUES
('FW00014', '齐永新', '男', '福天下合作社', '技术员', NULL),
('FW00264', '夏良军', '男', '米业生产', '监仓', NULL),
('FW006133', '徐炎浩', '男', '福天下合作社', '基地管理员', NULL),
('FW006222', '张道海', '男', '福天下合作社', '稻虾基地管理员', NULL),
('FW006232', '杜宝元', '男', '福天下合作社', '稻虾基地管理员', NULL),
('FW006281', '张从仿', '男', '福天下合作社', '基地管理员', NULL),
('FW00630', '谢申华', '男', '福天下合作社', '工程部经理', NULL),
('FW00656', '谢星', '男', '米业生产', '主检', NULL),
('FW00702', '谢均涛', '男', '米业生产', '原料保管', NULL),
('FW00803', '谢松才', '男', '福天下合作社', '基地管理员', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `plant`
--

CREATE TABLE `plant` (
  `plantid` varchar(200) NOT NULL COMMENT '种植id',
  `place` varchar(200) DEFAULT NULL COMMENT '种植地点',
  `startdate` date DEFAULT NULL COMMENT '种植起始日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `plant`
--

INSERT INTO `plant` (`plantid`, `place`, `startdate`, `personid`) VALUES
('NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', '1号农田', '2017-04-04', NULL),
('NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501', '2号农田', '2017-04-05', NULL),
('NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301', '3号农田', '2017-04-03', NULL),
('NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201', '4号农田', '2017-04-05', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `plant2input`
--

CREATE TABLE `plant2input` (
  `plantid` varchar(200) NOT NULL COMMENT '种植id',
  `inputid` varchar(200) NOT NULL COMMENT '入库id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `plant2input`
--

INSERT INTO `plant2input` (`plantid`, `inputid`, `personid`) VALUES
('NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', 'NT12017072243990601', NULL),
('NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501', 'NT22017072342178101', NULL),
('NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301', 'NT32017072141156501', NULL),
('NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201', 'NT4201707202627101', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `plantusedrug`
--

CREATE TABLE `plantusedrug` (
  `plantid` varchar(200) NOT NULL COMMENT '种植id',
  `drugid` varchar(200) NOT NULL COMMENT '药品id',
  `dosage` varchar(200) DEFAULT NULL COMMENT '浓度（ml/L）',
  `effect` varchar(200) DEFAULT NULL COMMENT '作用',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `plantusedrug`
--

INSERT INTO `plantusedrug` (`plantid`, `drugid`, `dosage`, `effect`, `date`, `personid`) VALUES
('NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', 'CCJHB0317012501', '9', '防治稻瘟、叶枯、纹枯', '2017-06-18', NULL),
('NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501', 'CCJHB0317012501', '8.7', '防治稻瘟、叶枯、纹枯', '2017-06-18', NULL),
('NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301', 'CCJHB0317012501', '9', '防治稻瘟、叶枯、纹枯', '2017-06-18', NULL),
('NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201', 'CCJHB0317012501', '8.5', '防治稻瘟、叶枯、纹枯', '2017-06-18', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `product`
--

CREATE TABLE `product` (
  `productid` varchar(200) NOT NULL COMMENT '产品id',
  `name` varchar(200) DEFAULT NULL COMMENT '名称',
  `specification` varchar(200) DEFAULT NULL COMMENT '口味',
  `flavor` varchar(200) DEFAULT NULL COMMENT '规格',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `product`
--

INSERT INTO `product` (`productid`, `name`, `specification`, `flavor`, `date`, `personid`) VALUES
('JC1CMGUO01170828ZS01', '糙米果', '芝士味', '2kg*1', '2017-08-28', NULL),
('JC2CMGUO01170828HG01', '糙米果', '黄瓜味', '2kg*1', '2017-08-28', NULL),
('JC3CMS01170830ZM01', '糙米酥', '芝麻味', '32g*12*6', '2017-08-30', NULL),
('JC4CMGEN01170830HZ01', '糙米羹', '红枣味', '360g*12', '2017-08-30', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `puffed`
--

CREATE TABLE `puffed` (
  `puffedid` int(11) NOT NULL COMMENT '膨化id',
  `productid` varchar(200) DEFAULT NULL COMMENT '产品id',
  `mach_temp` int(11) DEFAULT NULL COMMENT '主机温度',
  `feed_speed` int(11) DEFAULT NULL COMMENT '下料速度（Hz）',
  `mach_speed` int(11) DEFAULT NULL COMMENT '主机转速（Hz）',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `puffed`
--

INSERT INTO `puffed` (`puffedid`, `productid`, `mach_temp`, `feed_speed`, `mach_speed`, `date`, `personid`) VALUES
(1, 'JC1CMGUO01170828ZS01', 155, 18, 39, '2017-08-28', NULL),
(2, 'JC2CMGUO01170828HG01', 160, 19, 39, '2017-08-28', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seed`
--

CREATE TABLE `seed` (
  `seedid` varchar(200) NOT NULL COMMENT '种子ID',
  `variety` varchar(200) DEFAULT NULL COMMENT '品种',
  `type` varchar(200) DEFAULT NULL COMMENT '类型',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seed`
--

INSERT INTO `seed` (`seedid`, `variety`, `type`, `personid`) VALUES
('JZ361JX0117020101', '江早361', '早稻', NULL),
('TLY83HN0117020101', '潭两优83', '早稻', NULL),
('ZJZ17ZJJ0217020101', '中嘉早17', '早稻', NULL),
('ZLY171ZJ0217020101', '株两优171', '早稻', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seed2seedling`
--

CREATE TABLE `seed2seedling` (
  `seedid` varchar(200) NOT NULL COMMENT '种子id',
  `seedlingid` varchar(200) NOT NULL COMMENT '秧苗id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seed2seedling`
--

INSERT INTO `seed2seedling` (`seedid`, `seedlingid`, `personid`) VALUES
('JZ361JX0117020101', 'JZ361JX0117020101XDJBJ031701200117031401DP217031601', NULL),
('TLY83HN0117020101', 'TLY83HN0117020101XDJBJ031701200117031701HY17031901', NULL),
('ZJZ17ZJJ0217020101', 'ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701', NULL),
('ZLY171ZJ0217020101', 'ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seedling`
--

CREATE TABLE `seedling` (
  `seedlingid` varchar(200) NOT NULL COMMENT '秧苗id',
  `nurseryplace` varchar(200) DEFAULT NULL COMMENT '育秧地点',
  `startdate` date DEFAULT NULL COMMENT '育秧起始日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seedling`
--

INSERT INTO `seedling` (`seedlingid`, `nurseryplace`, `startdate`, `personid`) VALUES
('JZ361JX0117020101XDJBJ031701200117031401DP217031601', '2号大棚', '2017-03-16', NULL),
('TLY83HN0117020101XDJBJ031701200117031701HY17031901', '红阳育秧工厂', '2017-03-19', NULL),
('ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701', '1号大棚', '2017-03-17', NULL),
('ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801', '红阳育秧工厂', '2017-03-18', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seedling2plant`
--

CREATE TABLE `seedling2plant` (
  `seedlingid` varchar(200) NOT NULL COMMENT '秧苗id',
  `plantid` varchar(200) NOT NULL COMMENT '种植id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seedling2plant`
--

INSERT INTO `seedling2plant` (`seedlingid`, `plantid`, `personid`) VALUES
('ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801', 'NT1ZLY171ZJ0217020101XDJBJ031701200117031601HY1703180117040401', NULL),
('TLY83HN0117020101XDJBJ031701200117031701HY17031901', 'NT2TLY83HN0117020101XDJBJ031701200117031701HY1703190117040501', NULL),
('ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701', 'NT3ZJZ17ZJJL0217020101XDJBJ031701200117031501DP11703170117040301', NULL),
('JZ361JX0117020101XDJBJ031701200117031401DP217031601', 'NT4JZ361JX0117020101XDJBJ031701200117031401DP21703160117040201', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seedling_quality_inspect`
--

CREATE TABLE `seedling_quality_inspect` (
  `seedling_inspectid` int(11) NOT NULL COMMENT '质检id',
  `seedlingid` varchar(200) NOT NULL COMMENT '秧苗id',
  `density` varchar(200) DEFAULT NULL COMMENT '密度（kg/㎡）',
  `pest` varchar(200) DEFAULT NULL COMMENT '虫害',
  `date` date DEFAULT NULL COMMENT '日期',
  `growth` varchar(200) DEFAULT NULL,
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seedling_quality_inspect`
--

INSERT INTO `seedling_quality_inspect` (`seedling_inspectid`, `seedlingid`, `density`, `pest`, `date`, `growth`, `personid`) VALUES
(1, 'JZ361JX0117020101XDJBJ031701200117031401DP217031601', '0.344', '无', '2017-03-26', NULL, NULL),
(2, 'ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701', '0.356', '无', '2017-03-26', NULL, NULL),
(3, 'ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801', '0.360', '无', '2017-03-28', NULL, NULL),
(4, 'TLY83HN0117020101XDJBJ031701200117031701HY17031901', '0.350', '无', '2017-03-28', NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seedlingspraydrug`
--

CREATE TABLE `seedlingspraydrug` (
  `seedlingid` varchar(200) NOT NULL COMMENT '秧苗id',
  `drugid` varchar(200) NOT NULL COMMENT '药品id',
  `dosage` varchar(200) DEFAULT NULL COMMENT '用量(ml/kg)',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seedlingspraydrug`
--

INSERT INTO `seedlingspraydrug` (`seedlingid`, `drugid`, `dosage`, `date`, `personid`) VALUES
('JZ361JX0117020101XDJBJ031701200117031401DP217031601', 'KJJBJ0317020101', '20.5', '2017-03-16', NULL),
('TLY83HN0117020101XDJBJ031701200117031701HY17031901', 'KJJBJ0317020101', '21', '2017-03-19', NULL),
('ZJZ17ZJJL0217020101XDJBJ031701200117031501DP117031701', 'KJJBJ0317020101', '22', '2017-03-17', NULL),
('ZLY171ZJ0217020101XDJBJ031701200117031601HY17031801', 'KJJBJ0317020101', '23', '2017-03-18', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `seedsoakdrug`
--

CREATE TABLE `seedsoakdrug` (
  `seedid` varchar(200) NOT NULL COMMENT '种子id',
  `drugid` varchar(200) NOT NULL COMMENT '药品id',
  `concentration` varchar(200) DEFAULT NULL COMMENT '浓度（ml/L）',
  `startdate` date DEFAULT NULL COMMENT '起始日期',
  `enddata` date DEFAULT NULL COMMENT '结束日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `seedsoakdrug`
--

INSERT INTO `seedsoakdrug` (`seedid`, `drugid`, `concentration`, `startdate`, `enddata`, `personid`) VALUES
('JZ361JX0117020101', 'XDJBJ0317012001', '29.5', '2017-03-14', '2017-03-15', NULL),
('TLY83HN0117020101', 'XDJBJ0317012001', '32', '2017-03-17', '2017-03-18', NULL),
('ZJZ17ZJj0217020101', 'XDJBJ0317012001', '30', '2017-03-15', '2017-03-16', NULL),
('ZLY171ZJ0217020101', 'XDJBJ0317012001', '33', '2017-03-16', '2017-03-17', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `warehouse`
--

CREATE TABLE `warehouse` (
  `warehouseid` varchar(200) NOT NULL COMMENT '仓库id',
  `place` varchar(200) DEFAULT NULL COMMENT '地点',
  `capacity` varchar(200) DEFAULT NULL COMMENT '容量（t）',
  `standard` varchar(200) DEFAULT NULL COMMENT '标准',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `warehouse`
--

INSERT INTO `warehouse` (`warehouseid`, `place`, `capacity`, `standard`, `personid`) VALUES
('CK5', '湖北监利新沟银欣米业厂', '9676', '标准', NULL),
('CK6', '湖北监利新沟银欣米业厂', '9370', '标准', NULL),
('CK7', '湖北监利新沟银欣米业厂', '8869', '标准', NULL),
('CK8', '湖北监利新沟银欣米业厂', '9641', '标准', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `warehouse2feed`
--

CREATE TABLE `warehouse2feed` (
  `warehouseid` varchar(200) NOT NULL COMMENT '仓库id',
  `feedid` varchar(200) NOT NULL COMMENT '进料id',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `warehouse2feed`
--

INSERT INTO `warehouse2feed` (`warehouseid`, `feedid`, `personid`) VALUES
('CK5', 'CK5DMJG0117082801', NULL),
('CK6', 'CK6DMJG0117082801', NULL),
('CK7', 'CK7DMJG0117083001', NULL),
('CK8', 'CK8DMJG0117083001', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `warehouse_inspect`
--

CREATE TABLE `warehouse_inspect` (
  `houseinspectid` int(11) NOT NULL COMMENT '质检id',
  `warehouseid` varchar(200) NOT NULL COMMENT '仓库id',
  `watercontent` varchar(200) DEFAULT NULL COMMENT '稻谷水分（%）',
  `pest` varchar(200) DEFAULT NULL COMMENT '害虫',
  `housetemp` varchar(200) DEFAULT NULL COMMENT '仓库温度',
  `housemoisture` varchar(200) DEFAULT NULL COMMENT '仓库相对湿度（%）',
  `graintemp` varchar(200) DEFAULT NULL COMMENT '粮堆温度',
  `date` date DEFAULT NULL COMMENT '日期',
  `personid` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `warehouse_inspect`
--

INSERT INTO `warehouse_inspect` (`houseinspectid`, `warehouseid`, `watercontent`, `pest`, `housetemp`, `housemoisture`, `graintemp`, `date`, `personid`) VALUES
(1, 'CK5', '13.5', '无', '26.5', '70', '31', '2017-08-27', NULL),
(2, 'CK6', '13.5', '无', '26.5', '70', '31', '2017-08-27', NULL),
(3, 'CK7', '13.4', '无', '25', '71', '30', '2017-08-29', NULL),
(4, 'CK8', '13.4', '无', '225', '71', '30', '2017-08-29', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `baking_quality_inspect`
--
ALTER TABLE `baking_quality_inspect`
  ADD PRIMARY KEY (`baking_inspectid`),
  ADD KEY `fk_baking_quality_inspect_product` (`productid`);

--
-- Indexes for table `cooling_quality_inspect`
--
ALTER TABLE `cooling_quality_inspect`
  ADD PRIMARY KEY (`cooling_inspectid`),
  ADD KEY `fk_cooling_quality_inspect_product` (`productid`);

--
-- Indexes for table `drug`
--
ALTER TABLE `drug`
  ADD PRIMARY KEY (`drugid`),
  ADD KEY `fk_drug_person` (`personid`);

--
-- Indexes for table `dry`
--
ALTER TABLE `dry`
  ADD PRIMARY KEY (`dryid`),
  ADD KEY `fk_dry_person` (`personid`),
  ADD KEY `fk_dry_product` (`productid`);

--
-- Indexes for table `feed`
--
ALTER TABLE `feed`
  ADD PRIMARY KEY (`feedid`),
  ADD KEY `fk_feed_person` (`personid`);

--
-- Indexes for table `feed2product`
--
ALTER TABLE `feed2product`
  ADD PRIMARY KEY (`productid`,`feedid`),
  ADD KEY `fk_feed2product_person` (`personid`),
  ADD KEY `fk_feed2product_feed` (`feedid`);

--
-- Indexes for table `injection`
--
ALTER TABLE `injection`
  ADD PRIMARY KEY (`injectionid`),
  ADD KEY `fk_injection_person` (`personid`),
  ADD KEY `fk_injection_product` (`productid`);

--
-- Indexes for table `input`
--
ALTER TABLE `input`
  ADD PRIMARY KEY (`inputid`),
  ADD KEY `fk_input_person` (`personid`);

--
-- Indexes for table `input2warehouse`
--
ALTER TABLE `input2warehouse`
  ADD PRIMARY KEY (`inputid`,`warehouseid`),
  ADD KEY `fk_input2warehouse_person` (`personid`),
  ADD KEY `fk_input2warehouse_warehouse` (`warehouseid`);

--
-- Indexes for table `input_quality_inspect`
--
ALTER TABLE `input_quality_inspect`
  ADD PRIMARY KEY (`input_inspectid`),
  ADD KEY `fk_input_quality_inspect_input` (`inputid`);

--
-- Indexes for table `material`
--
ALTER TABLE `material`
  ADD PRIMARY KEY (`materialid`),
  ADD KEY `fk_material_person` (`personid`);

--
-- Indexes for table `material2product`
--
ALTER TABLE `material2product`
  ADD PRIMARY KEY (`productid`,`materialid`),
  ADD KEY `fk_material2product_person` (`personid`),
  ADD KEY `fk_material2product_material` (`materialid`);

--
-- Indexes for table `package_quality_inspect`
--
ALTER TABLE `package_quality_inspect`
  ADD PRIMARY KEY (`package_inspectid`),
  ADD KEY `fk_package_quality_inspect_product` (`productid`);

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`personid`);

--
-- Indexes for table `plant`
--
ALTER TABLE `plant`
  ADD PRIMARY KEY (`plantid`),
  ADD KEY `fk_plant_person` (`personid`);

--
-- Indexes for table `plant2input`
--
ALTER TABLE `plant2input`
  ADD PRIMARY KEY (`plantid`,`inputid`),
  ADD KEY `fk_plant2input_person` (`personid`),
  ADD KEY `fk_plant2input_input` (`inputid`);

--
-- Indexes for table `plantusedrug`
--
ALTER TABLE `plantusedrug`
  ADD PRIMARY KEY (`plantid`,`drugid`),
  ADD KEY `fk_plantusedrug_person` (`personid`),
  ADD KEY `fk_plantusedrug_drug` (`drugid`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`productid`),
  ADD KEY `fk_product_person` (`personid`);

--
-- Indexes for table `puffed`
--
ALTER TABLE `puffed`
  ADD PRIMARY KEY (`puffedid`),
  ADD KEY `fk_puffed_person` (`personid`),
  ADD KEY `fk_puffed_product` (`productid`);

--
-- Indexes for table `seed`
--
ALTER TABLE `seed`
  ADD PRIMARY KEY (`seedid`),
  ADD KEY `fk_seed_person` (`personid`);

--
-- Indexes for table `seed2seedling`
--
ALTER TABLE `seed2seedling`
  ADD PRIMARY KEY (`seedlingid`,`seedid`),
  ADD KEY `fk_seed2seedling_person` (`personid`),
  ADD KEY `fk_seed2seedling_seed` (`seedid`);

--
-- Indexes for table `seedling`
--
ALTER TABLE `seedling`
  ADD PRIMARY KEY (`seedlingid`),
  ADD KEY `fk_seedling_person` (`personid`);

--
-- Indexes for table `seedling2plant`
--
ALTER TABLE `seedling2plant`
  ADD PRIMARY KEY (`plantid`,`seedlingid`),
  ADD KEY `fk_seedling2plant_person` (`personid`),
  ADD KEY `fk_seedling2plant_seedling` (`seedlingid`);

--
-- Indexes for table `seedling_quality_inspect`
--
ALTER TABLE `seedling_quality_inspect`
  ADD PRIMARY KEY (`seedling_inspectid`),
  ADD KEY `fk_seedling_quality_inspect_seedling` (`seedlingid`);

--
-- Indexes for table `seedlingspraydrug`
--
ALTER TABLE `seedlingspraydrug`
  ADD PRIMARY KEY (`drugid`,`seedlingid`),
  ADD KEY `fk_seedlingspraydrug_person` (`personid`),
  ADD KEY `fk_seedlingspraydrug_seed` (`seedlingid`);

--
-- Indexes for table `seedsoakdrug`
--
ALTER TABLE `seedsoakdrug`
  ADD PRIMARY KEY (`seedid`,`drugid`),
  ADD KEY `fk_seedsoakdrug_person` (`personid`),
  ADD KEY `fk_seedsoakdrug_drug` (`drugid`);

--
-- Indexes for table `warehouse`
--
ALTER TABLE `warehouse`
  ADD PRIMARY KEY (`warehouseid`),
  ADD KEY `fk_warehouse_person` (`personid`);

--
-- Indexes for table `warehouse2feed`
--
ALTER TABLE `warehouse2feed`
  ADD PRIMARY KEY (`warehouseid`,`feedid`),
  ADD KEY `fk_warehousetofeed_person` (`personid`),
  ADD KEY `fk_warehousetofeed_feed` (`feedid`);

--
-- Indexes for table `warehouse_inspect`
--
ALTER TABLE `warehouse_inspect`
  ADD PRIMARY KEY (`houseinspectid`),
  ADD KEY `fk_warehouse_inspect_warehouse` (`warehouseid`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `dry`
--
ALTER TABLE `dry`
  MODIFY `dryid` int(11) NOT NULL AUTO_INCREMENT COMMENT '干燥id', AUTO_INCREMENT=4;
--
-- 使用表AUTO_INCREMENT `injection`
--
ALTER TABLE `injection`
  MODIFY `injectionid` int(11) NOT NULL AUTO_INCREMENT COMMENT '注模id', AUTO_INCREMENT=2;
--
-- 使用表AUTO_INCREMENT `input_quality_inspect`
--
ALTER TABLE `input_quality_inspect`
  MODIFY `input_inspectid` int(11) NOT NULL AUTO_INCREMENT COMMENT '质检id', AUTO_INCREMENT=5;
--
-- 使用表AUTO_INCREMENT `seedling_quality_inspect`
--
ALTER TABLE `seedling_quality_inspect`
  MODIFY `seedling_inspectid` int(11) NOT NULL AUTO_INCREMENT COMMENT '质检id', AUTO_INCREMENT=5;
--
-- 使用表AUTO_INCREMENT `warehouse_inspect`
--
ALTER TABLE `warehouse_inspect`
  MODIFY `houseinspectid` int(11) NOT NULL AUTO_INCREMENT COMMENT '质检id', AUTO_INCREMENT=5;
--
-- 限制导出的表
--

--
-- 限制表 `baking_quality_inspect`
--
ALTER TABLE `baking_quality_inspect`
  ADD CONSTRAINT `fk_baking_quality_inspect_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `cooling_quality_inspect`
--
ALTER TABLE `cooling_quality_inspect`
  ADD CONSTRAINT `fk_cooling_quality_inspect_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `drug`
--
ALTER TABLE `drug`
  ADD CONSTRAINT `fk_drug_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `dry`
--
ALTER TABLE `dry`
  ADD CONSTRAINT `fk_dry_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_dry_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `feed`
--
ALTER TABLE `feed`
  ADD CONSTRAINT `fk_feed_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `feed2product`
--
ALTER TABLE `feed2product`
  ADD CONSTRAINT `fk_feed2product_feed` FOREIGN KEY (`feedid`) REFERENCES `feed` (`feedid`),
  ADD CONSTRAINT `fk_feed2product_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_feed2product_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `injection`
--
ALTER TABLE `injection`
  ADD CONSTRAINT `fk_injection_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_injection_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `input`
--
ALTER TABLE `input`
  ADD CONSTRAINT `fk_input_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `input2warehouse`
--
ALTER TABLE `input2warehouse`
  ADD CONSTRAINT `fk_input2warehouse_input` FOREIGN KEY (`inputid`) REFERENCES `input` (`inputid`),
  ADD CONSTRAINT `fk_input2warehouse_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_input2warehouse_warehouse` FOREIGN KEY (`warehouseid`) REFERENCES `warehouse` (`warehouseid`);

--
-- 限制表 `input_quality_inspect`
--
ALTER TABLE `input_quality_inspect`
  ADD CONSTRAINT `fk_input_quality_inspect_input` FOREIGN KEY (`inputid`) REFERENCES `input` (`inputid`);

--
-- 限制表 `material`
--
ALTER TABLE `material`
  ADD CONSTRAINT `fk_material_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `material2product`
--
ALTER TABLE `material2product`
  ADD CONSTRAINT `fk_material2product_material` FOREIGN KEY (`materialid`) REFERENCES `material` (`materialid`),
  ADD CONSTRAINT `fk_material2product_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_material2product_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `package_quality_inspect`
--
ALTER TABLE `package_quality_inspect`
  ADD CONSTRAINT `fk_package_quality_inspect_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `plant`
--
ALTER TABLE `plant`
  ADD CONSTRAINT `fk_plant_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `plant2input`
--
ALTER TABLE `plant2input`
  ADD CONSTRAINT `fk_plant2input_input` FOREIGN KEY (`inputid`) REFERENCES `input` (`inputid`),
  ADD CONSTRAINT `fk_plant2input_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_plant2input_plant` FOREIGN KEY (`plantid`) REFERENCES `plant` (`plantid`);

--
-- 限制表 `plantusedrug`
--
ALTER TABLE `plantusedrug`
  ADD CONSTRAINT `fk_plantusedrug_drug` FOREIGN KEY (`drugid`) REFERENCES `drug` (`drugid`),
  ADD CONSTRAINT `fk_plantusedrug_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_plantusedrug_plant` FOREIGN KEY (`plantid`) REFERENCES `plant` (`plantid`);

--
-- 限制表 `product`
--
ALTER TABLE `product`
  ADD CONSTRAINT `fk_product_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `puffed`
--
ALTER TABLE `puffed`
  ADD CONSTRAINT `fk_puffed_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_puffed_product` FOREIGN KEY (`productid`) REFERENCES `product` (`productid`);

--
-- 限制表 `seed`
--
ALTER TABLE `seed`
  ADD CONSTRAINT `fk_seed_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `seed2seedling`
--
ALTER TABLE `seed2seedling`
  ADD CONSTRAINT `fk_seed2seedling_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_seed2seedling_seed` FOREIGN KEY (`seedid`) REFERENCES `seed` (`seedid`),
  ADD CONSTRAINT `fk_seed2seedling_seedling` FOREIGN KEY (`seedlingid`) REFERENCES `seedling` (`seedlingid`);

--
-- 限制表 `seedling`
--
ALTER TABLE `seedling`
  ADD CONSTRAINT `fk_seedling_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `seedling2plant`
--
ALTER TABLE `seedling2plant`
  ADD CONSTRAINT `fk_seedling2plant_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_seedling2plant_plant` FOREIGN KEY (`plantid`) REFERENCES `plant` (`plantid`),
  ADD CONSTRAINT `fk_seedling2plant_seedling` FOREIGN KEY (`seedlingid`) REFERENCES `seedling` (`seedlingid`);

--
-- 限制表 `seedling_quality_inspect`
--
ALTER TABLE `seedling_quality_inspect`
  ADD CONSTRAINT `fk_seedling_quality_inspect_seedling` FOREIGN KEY (`seedlingid`) REFERENCES `seedling` (`seedlingid`);

--
-- 限制表 `seedlingspraydrug`
--
ALTER TABLE `seedlingspraydrug`
  ADD CONSTRAINT `fk_seedlingspraydrug_drug` FOREIGN KEY (`drugid`) REFERENCES `drug` (`drugid`),
  ADD CONSTRAINT `fk_seedlingspraydrug_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_seedlingspraydrug_seed` FOREIGN KEY (`seedlingid`) REFERENCES `seedling` (`seedlingid`);

--
-- 限制表 `seedsoakdrug`
--
ALTER TABLE `seedsoakdrug`
  ADD CONSTRAINT `fk_seedsoakdrug_drug` FOREIGN KEY (`drugid`) REFERENCES `drug` (`drugid`),
  ADD CONSTRAINT `fk_seedsoakdrug_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_seedsoakdrug_seed` FOREIGN KEY (`seedid`) REFERENCES `seed` (`seedid`);

--
-- 限制表 `warehouse`
--
ALTER TABLE `warehouse`
  ADD CONSTRAINT `fk_warehouse_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`);

--
-- 限制表 `warehouse2feed`
--
ALTER TABLE `warehouse2feed`
  ADD CONSTRAINT `fk_warehousetofeed_feed` FOREIGN KEY (`feedid`) REFERENCES `feed` (`feedid`),
  ADD CONSTRAINT `fk_warehousetofeed_person` FOREIGN KEY (`personid`) REFERENCES `person` (`personid`),
  ADD CONSTRAINT `fk_warehousetofeed_warehouse` FOREIGN KEY (`warehouseid`) REFERENCES `warehouse` (`warehouseid`);

--
-- 限制表 `warehouse_inspect`
--
ALTER TABLE `warehouse_inspect`
  ADD CONSTRAINT `fk_warehouse_inspect_warehouse` FOREIGN KEY (`warehouseid`) REFERENCES `warehouse` (`warehouseid`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
