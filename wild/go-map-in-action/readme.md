# Go map in action

> Link: https://go.dev/blog/maps

## introduction

Hash 表 是最有用的数据结构之一，他们提供了**快速查找**、**添加**和**删除**的功能，Go 提供了一种实现 hash table 的内置 `map` 类型。

## 声明和初始化：
Go 的 `map` 类型就像这样：
```
map[KeyType]ValueType
```
其中 KeyType 是任何可**比较的**类型，