### 背景&解决方案:
09/17发现 prod-wpt-api-44 rainbow 占满一核cpu,
经过定位为loopRead中接受到stopRead消息仅break select,
没有退出整体for loop,此处从break改为return即可解决该问题

    原因追踪:
    在当前版本loopRead中可以退出for loop的方式有且仅有一种,
    select default中读包异常(例如EOF等等),
    正常case为 [读包->处理->响应] & [再次读包->EOF] 退出for loop,
    既出现该case的场景应该为"在走 select default读取到EOF之前,先收到stopRead,导致死循环"
        (1) 倒推 stopRead, 出现的场景应该为 handle write 失败
        (2) 再次读包 reader 时出现延迟

    构造以下简单 demo

    (1) ./rainbow-cpu-case
    默认情况, 无写包失败, 无再次读包延迟, 且为当前break select方式

    (2) ./rainbow-cpu-case -wf=true
    出现写包失败, 无再次读包延迟, 且为当前break select方式

    (3) ./rainbow-cpu-case -wf=true -sr=true
    出现问题场景, 在写包出现失败时, 且第二次读包出现延迟, 且为当前break select方式

    (4)./rainbow-cpu-case -el=true
       ./rainbow-cpu-case -wf=true -el=true
       ./rainbow-cpu-case -wf=true -sr=true -el=true
    修改为 exit loop, 解决问题
