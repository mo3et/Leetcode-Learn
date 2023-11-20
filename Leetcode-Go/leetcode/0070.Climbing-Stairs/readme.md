# 70.爬楼梯

图解：[link](https://blog.csdn.net/yangxinxiang84/article/details/121278068)
![](https://img-blog.csdnimg.cn/7b0d703445ee4359bb69e04c7a2c8ef5.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAeWFuZ3hpbnhpYW5nODQ=,size_20,color_FFFFFF,t_70,g_se,x_16)

[视频讲解](https://www.bilibili.com/video/BV17h411h7UH)

主要思路是 DP 动态规划

可以使用滑动数组

## 扩展

若要爬台阶，一次跨 N 阶，则公式为
`F(n)=F(n-1)+F(n-2)+F(n-3)`
