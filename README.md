基于 `gin` + `Angular` 实现的 `markdown` 静态博客

# 第三方库

**后端**

* gin

**前端**

* Angular Material
* Angular Flex
* ngx-markdown

# 基本流程

* 后端把 markdown 发送给前端，前端渲染 markdown
* 前端不在内存中储存 markdown，需要查询时直接向后端发送请求

# 前端

## Component

Component|作用
---|---
Home|主页
Article|文章详情
Navbar|导航栏

# 后端

## 目录结构

```console
-backend
│  go.mod
│  go.sum
│  main.go
│
├─handlers
│      ArticleHandler.go
│
├─markdown
│      HelloWorld-2021.11.02-oneman233.md
│      README-2021.11.03-oneman233.md
│
├─middlewares
│      CORS.go
│
├─models
│      article.go
│
├─pics
├─routers
│      routers.go
│
└─utils
        const.go
        getmarkdown.go
```

## RESTful

Url|Method|作用
---|---|---
/v1/articles|GET|获取文章列表
/v1/article/:id|GET|根据 ID 获取某篇文章

# TODOs

- [ ] 前端错误处理以及 404 页面设计
- [ ] About页面设计
- [ ] 评论模块开发

# 效果图

![homepage](/results/homepage.png)
![articlepage](/results/articlepage.png)