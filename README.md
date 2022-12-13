核心价值观加解密工具(golang版本)

### 版本
v1.0.0

### 最后更新 
2022-10-27

### 命令说明

算法来自于 [@sym233](https://www.github.com/sym233/core-values-encoder/)

命令格式: jzg <命令> <要加密或解密的串>

命令包含以下两种:
  d  要解密的文本.
  e  需要加密的文本

示例:

```
> go run main.go e 你好，世界
>
> 友善爱国自由友善平等诚信和谐友善自由富强友善爱国平等诚信富强平等诚信民主诚信和谐文明友善公正诚信自由自由友善平等爱国敬业公正诚信自由法治敬业平等爱国友善公正

> go run main.go d 友善爱国自由友善平等诚信和谐友善自由富强友善爱国平等诚信富强平等诚信民主诚信和谐文明友善公正诚信自由自由友善平等爱国敬业公正诚信自由法治敬业平等爱国友善公正
>
> 你好，世界
```

### 版权声明
算法版权归原作者 [@sym233](https://www.github.com/sym233/core-values-encoder/) 所有，请勿使用于任何违法行为，一切用于违法行为的后果请自负
