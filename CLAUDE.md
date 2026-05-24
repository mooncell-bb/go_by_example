# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

这是一个个人的 Go 学习仓库，学习的网站为 [Go by Example](https://gobyexample.com/)。需要在学习此网站后，从零构建起完整的 Go 知识体系。

知识体系可分为：

- 基础语法与控制结构 (Basics) - 熟悉变量声明方式和基本的逻辑控制。
  - Hello World
  - Values
  - Variables
  - Constants
  - For
  - If/Else
  - Switch
- 数据结构与函数 (Data Structures & Functions) - 掌握如何组织数据和封装逻辑。
  - Arrays
  - Slices
  - Maps
  - Range over Built-in Types
  - Functions
  - Multiple Return Values
  - Variadic Functions
  - Closures
  - Recursion
  - Pointers
  - Strings and Runes
- 面向对象特性与泛型 (Types & Interfaces) - Go 没有传统的类和继承，而是通过结构体和接口实现代码组织。
  - Structs
  - Methods
  - Interfaces
  - Enums
  - Struct Embedding
  - Generics
- 错误处理 (Error Handling) - Go 提倡显式处理错误，而不是抛出异常。
  - Errors
  - Custom Errors
  - Panic
  - Defer
  - Recover
- 并发编程 (Concurrency) - goroutine 和 channel 是 Go 并发模型的核心。
  - Goroutines
  - Channels
  - Channel Buffering
  - Channel Synchronization
  - Channel Directions
  - Select
  - Timeouts
  - Non-Blocking Channel Operations
  - Closing Channels
  - Range over Channels
  - Timers
  - Tickers
  - Worker Pools
  - WaitGroups
  - Rate Limiting
  - Atomic Counters
  - Mutexes
  - Stateful Goroutines
- 标准库与常用工具 (Standard Library & Utilities) - 了解处理字符串、JSON、时间、文件操作等日常开发任务。
  - Sorting
  - Sorting by Functions
  - String Functions
  - String Formatting
  - Text Templates
  - Regular Expressions
  - JSON
  - XML
  - Time
  - Epoch
  - Time Formatting / Parsing
  - Random Numbers
  - Number Parsing
  - URL Parsing
  - SHA256 Hashes
  - Base64 Encoding
  - Reading Files
  - Writing Files
  - Line Filters
  - File Paths
  - Directories
  - Temporary Files and Directories
  - Embed Directive
  - Range over Iterators
  - Testing and Benchmarking
- 系统交互与网络编程 (System & Networking) - 编写服务器、调用 API 以及处理底层系统信号。
  - Command-Line Arguments
  - Command-Line Flags
  - Command-Line Subcommands
  - Environment Variables
  - Logging
  - HTTP Client
  - HTTP Server
  - TCP Server
  - Context
  - Spawning Processes
  - Exec'ing Processes
  - Signals
  - Exit

每个目录包含独立的 Go 源文件，演示特定的语言特性，对应的 `.md` 文件（如 `ch1 - basics.md`）包含详细的中文笔记和深入讲解。

## 学习目标

完成 "Go by Example" 的学习后，需要达到初级工程师水平，你需要对我给定章节的学习笔记进行知识点的补充、完善。

- 修改的风格请参考之前章节的笔记，要求做到清晰、简洁。
- 补充内容仅限本章节指定部分内容，不得额外补充在学习其它章节时的内容。例如学习 Interfaces 章节时，就无需补充 Struct Embedding 相关部分内容。

## 输出提示

每轮对话时，需在开头添加："章节内容" + "补充说明"。例如我需要让你完善 basics 的 Hello World 章节内容时，你应该在开头添加："basics - Hello World 补充说明"。
