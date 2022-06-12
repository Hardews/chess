# chess
不够时间了就不按规范写了。。。（dbq

这里展示棋盘是用的cmd，然后介绍两个接口


## `GET` `:8080/new`
### `创建一个新房间`
#### `这个是用websocket做的，在websocket在线测试网站上可以链接到，在都准备后可以游戏开始然后才可以发送指令`
### `指令说明`

| 指令        | 类型 | 备注                 |
|-----------| ---- |--------------------|
| choose    | int | 必选，选择移动哪个棋子        |
| num       | int | 必选，选择移动这个棋子的第几个棋子  |
| direction | int | 必选，棋子移动的方向         |
| step      | int | 必选，棋子移动的步数         |
以上指令都要填，如果漏填指令不会执行

| choose | 备注   |
|--------|------|
| 1      | 炮    |
| 2      | 马    |
| 3      | 士    |
| 4      | 车    |
| 5      | 象    |
| 6      | 帅    |
| 7      | 兵    |


| num | 备注                           |
|-----|------------------------------|
| 0   | 红方是从左边数的第一个这个类型的棋子，绿方相反，下面类同 |
| 1   |                              |
| 2   |                              |
| 3   |                              |
| 4   |                              |

| direction | 备注             |
|-----------|----------------|
| 1         | 北              |
| 2         | 东              |
| 3         | 南              |
| 4         | 西              |
| 12        | 东北             |
| 23        | 东南             |
| 34        | 西南             |
| 41        | 西北             |
| 122       | 这个为马的专属方向，即走横日 |
| 233       |                |
| 344       |                |
| 411       |                |

| step | 备注       |
|------|----------|
| 1    | 向某个方向走一步 |


## `GET` `:8080/join/:num`
### `加入一个新房间`

| 请求参数 | 类型     | 备注      |
|------|--------|---------|
| :num | string | 必选，房间号  |

发送指令过后不会是用websocket发信息给客户端（因为棋盘美观问题），但会在程序的cmd处显示出来效果


1.实现功能

登陆注册：用了jwt，但是方便测试就没有把创建房间这些功能放在鉴权里面

加入和创建房间：这个功能用的是websocket实现，还用了redis的集合来存储房间号，然后在展示房间的时候取出里面的元素。
用了一个map来存房间的client，在加入房间时会判断房间是否满人，满人了就不能再加入。

准备和未准备状态： 在用户创建房间后（默认未准备）向websocket服务端发送一个1代表准备，
发送一个0代表未准备，当双方都准备时可开始游戏

基础功能差俩没写。
基础的功能判断输赢没写，判断输赢的逻辑来不及写了，想法是每次用户动棋子时判断下一步的方向从而实现是否提醒将军，
然后通过遍历坐标来判断是否为死棋（飞将，马后炮等等）
轮流着棋的话其实没有严格要求，但如果用户同时动两次会报错。
但是有个想法就是用两个channel来控制他们的着棋

还想用websocket参照官方的chat来写一个房间的聊天室

2.象棋逻辑

象棋的移动是没有问题的，可以正常吃子，有些细节被我添加了，如：帅和士不能出田字格，兵不能后退等。但是有些细节没做到位：比如马撇脚这些。
然后根据红方和绿方差异，遍历打印这个棋盘的方法也不一样（有一边是倒着遍历的），这样
无论是红方还是绿方，都是自己这边的棋子（帅）对着自己（我不懂怎么解释这里我的想法，我在说什么）

发现自己不会的东西很多，写的很水，可能也会有bug，实现的东西也不多。谢谢学长可以看到这里。