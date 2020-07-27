struct:

- 育秧 Seeding
	- 种子管理 SeedManagement
		- 种子基本信息 seed
		- 种子浸泡信息 seedsoakdrug
		- 种子喷药信息 seedspraydrug
	- 药品管理 DrugManagement
		- 药品信息 drug
	- 秧苗管理 SeedlingManagement
		- 秧苗信息 seeding * seedling

- 种植 Plant
	- 种植管理 PlantManagement
		- 种植信息 plant
		- 种植用药信息（除虫施肥） plantusedrug
	- 收割管理 HarvestManagement
		- 收割信息 harvest  ->delete


			+plant2input
- 仓储 Storage
	- 入库管理 input * 
	 		  +input2warehouse
	- 仓库管理 warehouse *
	- 出库管理 output -> warehouse2feed

- 粗加工 RoughProcess
	- 进料管理 feed *
	- 出料管理 feed2product

- 深加工 DeepProcess
	- 原料加工 RawMaterialProcessing
		- 原料信息 material ?
				  + material2product
	- 产品管理 ProductManagement 
		- 产品信息 product
		 		+ injection
		 		+ dry
		 		+ puffed


- 质检管理 Inspect
	- 秧苗质检 seeding_quality_inspect
	- 入库质检 input_quality_inspect
	- 仓库质检 warehouse_inspect
	- 深加工质检 QualityInspectionDeepProcessing
		- 烘烤质检 baking_quality_inspect
		- 冷却质检 cooling_quality_inspect
		- 内包装质检 package_quality_inspect

- 用户管理 Person
	- 用户信息 person *

- 登陆 Login
	- 登陆 login
	- 注册 register
	
- index
- error
