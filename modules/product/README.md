# Product模块

#####  

- 分类、商品、sku 等的 新增 、创建 、 修改 、查询

##### 接口方法
- 产品分类 增删改查
- 产品属性 增删改查
- 产品sku 增删改查

#### 产品类目(Category)
开始默认 使用 系统类目(全部), 无类别 ，商品数量多了以后使用 一级分类或多级分类。 参考设计  http://www.woshipm.com/pd/1919218.html
 

### 数据库表结构

- [数据库表关系图](https://dbdiagram.io/embed/5ce3b1d51f6a891a6a656266) 

**商品类目表 categorys**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   pid          |  bigint(20)   |   NOT NULL          |   父级CategorysId    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   status          |  int(11)   |   DEFAULT NULL         |   状态    |
|   sort          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   resource_id          |  bigint(20)   |   DEFAULT NULL         |   商品图片资源Id    |
|   image_json          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |



**规格属性 attributes**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   category_id          |  bigint(20)   |   NOT NULL          |   类目ID    |
|   status          |  int(11)   |   DEFAULT NULL         |   状态    |
|   sort          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |


**规格属性组 attribute_groups**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   product_id          |  bigint(20)   |   NOT NULL          |   产品ID    |
|   status          |  int(11)   |   DEFAULT NULL         |   状态    |
|   sort          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

 **规格属性SKU组 attribute_group_sku**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   product_sku_id          |  bigint(20)   |   NOT NULL          |   产品skuID    |
|   attribute_group_id          |  bigint(20)   |   NOT NULL          |   规格属性组ID     |
|   status          |  int(11)   |   DEFAULT NULL         |   状态    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**规格属性选项 attribute_options**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   attribute_id          |  bigint(20)   |   NOT NULL          |   属性ID    |
|   status          |  int(11)   |   DEFAULT NULL         |   状态    |
|   sort          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**规格属性 选项 组 attribute_option_groups**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    | 
|   attribute_group_id          |  bigint(20)   |   NOT NULL          |   属性组ID    |
|   attribute_option_id          |  bigint(20)   |   NOT NULL          |   属性选项ID    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**产品表 products**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   guid          |  varchar(255)   |   NOT NULL  unique_index       | GUID 产品数据库唯一标记    |
|   category_id          |  bigint(20)   |   NOT NULL          |   类目ID    |
|   mer_id          |  bigint(20)   |   NOT NULL          |   商户ID    |
|   type_id          |  bigint(20)   |   NOT NULL|类型ID 默认 0  [1、套餐 2、虚拟产品 3、积分产品]   |
|   brand_id          |  bigint(20)   |   NOT NULL          |   品牌ID (暂无、保留字段)    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   产品名称    |
|   status          |  int(11)   |   DEFAULT NULL         |   状态   0默认未上架 1 上架 3 下架  |
|   content          |  text   |            |   产品富文本信息   |
|   description          |  varchar(2000)   |            |   描述信息   |
|   keywords          |  varchar(255)   |            |   产品关键词   |
|   tags          |  varchar(255)   |            |   标签   |
|   original_price          |  decimal(18,2)   |            |   原始价格   |
|   min_price          |  decimal(18,2)   |            |   最低价格   |
|   current_price          |  decimal(18,2)   |            |   当前销售价格 (待删除)  |
|   sales          |  int(11)   |   DEFAULT NULL         |   销量(总计)    |
|   location          |  varchar(255)   |            |   经纬度信息(暂留)   |
|   width          |  decimal(10,2)   |            |   宽   |
|   height          |  decimal(10,2)   |            |   高   |
|   depth          |  decimal(10,2)   |            |   深度   |
|   weight          |  decimal(10,2)   |            |   重量(kg)   |
|   product_type          |  int(11)|  DEFAULT NULL|产品类型  0:商品(默认) 1:服务    |
|   integral          |  int(11)   |            |   可以使用积分抵消   |
|   active          |  bool   |            |   是否启用   | 
|   image_json          |  text   |            |   产品图片对象json 字符串(冗余字段)   |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**产品价格表(如果发生变化会记录下来，默认最后一条记录的价格是最新的) product_price**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   product_id          |  bigint(20)   |   NOT NULL          |   产品Id    |
|   product_sku_id          |  bigint(20)   |   NOT NULL          |   产品SKUID    |
|   price          |  decimal(18,2)   |            |   价格   |
|   operator_admin_id          |  bigint(20)   |   NOT NULL          |   操作管理员用户编号    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**产品 静态资源表(图片、视频) product_resources**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   resource_id          |  bigint(20)   |   NOT NULL          |   ResourceId    |
|   resource_guid          |  varchar(255)   |   NOT NULL  unique_index       | ResourceGUID|
|   product_id          |  bigint(20)   |   NOT NULL          |   产品ID    |
|   product_guid          |  varchar(255)   |   NOT NULL  unique_index       | 产品GUID|
|   cover          |  bool   |            |   是否封面   |
|   type          |  int(11)   |   NOT NULL          |   类型(暂留字段)    |
|   position          |  int(11)   |   NOT NULL          |  位置    |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |

**产品sku product_sku**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   guid          |  varchar(255)   |   NOT NULL  unique_index       | GUID 产品数据库唯一标记    |
|   product_id          |  bigint(20)   |   NOT NULL          |   产品ID    |
|   name          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   code          |  varchar(255)   |   DEFAULT NULL         |   商品编码    |
|   bar_code          |  varchar(255)   |   DEFAULT NULL         |   商品条形码    |
|   original_price          |  decimal(18,2)   |            |   原始价格   |
|   price          |  decimal(18,2)   |            |   价格   |
|   stock          |  int(11)   |   DEFAULT NULL         |  库存 |
|   width          |  decimal(10,2)   |            |   宽   |
|   height          |  decimal(10,2)   |            |   高   |
|   depth          |  decimal(10,2)   |            |   深度   |
|   weight          |  decimal(10,2)   |            |   重量(kg)   |
|   status          |  int(11)   |   DEFAULT NULL         |   默认 0 预留字段 暂无意义  |
|   sort          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   resource_id          |  bigint(20)   |   DEFAULT NULL         |   商品图片资源Id    |
|   attribute_info|varchar(255)|DEFAULT NULL| 规格属性信息 ，冗余字段 更新商品sku 规格属性需要更新该字段|
|   image_json          |  text   |            |   产品图片对象json 字符串(冗余字段)   |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |


