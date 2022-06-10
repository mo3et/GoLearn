// 位运算
// %b 是用于表示位的格式化标识符

// 二元运算符

// 按位与 &

// 两者都为1 为T 否则为0

//   1 & 1 -> 1
//   1 & 0 -> 0
//   0 & 1 -> 0
//   0 & 0 -> 0

// 按位或 ｜

// 有一者为1 为T

// 1 | 1 -> 1
// 1 | 0 -> 1
// 0 | 1 -> 1
// 0 | 0 -> 0

// 按位异或 ^

// 两者相同为 0 不同为1

// 1 ^ 1 -> 0
// 1 ^ 0 -> 1
// 0 ^ 1 -> 1
// 0 ^ 0 -> 0

// 位清除 &^：将指定位置上的值设置为 0

// 一元运算符

// 按位补足 ^

// 位左移 <<
// 用法： bitP << n
// bitP 的位向左移动n位,右侧空白部分使用0填充; 如果n等于2,则结果是2的相应倍数,即2的n次方.

// 1 << 10 // 等于 1 KB
// 1 << 20 // 等于 1 MB
// 1 << 30 // 等于 1 GB

// 位右移 >> :
// 用法：bitP >> n。
// bitP 的位向右移动 n 位，左侧空白部分使用 0 填充；如果 n 等于 2，则结果是当前值除以 2 的 n 次方。

// 当希望把结果赋值给第一个操作数时，可以简写为 a <<= 2 或者 b ^= a & 0xfffffff

// Usually example

// 使用位左移与iota计数实现存储单位的常量枚举：

