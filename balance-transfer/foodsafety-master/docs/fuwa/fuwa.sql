/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017/8/22 2:22:16                            */
/*==============================================================*/

-- 6
drop table if exists baking_quality_inspect;
drop table if exists cooling_quality_inspect;
drop table if exists package_quality_inspect;
drop table if exists seeding_quality_inspect;
drop table if exists input_quality_inspect;
drop table if exists warehouse_inspect;

-- 5
drop table if exists seed;
drop table if exists seed2seeding;
drop table if exists seedling;
drop table if exists seeding2plant;
drop table if exists plant;
drop table if exists harvest;

-- 3
drop table if exists input;
drop table if exists warehouse;
drop table if exists output;

-- 4
drop table if exists feed;
drop table if exists feed2product;
drop table if exists material;
drop table if exists product;

-- 4
drop table if exists drug;
drop table if exists plantusedrug;
drop table if exists seedsoakdrug;
drop table if exists seedspraydrug;




drop table if exists person;





/*==============================================================*/
/* Table: baking_quality_inspect 烘烤质检      1                 */
/*==============================================================*/
create table baking_quality_inspect (
   weight               varchar(100),
   temperature          varchar(100),
   personid             int,
   baking_inspectid     int not null,
   productid            int not null,
   primary key (baking_inspectid)
);


/*==============================================================*/
/* Table: cooling_quality_inspect         2                       */
/*==============================================================*/
create table cooling_quality_inspect (
   date                 date,
   temperature          varchar(100),
   personid           int,
   cooling_inspectid           int not null,
   productid            int not null,
   primary key (cooling_inspectid)
);



/*==============================================================*/
/* Table: package_quality_inspect          3                     */
/*==============================================================*/
create table package_quality_inspect (
   oxygencontent        varchar(100),
   innercapsule         varchar(100),
   date                 date,
   personid           int,
   package_inspectid           int not null,
   productid            int not null,
   primary key (package_inspectid)
);


/*==============================================================*/
/* Table: seeding_quality_inspect         4                      */
/*==============================================================*/
create table seeding_quality_inspect (
   date                 date,
   growth               varchar(100),
   density              varchar(100),
   pest                 varchar(100),
   personid           int,
   seeding_inspectid           int not null,
   seedlingid           int not null,
   primary key (seeding_inspectid)
);


/*==============================================================*/
/* Table: input_quality_inspect           5                  */
/*==============================================================*/
create table input_quality_inspect (
   mildew               varchar(100),
   pest                 varchar(100),
   watercontent         varchar(100),
   input_inspectid          int not null,
   inputid            int not null,
   personid           int,
   primary key (input_inspectid)
);


/*==============================================================*/
/* Table: warehouse_inspect               6                      */
/*==============================================================*/
create table warehouse_inspect (
   houseinspectid       int not null,
   warehouseid          int not null,
   date                 date,
   watercontent         varchar(100),
   housetemp            varchar(100),
   graintemp            varchar(100),
   grainmoisture        varchar(100),
   pest                 varchar(100),
   personid           int,
   primary key (houseinspectid)
);




/*==============================================================*/
/* Table: seed                           7                       */
/*==============================================================*/
create table seed (
   seedid               int not null,
   variety              varchar(100),
   type                 varchar(100),
   personid           int,
   primary key (seedid)
);


/*==============================================================*/
/* Table: seed2seeding                      8                          */
/*==============================================================*/
create table seed2seeding (
   seedlingid           int not null,
   seedid               int not null,
   personid           int,
   primary key (seedlingid, seedid)
);


/*==============================================================*/
/* Table: seedling                          9                    */
/*==============================================================*/
create table seedling (
   seedlingid           int not null,
   nurseryplace         varchar(100),
   date                 date,
   personid           int,
   primary key (seedlingid)
);


/*==============================================================*/
/* Table: seeding2plant                      10                   */
/*==============================================================*/
create table seeding2plant (
   plantid              int not null,
   seedlingid           int not null,
   personid             int,
   primary key (plantid, seedlingid)
);


/*==============================================================*/
/* Table: plant                              11                   */
/*==============================================================*/
create table plant (
   plantid              int not null,
   place                varchar(100),
   personid           int,
   primary key (plantid)
);


/*==============================================================*/
/* Table: harvest                           12                   */
/*==============================================================*/
create table harvest (
   plantid              int not null,
   inputid            int not null,
   date                 date,
   productivity         varchar(100),
   personid           int,
   primary key (plantid, inputid)
);



/*==============================================================*/
/* Table: input                           13                    */
/*==============================================================*/
create table input (
   inputid              int not null,
   warehouseid          int not null,
   personid           int,
   primary key (inputid)
);


/*==============================================================*/
/* Table: warehouse                        14                     */
/*==============================================================*/
create table warehouse (
   warehouseid          int not null,
   feedid               int not null,
   surface              varchar(100),
   capacity             varchar(100),
   standard             varchar(100),
   personid           int,
   primary key (warehouseid)
);


-- 出库
create table output (
   outputid             int not null,
   warehouseid          int not null,
   personid           int,
   primary key (outputid)
);


/*==============================================================*/
/* Table: feed                             15                     */
/*==============================================================*/
create table feed (
   feedid               int not null,
   weight               varchar(100),
   personid           int,
   primary key (feedid)
);


/*==============================================================*/
/* Table: feed2product                       16                   */
/*==============================================================*/
create table feed2product (
   productid            int not null,
   feedid               int not null,
   watercontent         varchar(100),
   brokenrice           varchar(100),
   millet               varchar(100),
   personid           int,
   primary key (productid, feedid)
);


/*==============================================================*/
/* Table: material                           17                   */
/*==============================================================*/
create table material (
   materialid           int not null,
   productid            int not null,
   date                 date,
   weight               varchar(100),
   kind                 varchar(100),
   source               varchar(100),
   personid           int,
   primary key (materialid)
);


/*==============================================================*/
/* Table: product                              18                 */
/*==============================================================*/
create table product (
   productid            int not null,
   name                 varchar(100),
   specification        varchar(100),
   date                 date,
   flavor               varchar(100),
   personid           int,
   primary key (productid)
);




/*==============================================================*/
/* Table: drug                                 19                 */
/*==============================================================*/
create table drug (
   drugid              int not null,
   standard             varchar(100),
   name                 varchar(100),
   dosage               varchar(100),
   effect               varchar(100),
   personid           int,
   primary key (drugid)
);


/*==============================================================*/
/* Table: plantusedrug                           20               */
/*==============================================================*/
create table plantusedrug (
   plantid              int not null,
   drugid              int not null,
   effect               varchar(100),
   date                 date,
   dosage               varchar(100),
   personid           int,
   primary key (plantid, drugid)
);


/*==============================================================*/
/* Table: seedsoakdrug                              21                    */
/*==============================================================*/
create table seedsoakdrug (
   seedid               int not null,
   drugid              int not null,
   concentration        varchar(100),
   startdate            date,
   enddata              date,
   personid           int,
   primary key (seedid, drugid)
);



/*==============================================================*/
/* Table: seedspraydrug                            22                     */
/*==============================================================*/
create table seedspraydrug (
   drugid              int not null,
   seedid               int not null,
   date                 date,
   dosage               varchar(100),
   personid           int,
   primary key (drugid, seedid)
);




create table person (
   personid             int not null,
   account              varchar(100),
   name                 varchar(100),
   password             varchar(100),
   primary key (personid)
);



alter table baking_quality_inspect add constraint FK_product_bake_inspect foreign key (productid)
      references product (productid) on delete restrict on update restrict;

alter table cooling_quality_inspect add constraint FK_product_cool_inspect foreign key (productid)
      references product (productid) on delete restrict on update restrict;

alter table seeding_quality_inspect add constraint FK_checkseeding foreign key (seedlingid)
      references seedling (seedlingid) on delete restrict on update restrict;

alter table input_quality_inspect add constraint FK_checkinput foreign key (inputid)
      references input (inputid) on delete restrict on update restrict;

alter table warehouse_inspect add constraint FK_checkhouse foreign key (warehouseid)
      references warehouse (warehouseid) on delete restrict on update restrict;

alter table package_quality_inspect add constraint FK_product_package_inspect foreign key (productid)
      references product (productid) on delete restrict on update restrict;



-- process

alter table seed2seeding add constraint FK_seed2seeding foreign key (seedlingid)
      references seedling (seedlingid) on delete restrict on update restrict;

alter table seed2seeding add constraint FK_seed2seeding2 foreign key (seedid)
      references seed (seedid) on delete restrict on update restrict;


alter table seeding2plant add constraint FK_seeding2plant foreign key (plantid)
      references plant (plantid) on delete restrict on update restrict;

alter table seeding2plant add constraint FK_seeding2plant2 foreign key (seedlingid)
      references seedling (seedlingid) on delete restrict on update restrict;

alter table input add constraint FK_enterhouse foreign key (warehouseid)
      references warehouse (warehouseid) on delete restrict on update restrict;

alter table output add constraint FK_outputhouse foreign key (warehouseid)
      references warehouse (warehouseid) on delete restrict on update restrict;

alter table warehouse add constraint FK_house2feed foreign key (feedid)
      references feed (feedid) on delete restrict on update restrict;

alter table harvest add constraint FK_harvest foreign key (inputid)
      references input (inputid) on delete restrict on update restrict;

alter table harvest add constraint FK_harvest2 foreign key (plantid)
      references plant (plantid) on delete restrict on update restrict;

alter table material add constraint FK_material2product foreign key (productid)
      references product (productid) on delete restrict on update restrict;

alter table feed2product add constraint FK_feed2product foreign key (feedid)
      references feed (feedid) on delete restrict on update restrict;

alter table feed2product add constraint FK_feed2product2 foreign key (productid)
      references product (productid) on delete restrict on update restrict;





-- use drug

alter table plantusedrug add constraint FK_plantusedrug foreign key (drugid)
      references drug (drugid) on delete restrict on update restrict;

alter table plantusedrug add constraint FK_plantusedrug2 foreign key (plantid)
      references plant (plantid) on delete restrict on update restrict;


alter table seedsoakdrug add constraint FK_seedsoakdrug foreign key (drugid)
      references drug (drugid) on delete restrict on update restrict;

alter table seedsoakdrug add constraint FK_seedsoakdrug2 foreign key (seedid)
      references seed (seedid) on delete restrict on update restrict;


alter table seedspraydrug add constraint FK_seedspraydrug foreign key (seedid)
      references seed (seedid) on delete restrict on update restrict;

alter table seedspraydrug add constraint FK_seedspraydrug2 foreign key (drugid)
      references drug (drugid) on delete restrict on update restrict;


-- person
alter table baking_quality_inspect add constraint FK_baking_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table cooling_quality_inspect add constraint FK_cooling_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table package_quality_inspect add constraint FK_package_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seeding_quality_inspect add constraint FK_seeding_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table input_quality_inspect add constraint FK_input_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table warehouse_inspect add constraint FK_warehouse_inspect_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seed add constraint FK_seed_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seed2seeding add constraint FK_seed2seeding_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seedling add constraint FK_seedling_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seeding2plant add constraint FK_seeding2plant_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table plant add constraint FK_plant_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table harvest add constraint FK_harvest_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table input add constraint FK_input_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table warehouse add constraint FK_warehouse_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table output add constraint FK_output_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table feed add constraint FK_feed_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table feed2product add constraint FK_feed2product_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table material add constraint FK_material_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table product add constraint FK_product_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table drug add constraint FK_drug_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table plantusedrug add constraint FK_plantusedrug_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;

alter table seedsoakdrug add constraint FK_seedsoakdrug_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;
      
alter table seedspraydrug add constraint FK_seedspraydrug_person foreign key (personid)
      references person (personid) on delete restrict on update restrict;