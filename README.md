# goWLB 

# go-王老板

*用gogs-webhooks来统计你的员工写了多少代码！*

*generate coding statistic information of your group， with gogs-webhooks*

## 动机

最近在实验室实践基于工单--提交的闭环管理模式，
但是发现实验室的私有gogs服务自己没有代码量统计、每个用户的提交量的分析（等等的一众统计功能）
等等的一些必要功能，搜了一下，不知道是因为这种东西太简单还是什么别的原因，
([搜到了一个这个东西，但是我不懂这是西班牙语还是啥反正不会用就是了](https://github.com/lfkimura/webhook-statistics))
反正是没人写原生的gogs插件来实现这个功能。这就对俺们的KPI考核产生了一定的影响。

既然没得原生插件，那可以简单的用webhook来做一个外挂的服务：

```
-----------                -----------------
|   repos  |--webhooks---->|   go-WLB     |
-----------                -----------------
```

在你的私有git服务器上启动一个go-WLB，
然后要求每一个repo都把webhook推送到go-WLB，
后者就会把这些信息汇总起来，然后存到自己的小数据库里。

然后就可以按项目、成员、组织、时间分类来统计每一个小可爱的代码贡献了。

## 设计

WLB的功能非常简单：接受webhook的post请求，然后写到一个数据库里面，
分析和统计的功能都交给数据库和SQL来做（这样是不是有点偷懒）

对于Webhook的格式，一开始准备只支持gogs格式，后续可以把其他格式再支持进来，这列了个表：

| Webhook | Support |
|---------|---------|
|  gogs   |    Y    |
|  Slack  |    X    |
|  Github |  Unknown|

各种各样的repo活动其实都会推送webhook，但是一开始我打算只处理push活动，
gogs的push webhook粒度最细是这次push包含哪些commit，
不包含每个commit有哪些diff（实际上我希望后续可以细化到能够统计出来代码的反复横跳，这样就可以狠怼需求方）
所以目前只能做到统计commit的粒度。


前端我是一点也不想写的，所以准备用grafana来作前端。



