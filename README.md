# Go 语言学习代码库

这个代码库包含了一系列 Go 语言的学习示例代码，涵盖了从基础语法到高级特性的各个方面，适合初学者快速上手 Go 语言的基础概念和常用操作。

## 📁 项目结构

```
go/
├── 基础语法示例/
│   ├── hello.go              # 基础 Hello World 程序
│   ├── hello_v2.0.go         # Hello World 进阶版本
│   ├── helloword_v2.go       # Hello World 另一个版本
│   ├── list.go               # 数组初始化与打印示例
│   ├── 数组特性.go           # 数组特性演示
│   ├── 字符串.go             # 字符串操作示例
│   └── go_slice.go           # Slice 切片操作详解
├── 函数与面向对象/
│   ├── func.go               # 函数定义与使用示例
│   └── go_oop.go             # Go 面向对象编程示例
├── 并发编程/
│   ├── routine.go            # Goroutine 协程示例
│   ├── channel.go            # Channel 通道使用示例
│   ├── 生产者消费模型简化版.go # 生产者消费者模式实现
│   └── pubsub/               # 发布订阅模式
│       └── pubsub.go         # 发布订阅实现
├── 设计模式/
│   └── Singleton_Pattern.go  # 单例模式实现
├── 算法练习/
│   └── leetcode/             # LeetCode 算法题
│       ├── two-sum.go        # 两数之和
│       └── revertedNumber.go # 数字反转
└── 项目文件/
    ├── go.mod                # Go 模块文件
    ├── pubsub_main.go        # 发布订阅主程序
    └── README.md             # 项目说明文档
```

## 🚀 快速开始

### 环境要求
- Go 1.16 或更高版本

### 安装与运行

1. **克隆仓库**
   ```bash
   git clone <repository-url>
   cd go
   ```

2. **运行基础示例**
   ```bash
   # 运行 Hello World
   go run hello.go
   
   # 运行数组示例
   go run list.go
   
   # 运行字符串操作
   go run 字符串.go
   ```

3. **运行并发示例**
   ```bash
   # 运行 Goroutine 示例
   go run routine.go
   
   # 运行 Channel 示例
   go run channel.go
   
   # 运行生产者消费者模型
   go run 生产者消费模型简化版.go
   ```

4. **运行发布订阅示例**
   ```bash
   go run pubsub_main.go
   ```

## 📚 学习内容

### 基础语法
- **Hello World**: 学习 Go 程序的基本结构
- **变量与数据类型**: 了解 Go 的基本数据类型
- **数组与切片**: 掌握数组和切片的操作
- **字符串操作**: 学习字符串的常用方法

### 函数与面向对象
- **函数定义**: 学习函数的声明和调用
- **面向对象**: 了解 Go 中的结构体和方法

### 并发编程
- **Goroutine**: 学习轻量级线程的使用
- **Channel**: 掌握通道的创建和使用
- **生产者消费者**: 理解并发编程的实际应用
- **发布订阅**: 学习事件驱动的编程模式

### 设计模式
- **单例模式**: 学习 Go 中单例模式的实现

### 算法练习
- **LeetCode 题目**: 包含经典算法题的 Go 实现

## 🛠️ 开发工具

推荐使用以下工具进行开发：
- **IDE**: GoLand, VS Code (安装 Go 扩展)
- **调试**: Delve
- **测试**: Go 内置测试框架

## 📝 注意事项

1. 所有示例代码都经过测试，可以直接运行
2. 建议按照文件名的顺序学习，从基础到进阶
3. 并发编程示例需要特别注意资源管理
4. 算法练习部分适合巩固 Go 语法

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个学习代码库！

## 📄 许可证

本项目采用 MIT 许可证。
