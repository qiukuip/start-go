---
title: go-node.md
date: 2026-04-24 13:38:58
---



## 模块

模块初始化：
```bash
go mod init
```

模块引用：
```bash
go mod edit -replace day.happy365/greetings=../greetings
```

获取依赖
```bash
go mod tidy

go get .
```

