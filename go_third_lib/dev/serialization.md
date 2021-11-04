# 序列化/反序列化

## jsoniter

jsoniter可以解决这个痛点，其是一款快且灵活的 JSON 解析器，具有良好的性能并能100%兼容标准库，我们可以使用jsoniter替代encoding/json，官方文档称可以比标准库快6倍多，后来Go官方在go1.12版本对 json.Unmarshal 函数使用 sync.Pool 缓存了 decoder，性能较之前的版本有所提升，所以现在达不到快6倍多。

github地址：https://github.com/json-iterator/go