# dxgi

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

DXGI (DirectX Graphics Infrastructure) 的 Go 语言绑定库，提供了完整的 Windows DXGI API 接口封装。

## 简介

`dxgi` 是一个自用的 DXGI 绑定库，通过 Go 语言的 `syscall` 包直接调用 Windows DXGI DLL，实现了从 DXGI 1.0 到 DXGI 1.6 的主要接口。该库采用 COM 接口绑定方式，通过虚函数表（vtbl）实现接口调用。

## 功能特性

- ✅ 完整的 DXGI 接口支持（1.0 - 1.6）
- ✅ 工厂接口（IDXGIFactory 到 IDXGIFactory7）
- ✅ 适配器管理（IDXGIAdapter 到 IDXGIAdapter4）
- ✅ 交换链管理（IDXGISwapChain 到 IDXGISwapChain4）
- ✅ 输出管理（IDXGIOutput 到 IDXGIOutput6）
- ✅ 资源管理（IDXGIResource, IDXGIResource1）
- ✅ 表面管理（IDXGISurface 到 IDXGISurface2）
- ✅ 设备接口（IDXGIDevice 到 IDXGIDevice4）
- ✅ 输出复制（IDXGIOutputDuplication）
- ✅ 调试接口支持
- ✅ 类型安全的 Go API 封装

## 系统要求

- Windows 操作系统
- Go 1.16 或更高版本
- DirectX 运行时（通常随 Windows 安装）

## 安装

```bash
go get github.com/liuqi0826/dxgi
```

## 快速开始

### 创建 DXGI 工厂

```go
package main

import (
    "fmt"
    "github.com/liuqi0826/dxgi"
)

func main() {
    // 创建 DXGI 工厂
    factory, err := dxgi.CreateDXGIFactory1(&dxgi.IID_IDXGIFactory1)
    if err != nil {
        fmt.Printf("创建工厂失败: %v\n", err)
        return
    }
    defer factory.Release()

    // 枚举适配器
    for i := uint32(0); ; i++ {
        adapter, err := factory.EnumAdapters1(i)
        if err != nil {
            break
        }
        defer adapter.Release()

        // 获取适配器描述
        desc, err := adapter.GetDesc1()
        if err != nil {
            continue
        }

        fmt.Printf("适配器 %d: %s\n", i, desc.Description)
    }
}
```

### 枚举显示输出

```go
// 获取第一个适配器
adapter, err := factory.EnumAdapters1(0)
if err != nil {
    return
}
defer adapter.Release()

// 枚举输出
for i := uint32(0); ; i++ {
    output, err := adapter.EnumOutputs(i)
    if err != nil {
        break
    }
    defer output.Release()

    // 获取输出描述
    desc, err := output.GetDesc()
    if err != nil {
        continue
    }

    fmt.Printf("输出 %d: %s\n", i, desc.DeviceName)
}
```

## 支持的接口

### 核心接口

- **IDXGIFactory** - 工厂接口（1-7）
- **IDXGIAdapter** - 适配器接口（1-4）
- **IDXGIOutput** - 输出接口（1-6）
- **IDXGISwapChain** - 交换链接口（1-4）
- **IDXGISurface** - 表面接口（1-2）
- **IDXGIResource** - 资源接口（1）
- **IDXGIDevice** - 设备接口（1-4）

### 扩展接口

- **IDXGIObject** - 对象基接口
- **IDXGIDeviceSubObject** - 设备子对象接口
- **IDXGIKeyedMutex** - 键控互斥锁接口
- **IDXGIOutputDuplication** - 输出复制接口
- **IDXGISwapChainMedia** - 交换链媒体接口
- **IDXGIDecodeSwapChain** - 解码交换链接口
- **IDXGIDisplayControl** - 显示控制接口

## 主要功能

### 适配器管理
- 枚举图形适配器
- 查询适配器信息（描述、内存、标志等）
- 检查接口支持
- 视频内存管理

### 交换链管理
- 创建和管理交换链
- 全屏模式支持
- 缓冲区管理
- 帧统计信息
- HDR 元数据支持

### 输出管理
- 枚举显示输出
- 获取显示模式列表
- 查找匹配模式
- 伽马控制
- 输出复制（桌面捕获）

### 资源管理
- 共享句柄管理
- 资源使用标志
- 驱逐优先级
- 子资源表面创建

## 类型定义

库提供了完整的类型定义，包括：

- `DXGI_ADAPTER_DESC` / `DXGI_ADAPTER_DESC1-3` - 适配器描述
- `DXGI_OUTPUT_DESC` / `DXGI_OUTPUT_DESC1` - 输出描述
- `DXGI_SWAP_CHAIN_DESC` / `DXGI_SWAP_CHAIN_DESC1` - 交换链描述
- `DXGI_MODE_DESC` / `DXGI_MODE_DESC1` - 模式描述
- `DXGI_FRAME_STATISTICS` - 帧统计信息
- `DXGI_QUERY_VIDEO_MEMORY_INFO` - 视频内存信息
- 以及更多...

## 枚举和常量

库定义了所有 DXGI 相关的枚举和常量：

- `DXGI_ADAPTER_FLAG` - 适配器标志
- `DXGI_SWAP_EFFECT` - 交换效果
- `DXGI_SWAP_CHAIN_FLAG` - 交换链标志
- `DXGI_MODE_SCANLINE_ORDER` - 扫描线顺序
- `DXGI_MODE_ROTATION` - 旋转模式
- `DXGI_COLOR_SPACE_TYPE` - 颜色空间类型
- 以及更多...

## 错误处理

所有方法都返回标准的 Go `error` 类型。错误码通过 `GetError()` 函数从 HRESULT 转换为相应的错误信息。

```go
factory, err := dxgi.CreateDXGIFactory1(&dxgi.IID_IDXGIFactory1)
if err != nil {
    // 处理错误
    fmt.Printf("错误: %v\n", err)
}
```

## 注意事项

1. **资源管理**: 所有 COM 对象都需要调用 `Release()` 方法释放资源
2. **线程安全**: DXGI 接口不是线程安全的，需要在同一线程中使用
3. **错误处理**: 始终检查方法返回的错误
4. **指针传递**: 某些方法需要传递指针参数，注意使用 `&` 操作符

## 示例

更多示例请参考项目中的测试文件或文档。

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 相关链接

- [Microsoft DXGI 文档](https://docs.microsoft.com/en-us/windows/win32/direct3ddxgi/dx-graphics-dxgi)
- [DirectX Graphics Infrastructure](https://en.wikipedia.org/wiki/DirectX_Graphics_Infrastructure)

## 作者

liuqi0826

---

**注意**: 这是一个自用的绑定库，主要用于个人项目。使用前请确保理解 DXGI API 的工作原理和 Windows COM 编程规范。
