# pprof


## 运行
- 通过url查看， 文字界面 http://127.0.0.1:8000/debug/pprof/
- 通过url查看， 文字界面 
  - http://127.0.0.1:8000/debug/pprof/goroutine?debug=1
  - http://127.0.0.1:8000/debug/pprof/goroutine?debug=2
- 通过图形界面：go tool pprof -http=:8087 http://127.0.0.1:8000/debug/pprof/profile

### from
https://www.jianshu.com/p/7161a8e17da5