# 📋 内部文件收集与管理工具（截图收集系统）

## 项目简介

这是一个用于收集和管理员工转发公众号/视频号内容截图的内网工具系统。系统实现了截图文件的智能化收集、规范化存储、便捷式上传和实时化统计。

## 技术栈

- **前端**: Vue 3 + TypeScript + Vite + Element Plus + Pinia + Vue Router
- **后端**: Go + Gin + GORM + MySQL

## 项目结构

```
collect-tools/
├── server/          # Go后端服务
│   ├── config/      # 配置管理
│   ├── database/    # 数据库初始化
│   ├── models/      # 数据模型
│   ├── router/      # 路由定义
│   └── main.go      # 入口文件
└── web/             # Vue前端应用
    ├── src/
    │   ├── components/  # 组件
    │   ├── layouts/     # 布局
    │   ├── router/      # 路由
    │   ├── store/       # Pinia状态管理
    │   ├── utils/       # 工具函数
    │   └── views/       # 页面视图
    └── ...
```

## 快速开始

### 前置要求

- Go 1.25.4+
- Node.js 18+
- pnpm
- MySQL 5.7+

### 后端启动

1. 进入server目录：
```bash
cd server
```

2. 安装依赖：
```bash
go mod download
```

3. 配置数据库：
   - 创建MySQL数据库
   - 复制 `.env.example` 为 `.env` 并修改配置

4. 启动服务：
```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动

### 前端启动

1. 进入web目录：
```bash
cd web
```

2. 安装依赖：
```bash
pnpm install
```

3. 启动开发服务器：
```bash
pnpm dev
```

前端将在 `http://localhost:5173` 启动

## 功能模块

- ✅ 部门管理
- ✅ 人员管理
- ✅ 转发任务管理
- ✅ 文件上传（支持剪贴板、拖拽、选择文件）
- ✅ 实时统计看板
- ✅ 文件导出（ZIP压缩）
- ✅ 报表导出（Excel）

## 开发进度

- [x] 项目初始化和基础架构
- [x] 路由和布局组件
- [x] 数据模型定义
- [ ] 部门管理功能
- [ ] 人员管理功能
- [ ] 转发任务管理
- [ ] 文件上传功能
- [ ] 统计看板
- [ ] 文件导出功能

## 许可证

内部使用

