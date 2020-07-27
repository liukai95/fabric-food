struct:

- 育秧
	- 种子管理
		- 种子基本信息 seed
		- 种子浸泡信息 seedsoakdrug
		- 种子喷药信息 seedspraydrug
		- 种子秧苗转换信息信息 seed2seedling
	- 药品管理
		- 药品信息 drug
	- 秧苗管理
		- 秧苗信息 seeding

- 种植
	- 种植管理
		- 种植信息 plant
		- 种植用药信息（除虫施肥） plantusedrug
		- 秧苗种植转换信息 seedling2plant
	- 收割管理
		- 收割信息 harvest

- 仓储
	- 入库管理 input
	- 仓库管理 warehouse
	- 出库管理 output

- 加工
	- 进料管理 feed
	- 出料管理 feed2product

- 深加工
	- 原料加工 
		- 原料信息 material
	- 产品管理 
		- 产品信息 product



- 质检管理
	- 秧苗质检 seeding_quality_inspect
	- 入库质检 input_quality_inspect
	- 仓库质检 warehouse_inspect
	- 深加工质检
		- 烘烤质检 baking_quality_inspect
		- 冷却质检 cooling_quality_inspect
		- 内包装质检 package_quality_inspect

- 用户管理
	- 用户信息 person