# MCP Host Web [第四课]

把 [MCP Host](https://github.com/guobinqiu/mcp-host) (命令行版) 改成 Web 版

## 运行

1. 启动后端服务

```
cd backend && go run main.go
```

2. 启动前端服务

```
cd frontend && npm run serve
```

## 效果图

<img width="363" alt="image" src="https://github.com/user-attachments/assets/c7cd91dd-bf51-4223-9d7f-d0c2dca1d381" />

## TODO

没有MCP的时候流式响应到前端很简单,但是带了MCP之后怎么做流式响应?因为llm返回tools也是流式的不好控制.
