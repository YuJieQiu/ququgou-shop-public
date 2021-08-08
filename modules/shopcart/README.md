### 购物车模块


### 数据库表结构

购物车结构: 一个用户和一个sku 为一条记录,可以执行更新删除等
区分商家



**购物车表 categorys**

|   字段    |    数据类型    |   NULL、DEFAULT...  |           描述          |
|   ---    |     ----      |   ----      |  -------------------------------    |
|   id          |  bigint(20)   |   NOT NULL AUTO_INCREMENT         |   主键自增    |
|   shop_cart_no          |  bigint(20)   |   NOT NULL          |   父级CategorysId    |
|   mer_id          |  varchar(255)   |   DEFAULT NULL         |   名称    |
|   product_id          |  int(11)   |   DEFAULT NULL         |   状态    |
|   product_sku_id          |  int(11)   |   DEFAULT NULL         |   顺序    |
|   number          |  bigint(20)   |   DEFAULT NULL         |   商品图片资源Id    |
|   price          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   original_price          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   total_price          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   product_name          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   product_description          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   location          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   width          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   height          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   depth          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   weight          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   is_package          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   is_virtual          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   is_integral          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   cover_image          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   attribute_info          |  text   |            |   图片对象json 字符串(冗余字段)   |
|   created_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   创建时间    |
|   updated_at  |  timestamp    |   NULL DEFAULT CURRENT_TIMESTAMP  |   更新时间    |
|   deleted_at  |  timestamp    |   NULL DEFAULT NULL               |   删除时间    |




