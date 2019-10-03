# CLI 命令行实用程序开发基础
CSDN同步博客地址:https://blog.csdn.net/dcy19991116/article/details/101987271
## 1.设计说明
我们这次的任务是开发一个CLI命令行实用程序，CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，
通过shell或script使得应用获得最大的灵活性与开发效率
上面这一段是作业网站告诉我们的关于我们本次作业的内容，更通俗一点的解释就是一个类似于shell的东西，往里面输入一个命令，就会得到相应的结果，比如，输入-ls，就会显示当前目录下的文件，不过我们这次要完成的目标如下

> 该实用程序从标准输入或从作为命令行参数给出的文件名读取文本输入。它允许用户指定来自该输入并随后将被输出的页面范围。例如，如果输入含有 100 页，则用户可指定只打印第 35 至 65 页。这种特性有实际价值，因为在打印机上打印选定的页面避免了浪费纸张。另一个示例是，原始文件很大而且以前已打印过，但某些页面由于打印机卡住或其它原因而没有被正确打印。在这样的情况下，则可用该工具来只打印需要打印的页面

所以我们本次作业的目的可以简单概括为用命令行来控制所打印的文本，因为我们需要对命令行进行操作，所以我们需要在Golang环境下处理命令行，这里面用到了flag包，我在这里使用的是pflag，下载操作如下:

```
go get github.com/spf13/pflag
```
flag包是一个命令行参数的解析工具，用法的话我在做的过程中参考了以下几篇博客
[go语言学习-flag包的使用](https://blog.csdn.net/len_yue_mo_fu/article/details/81041793)
[golang flag包使用笔记](https://www.jianshu.com/p/f9cf46a4de0e)

其实整个流程就是我们在一个flag.XXXVar()语句里面将变量，所对应的命令行参数，以及默认值与对这个变量的说明绑定在一起
我们第一步要做的事解析我们所要输入的命令
首先先定义一个结构体，来表明我们将要进行操作的参数

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003140112807.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

六个变量的作用依次如下
start_page:开始的页面

end_page:结束的页面

in_filename:用于输入的文件名

print_dest:用于输出的文件名 

page_len:每页的行数(在这里默认为72) 

page_type: 打印的模式，'l'按行打印，'f'按换页符打印 

然后我们将这些变量绑定在flag上，然后设置Usage，最后通过Parse()函数来进行解析

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003140757866.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

接下来就是对一些异常的情况进行处理

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003141340169.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

首先是对参数的个数进行判断，如果参数小于3，则一定有问题
其次对第一个位置和第三个位置的参数进行检查，若都一个参数不是“-s”或第三个参数不是“-e”，则错误
因为命令格式如下 -s (start_page) -e (end_page)
第二个参数是开始页面，第四个参数是终止页面
至于第零个参数，是下面这个东西

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003141827834.png)

接下来是对开始页面和结束页面进行判断
如果开始/结束页面小于1或大于最大页面(在这里我设置成了99999999)，则出现错误
如果结束页面比开始页面还要小，则也出现错误
接着是对页面长度进行判断，小于1或者大于最大页面，都是有错误的
如果还有剩余的参数的话，一定是作为输入的文件名，检查这个文件是否存在，若不存在，则报错

对命令行参数的解析至此可以告一段落了
接下来就是对输入输出进行操作
在这个程序里面，输入可以直接在终端使用输入流进行输入，也可以通过文件来进行输入

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003143531475.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

如果这个作为输入的文件的文件名字存在，那么就打开这个文件，若文件为空，则报错，并最终关闭文件，否则，输入则按照标准输入流来进行输入

而输出则可以通过文件输出，也可以通过管道进行输出或是直接输出到屏幕上面

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003144112751.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

以上主要是用来设置输出地点的函数，通过使用cmd来确定输出的地点在哪里

最后是设置打印的模式，-l(按行来打印)-f(按分页符打印)，具体操作分别如下

![在这里插入图片描述](https://img-blog.csdnimg.cn/201910031449235.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003144937523.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

## 2.使用与测试
因为这里面测试数据我只设置了10行，所以若是不出现错误的话，开始和结尾页码都是1，设置报错信息的话另说
test.txt为包含了测试数据的文件，这里用作输入文件
test1.txt为空白文件，这里用作输出文件
error.txt为空白文件，用于输出报错信息

-s 1 -e 1 test.txt
将test.txt的信息输出到屏幕上

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003151658215.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

-s 1 -e 1 <test.txt
将test.txt第一页的内容输出到屏幕上

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152043849.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

-s 1 -e 1 test.txt >test1.txt
将test的内容输出到test1中

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152302509.png)

执行后的test1

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152421118.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

-s 1 -e 5 test.txt 2>error.txt
将test的内容输出到屏幕上，报错的信息则输出到error中
这里面我设置的是输出第1页-第5页，但是整个文件就只有一页，所以错误出现在这里

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152552365.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152607299.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

-s 1 -e 5 test.txt >test1.txt 2>error.txt
将test的信息输出到test1中，错误信息输出到error.txt
错误信息同上

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152917405.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152933882.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003152945290.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)


-s 1 -e 1 test.txt >/dev/null
什么也不输出

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153258792.png)

-s 1 -e 1 test.txt | wc
输出行数，单词数以及字节数

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153513228.png)

-s 1 -e 1 -l 5 test.txt
输出5行

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153704770.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)


-s 1 -e 1 -f test.txt
按照分页符来输出

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153808983.png)


-s 1 -e 1 -d test1.txt test.txt
将test的内容输出到test1里面去

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153947155.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003153959666.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)


cat test.txt | -s 1 -e 1
连接test.txt并打印到标准输出设备上(在这里是屏幕)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191003154053704.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2RjeTE5OTkxMTE2,size_16,color_FFFFFF,t_70)

