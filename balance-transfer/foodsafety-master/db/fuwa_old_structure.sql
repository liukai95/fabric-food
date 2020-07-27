/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017/8/22 2:22:16                            */
/*==============================================================*/

drop table if exists person;

-- 5
drop table if exists seed;
drop table if exists seedling;
drop table if exists seed2seedling;

drop table if exists plant;
drop table if exists seedling2plant;

drop table if exists input;
drop table if exists plant2input;
-- 3
drop table if exists warehouse;
drop table if exists input2warehouse;

drop table if exists feed;
drop table if exists warehousetofeed;

drop table if exists product;
drop table if exists feed2product;
drop table if exists material;
drop table if exists material2product;
-- 4

-- 4
drop table if exists drug;
drop table if exists plantusedrug;
drop table if exists seedsoakdrug;
drop table if exists seedlingspraydrug;


drop table if exists dry;
drop table if exists puffed;
drop table if exists injection;


-- 6
drop table if exists baking_quality_inspect;
drop table if exists cooling_quality_inspect;
drop table if exists package_quality_inspect;
drop table if exists seedling_quality_inspect;
drop table if exists input_quality_inspect;
drop table if exists warehouse_inspect;



create table person (
   personid             varchar(200) not null,  -- 用户id
   name                 varchar(200),           -- 姓名
   sex                  varchar(200),           -- 性别
   workplace            varchar(200),           -- 工作单位
   job                  varchar(200),           -- 岗位
   password             varchar(200),           -- 密码
   primary key (personid)
);

/* Table: seed                           7                       */
create table seed (
   seedid               varchar(200) not null,
   variety              varchar(200),
   type                 varchar(200),
   personid             varchar(200),
   primary key (seedid),
   constraint fk_seed_person foreign key (personid) references person(personid)
);

/* Table: seedling                          9                    */
create table seedling (
   seedlingid           varchar(200) not null, -- 秧苗id
   nurseryplace         varchar(200), -- 育秧地点
   startdate            date,         -- 育秧起始日期
   personid             varchar(200),
   primary key (seedlingid),
   constraint fk_seedling_person foreign key (personid) references person(personid)
);

/* Table: seed2seedling                      8                          */
create table seed2seedling (
   seedid               varchar(200) not null, -- 种子id
   seedlingid           varchar(200) not null, -- 秧苗id
   personid             varchar(200),
   primary key (seedlingid, seedid),
   constraint fk_seed2seedling_person  foreign key (personid)  references person(personid),
   constraint fk_seed2seedling_seed    foreign key (seedid)    references seed(seedid),
   constraint fk_seed2seedling_seedling foreign key (seedlingid) references seedling(seedlingid)
);


/* Table: plant                              11                   */
create table plant (
   plantid              varchar(200) not null, -- 种植id
   place                varchar(200), -- 地点
   startdate            date,         -- 种植起始日期d         
   personid             varchar(200),
   primary key (plantid),
   constraint fk_plant_person foreign key (personid) references person(personid)
);

/* Table: seedling2plant                      10                   */
create table seedling2plant (
   seedlingid           varchar(200) not null, -- 秧苗id
   plantid              varchar(200) not null, -- 种植id
   personid             varchar(200),
   primary key (plantid, seedlingid),
   constraint fk_seedling2plant_person   foreign key (personid)   references person(personid),
   constraint fk_seedling2plant_seedling  foreign key (seedlingid)  references seedling(seedlingid),
   constraint fk_seedling2plant_plant    foreign key (plantid)    references plant(plantid)
);

/*  */
create table input (
   inputid              varchar(200) not null, -- 入库id
   harvestdate          date,                  -- 收割日期
   quantity             int,                   -- 入库量
   inputdate            date,                  -- 入库日期
   personid             varchar(200),
   primary key (inputid),
   constraint fk_input_person foreign key (personid) references person(personid)
);

/* Table: harvest                           12                   */
create table plant2input (
   plantid              varchar(200) not null, -- 种植id
   inputid              varchar(200) not null, -- 入库id
   personid             varchar(200),
   primary key (plantid, inputid),
   constraint fk_plant2input_person   foreign key (personid)   references person(personid),
   constraint fk_plant2input_plant    foreign key (plantid)    references plant(plantid),
   constraint fk_plant2input_input    foreign key (inputid)    references input(inputid)
);

/* Table: warehouse                        14                     */
create table warehouse (
   warehouseid          varchar(200) not null, -- 仓库id
   place                varchar(200), -- 地点
   capacity             varchar(200), -- 容量
   standard             varchar(200), -- 标准
   personid             varchar(200),
   primary key (warehouseid),
   constraint fk_warehouse_person foreign key (personid) references person(personid)
);

/* Table: input                           13                    */
create table input2warehouse (
   inputid              varchar(200) not null, -- 入库id
   warehouseid          varchar(200) not null, -- 仓库id
   personid             varchar(200),
   primary key (inputid, warehouseid),
   constraint fk_input2warehouse_person    foreign key (personid) references person(personid),
   constraint fk_input2warehouse_input    foreign key (inputid)    references input(inputid),
   constraint fk_input2warehouse_warehouse foreign key (warehouseid) references warehouse(warehouseid)
);

/* Table: feed                             15                     */
create table feed (
   feedid               varchar(200) not null, -- 进料id
   weight               varchar(200), -- 稻谷重量
   watercontent         varchar(200), -- 大米水分
   brokenrice           varchar(200), -- 碎米率
   qingmilv             varchar(200), -- 青米率
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (feedid),
   constraint fk_feed_person foreign key (personid) references person(personid)
);

-- 出库
create table warehouse2feed (
   warehouseid        varchar(200) not null, -- 仓库id
   feedid             varchar(200) not null, -- 进料id
   personid           varchar(200),
   primary key (warehouseid, feedid),
   constraint fk_warehousetofeed_person foreign key (personid) references person(personid),
   constraint fk_warehousetofeed_warehouse foreign key (warehouseid) references warehouse(warehouseid),
   constraint fk_warehousetofeed_feed foreign key (feedid) references feed(feedid)
);

/* Table: product                              18                 */
create table product (
   productid            varchar(200) not null, -- 产品id
   name                 varchar(200), -- 名称
   specification        varchar(200), -- 规格
   flavor               varchar(200), -- 口味
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (productid),
   constraint fk_product_person foreign key (personid) references person(personid)
);

/* Table: feed2product                       16                   */
create table feed2product (
   feedid               varchar(200) not null,
   productid            varchar(200) not null,
   personid             varchar(200),
   primary key (productid, feedid),
   constraint fk_feed2product_person foreign key (personid) references person(personid),
   constraint fk_feed2product_feed foreign key (feedid) references feed(feedid),
   constraint fk_feed2product_product foreign key (productid) references product(productid)
);

/* Table: material                           17                   */
create table material (
   materialid           varchar(200) not null, -- 原料id
   kind                 varchar(200), -- 种类
   weight               varchar(200), -- 重量
   source               varchar(200), -- 来源
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (materialid),
   constraint fk_material_person foreign key (personid) references person(personid)
);

create table material2product (
   materialid         varchar(200) not null, -- 原料id
   productid          varchar(200) not null, -- 产品id
   personid           varchar(200),
   primary key (productid, materialid),
   constraint fk_material2product_person foreign key (personid) references person(personid),
   constraint fk_material2product_material foreign key (materialid) references material(materialid),
   constraint fk_material2product_product foreign key (productid) references product(productid)
);

-- 干燥表
create table dry (
   dryid                varchar(200) not null,  -- 
   productid            varchar(200),  -- 产品id
   temp                 int,           -- 温度
   drytime              int,           -- 烘烤时间
   watercontent         varchar(200),  -- 水分
   date                 date,          -- 日期
   personid             varchar(200),
   primary key (dryid),
   constraint fk_dry_person foreign key (personid) references person(personid),
   constraint fk_dry_product foreign key (productid) references product(productid)           
);

-- 膨化表
create table puffed (
   puffedid             varchar(200) not null, -- 膨化id
   productid            varchar(200), -- 产品id
   mach_temp            int,          -- 主机温度
   feed_speed           int,          -- 下料速度
   mach_speed           int,          -- 主机转速
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (puffedid),
   constraint fk_puffed_person foreign key (personid) references person(personid),
   constraint fk_puffed_product foreign key (productid) references product(productid)    
);

-- 注模表
create table injection (
   injectionid          varchar(200) not null,
   productid            varchar(200), -- 产品id
   temp                 int,          -- 温度
   drytime              int,          -- 烘烤时间
   weight               int,          -- 十枚重量     
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (injectionid),
   constraint fk_injection_person foreign key (personid) references person(personid),
   constraint fk_injection_product foreign key (productid) references product(productid)    
);

/* Table: drug                                 19                 */
create table drug (
   drugid               varchar(200) not null,  -- 药品id
   name                 varchar(200), -- 名称
   dosage               varchar(200), -- 剂型
   standard             varchar(200), -- 生产标准
   effect               varchar(200), -- 作用
   personid             varchar(200),
   primary key (drugid),
   constraint fk_drug_person foreign key (personid) references person(personid)
);

/* Table: plantusedrug                           20               */
create table plantusedrug (
   plantid              varchar(200) not null, -- 种植id
   drugid               varchar(200) not null, -- 药品id
   dosage               varchar(200), -- 浓度
   effect               varchar(200), -- 作用
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (plantid, drugid),
   constraint fk_plantusedrug_person foreign key (personid) references person(personid),
   constraint fk_plantusedrug_plant foreign key (plantid) references plant(plantid),
   constraint fk_plantusedrug_drug foreign key (drugid) references drug(drugid)  
);

/* Table: seedsoakdrug                              21                    */
create table seedsoakdrug (
   seedid               varchar(200) not null, -- 种子id
   drugid               varchar(200) not null, -- 药品id
   concentration        varchar(200), -- 浓度
   startdate            date,         -- 起始日期
   enddata              date,         -- 结束日期
   personid             varchar(200),
   primary key (seedid, drugid),
   constraint fk_seedsoakdrug_person foreign key (personid) references person(personid),
   constraint fk_seedsoakdrug_seed foreign key (seedid) references seed(seedid),
   constraint fk_seedsoakdrug_drug foreign key (drugid) references drug(drugid) 
);

/* Table: seedlingspraydrug                            22                     */
create table seedlingspraydrug (
   seedlingid           varchar(200) not null, -- 秧苗id
   drugid               varchar(200) not null, -- 药品id
   dosage               varchar(200), -- 用量
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (drugid, seedlingid),
   constraint fk_seedlingspraydrug_person foreign key (personid) references person(personid),
   constraint fk_seedlingspraydrug_seed foreign key (seedlingid) references seedling(seedlingid),
   constraint fk_seedlingspraydrug_drug foreign key (drugid) references drug(drugid) 
);



/* Table: seedling_quality_inspect         4                      */
create table seedling_quality_inspect (
   seedling_inspectid    varchar(200) not null,
   seedlingid           varchar(200) not null, -- 秧苗id
   density              varchar(200), -- 密度
   pest                 varchar(200), -- 虫害
   date                 date,         -- 日期
   growth               varchar(200),
   personid             varchar(200),
   primary key (seedling_inspectid),
   constraint fk_seedling_quality_inspect_seedling foreign key (seedlingid) references seedling(seedlingid)
);

/* Table: input_quality_inspect           5                  */
create table input_quality_inspect (
   input_inspectid      varchar(200) not null, -- 质检id
   inputid              varchar(200) not null, -- 入库id
   mildew               varchar(200), -- 霉变
   pest                 varchar(200), -- 虫害
   watercontent         varchar(200), -- 水分含量
   personid             varchar(200),
   primary key (input_inspectid),
   constraint fk_input_quality_inspect_input foreign key (inputid) references input(inputid)
);



/* Table: warehouse_inspect               6                      */
create table warehouse_inspect (
   houseinspectid       varchar(200) not null,
   warehouseid          varchar(200) not null, -- 仓库id
   watercontent         varchar(200), -- 水分
   pest                 varchar(200), -- 害虫
   housetemp            varchar(200), -- 仓库温度
   housemoisture        varchar(200), -- 仓库湿度
   graintemp            varchar(200), -- 粮堆温度
   date                 date,         -- 日期
   personid             varchar(200),
   primary key (houseinspectid),
   constraint fk_warehouse_inspect_warehouse foreign key (warehouseid) references warehouse(warehouseid)
);


/* Table: baking_quality_inspect 烘烤质检      1                 */
create table baking_quality_inspect (
   baking_inspectid     varchar(200) not null,
   productid            varchar(200) not null,
   weight               varchar(200),
   temperature          varchar(200),
   personid             int,
   primary key (baking_inspectid),
   constraint fk_baking_quality_inspect_product foreign key (productid) references product(productid)  
);

/* Table: cooling_quality_inspect         2                       */
create table cooling_quality_inspect (
   cooling_inspectid     varchar(200) not null,
   productid             varchar(200) not null,
   date                  date,
   temperature           varchar(200),
   personid              varchar(200),
   primary key (cooling_inspectid),
   constraint fk_cooling_quality_inspect_product foreign key (productid) references product(productid)
);

/* Table: package_quality_inspect          3                     */
create table package_quality_inspect (
   package_inspectid    varchar(200) not null,
   productid            varchar(200) not null,
   oxygencontent        varchar(200),
   innercapsule         varchar(200),
   date                 date,
   personid             varchar(200),
   primary key (package_inspectid),
   constraint fk_package_quality_inspect_product foreign key (productid) references product(productid)
);



