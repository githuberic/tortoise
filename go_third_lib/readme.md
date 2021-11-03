# Network

## evio

Fast event-loop networking for Go

evio is an event loop networking framework that is fast and small. It makes direct epoll and kqueue syscalls rather than using the standard Go net package, and works in a similar manner as libuv and libevent.

https://github.com/tidwall/evio

## gnet

gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。

https://github.com/panjf2000/gnet

- 高性能 的基于多线程/Go 程网络模型的 event-loop 事件驱动
- 内置 goroutine 池，由开源库 ants 提供支持
- 内置 bytes 内存池，由开源库 bytebufferpool 提供支持
- 整个生命周期是无锁的
- 基于 Ring-Buffer 的高效内存利用
- 支持多种网络协议/IPC 机制： TCP 、 UDP 和 Unix Domain Socket
- 支持多种负载均衡算法： Round-Robin(轮询) 、 Source Addr Hash(源地址哈希) 和 Least-Connections(最少连接数)
- 支持两种事件驱动机制：Linux 里的 epoll 以及 FreeBSD/Darwin 里的 kqueue
- 支持异步写操作
- 灵活的事件定时器
- SO_REUSEPORT 端口重用
- 内置多种编解码器，支持对 TCP 数据流分包：LineBasedFrameCodec, DelimiterBasedFrameCodec, FixedLengthFrameCodec 和 LengthFieldBasedFrameCodec，参考自 netty codec，而且支持自定制编解码器

中文介绍 https://www.golangtc.com/t/5d7e89a4b17a820430406ee2






