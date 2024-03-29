### 云南省医疗保障信息平台定点医药机构接口动态库版本更新说明

##### CHSInterfaceYn.dll

| 版本号  | 更新说明                                                     |
| ------- | ------------------------------------------------------------ |
| 1.0.1.7 | 解决文件上传和下载问题                                       |
| 1.0.1.8 | 解决文件下载异常情况的错误提示不正确问题                     |
|         | 解决公告提醒问题                                             |
| 1.0.1.9 | 解决电子凭证扫码异常问题                                     |
| 1.0.2.1 | 解决 1191交易电子凭证返回节点【ecinfo】字段【 ectoken 】字母小写 |
| 1.0.2.2 | 解决1192核验成功标志返回不正确问题                           |
| 1.0.2.4 | 解决社保卡动态库兼容性问题                                   |
| 1.0.2.5 | 解决社保卡提醒窗口异常导致程序崩溃                           |
| 1.0.2.6 | 解决社保卡动态库路径问题，解决xp兼容性问题                   |
| 1.0.2.7 | 解决社保卡读卡提醒窗口bug                                    |
| 1.0.2.8 | 增加函数BusinessHandleW 支持unicode字符集                    |
| 1.0.2.9 | 解决 runtime error 226 问题以及1101返回值【姓名】存在繁体字时不能正常返回的bug |
| 1.0.3.0 | 1101返回值【姓名】存在特殊字符时不能正常返回的bug            |
| 1.0.3.1 | 支持自定义日志路径、文件下载路径                             |
|         | 1101增加密码输入功能【重要】                                 |
| 1.0.3.2 | 修复1.0.3.1版本的bug                                         |
| 1.0.3.3 | 修复文件下载路径不正确bug                                    |
| 1.0.3.4 | 解决社保卡姓名中含有特殊字符导致不能读卡的bug                |
|         | 优化初始化方法                                               |
|         | 1191返回报文【cardinfo】节点增加字段【certno】证件号码       |
| 1.0.3.5 | 解决社保卡卡号存在空格的问题，处理1101交易姓名中有特殊字符的问题 |
| 1.0.3.6 | 支持刷身份证就医                                             |
| 1.0.3.7 | 解决外省卡统筹区是FFFFFF的问题；优化刷身份证就医；解决1101特殊姓名的问题 |
| 1.0.3.9 | 解决社保卡token失效问题                                      |
| 1.0.4.0 | 解决xp系统兼容性问题                                         |
| 1.0.4.1 | 支持昆明三代社保卡                                           |
| 1.0.4.2 | 支持GB18030中文编码字符集，解决特殊汉字显示?问题             |
| 1.0.4.3 | 支持GB18030中文编码字符集，解决特殊汉字显示?问题（优化1042bug） |
| 1.0.4.4 | 支持终端设备                                                 |
| 1.0.4.5 | 增加 NationEcTrans 函数                                      |
| 1.0.4.6 | 临时版本                                                     |
| 1.0.4.7 | 解决生僻字问题                                               |



#### ziptool.dll

| 版本号  | 更新说明                           |
| ------- | ---------------------------------- |
| 1.0.1.1 | 解决解压时发生异常时文件占用的问题 |
| 1.0.2.0 | 解决特殊压缩格式不能解压的问题     |



#### SCardLibYN.dll

| 版本号  | 更新说明                                                |
| ------- | ------------------------------------------------------- |
| 1.0.2.1 | 增加读卡错误的中文描述                                  |
| 1.2.4.0 | 通过scardbrood.ini中节点【DeviceWorkDir】配置动态库驱动 |
| 1.2.6.0 | 增加日志记录                                            |
| 1.2.7.0 | 解决选择发卡方文件失败                                  |
| 1.2.8.0 | 解决错误： 初始化机构代码长度不正确(应为24)             |
| 1.2.9.0 | 兼容德卡读卡器（T10）                                   |
| 1.3.1.0 | 处理卡号中存在的空格                                    |
| 2.0.0.0 | 优化读卡                                                |
| 2.0.1.0 | 解决北京卡用户卡不合法问题                              |
| 2.0.2.0 | 非接触式读卡器优化                                      |
| 2.0.3.0 | 解决外省卡行政区划FFFFFF的问题                          |
| 2.0.7.0 | 支持昆明市三代卡                                        |
| 2.0.8.0 | 解决州市卡转移到昆明参保后不能修改密码的问题            |
| 2.0.9.0 | 优化                                                    |
| 2.0.9.1 | 版本冲突，升个小版本                                    |

#### chsinterfaceyn.ocx

| 版本号  | 更新说明                                  |
| ------- | ----------------------------------------- |
| 1.0.1.2 | 解决文件上传bug                           |
| 1.0.3.0 | 增加函数BusinessHandleW 支持unicode字符集 |
| 1.0.4.0 | 增加NationEcTrans方法                     |

#### GP700.dll

| 版本号   | 更新说明                                                     |
| -------- | ------------------------------------------------------------ |
| 3.0.1.7  | 解决gp700读卡器闪退                                          |
| 3.0.26.0 | 支持 t10 读卡器；自动选择卡片操作界面；改PIN操作流程的修改；读卡器验证逻辑 |

