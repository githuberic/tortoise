## 概述
gin用jwt-go+crypto/bcrypt实现用户登录验证

### 访问验证
- 测试生成一个加密后的密码
curl -X GET 'http://127.0.0.1:8090/user/pass?password=123'

- 测试生成token
curl -X GET 'http://127.0.0.1:8090/user/login?username=lgq&password=123456'

- 测试使用获得的token访问api
curl -X GET 'http://127.0.0.1:8090/user/info?auth=Bearer%20eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxncSIsImV4cCI6MTY1NDcwNzA0NSwiaXNzIjoibGdxIn0.PZyZPv7JeMqDxJi4Sfkw4TXe8gY7om5Pc1l1_5uMYhI'

