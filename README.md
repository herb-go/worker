# 工作者 模块

用于管理Herb go 框架中的全局单例对象(Worker)的模块

## 引入类

* Worker工作者对象，通过名称注册单例对象到系统内统一管理
* Overseer监工接口/类，提供统一的管理类，实现 查看工作对象状态，生成标准化报告，初始化工作对象属性，发布管理命令的接口。

## 解决的问题
* 提供全局范围内的工作者对象的引用，实现工作类对象如缓存，数据库等的复用，
* 实现预埋点，可以根据项目发展对指定已埋点的工作者对象做特殊设置
* 实现通用的系统状态报告切入点
* 实现通用的管理命令接口

## 约定
* 工作者类通过反射类型，匹配到对应的监工类。
* 工作者类必须为结构指针或者接口，第一次赋值时必须实例化不可为空，赋值后不可修改(指向其他对象)
* 监工类返回的评估信息和命令接口的返回值为可选项，返回的结构可以通过json或者其他兼容的方式序列化。

