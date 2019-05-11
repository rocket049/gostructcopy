# gostructcopy
func:
```
//CopyToStruct copy the exported value of a struct to a likely struct , with reflect.利用 reflect 技术把结构体的可 export 值复制到 dst 中
func StructCopy(src, dst interface{}) error
```

copy the exported value of a struct to a likely struct , with reflect. dst must be pointer to s likely struct.

利用 reflect 技术把结构体的可 export 值复制到 dst 中，dst 必须是相似结构体的指针。
