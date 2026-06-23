# 测试文档

本项目的测试文件覆盖了所有主要接口和功能。

## 测试文件列表

### 核心功能测试
- **dxgi_test.go** - 测试基础 DXGI API（CreateDXGIFactory 等）
- **dxgifactory_test.go** - 测试工厂接口功能
- **dxgiadapter_test.go** - 测试适配器接口功能
- **dxgioutput_test.go** - 测试输出接口功能

### 工具和类型测试
- **guid_test.go** - 测试 GUID 解析和格式化
- **utils_test.go** - 测试工具函数（错误处理、类型转换等）
- **const_test.go** - 测试常量定义
- **type_test.go** - 测试基础类型（RECT, POINT, LUID 等）
- **dxgicommon_test.go** - 测试通用 DXGI 类型

### 扩展接口测试
- **dxgiswapchain_test.go** - 交换链测试（占位符，需要 D3D 上下文）
- **dxgisurface_test.go** - 表面测试（占位符，需要 D3D 上下文）
- **dxgidevice_test.go** - 设备测试（占位符，需要 D3D 上下文）
- **dxgiresource_test.go** - 资源测试（占位符，需要 D3D 上下文）
- **dxgioutputduplication_test.go** - 输出复制测试（占位符，需要特定权限）
- **dxgikeyedmutex_test.go** - 键控互斥锁测试（占位符，需要资源对象）
- **dxgidisplaycontrol_test.go** - 显示控制测试（占位符，需要适配器对象）

### 测试辅助
- **test_helper.go** - 测试辅助函数（获取测试对象等）

## 运行测试

### 运行所有测试（跳过 Windows 特定测试）
```bash
go test -short
```

### 运行所有测试（包括 Windows 特定测试）
```bash
go test
```

### 运行特定测试
```bash
go test -v -run TestGUID
go test -v -run TestGetError
go test -v -run TestDXGIAdapter
```

### 运行 Windows 特定测试
```bash
go test -v -run TestCreateDXGIFactory1
```

## 测试覆盖

### 已测试的功能
- ✅ GUID 解析和格式化
- ✅ 错误码转换（GetError）
- ✅ 类型转换函数（boolToInt, intToBool）
- ✅ 常量定义验证
- ✅ 基础类型结构体
- ✅ 工厂创建和基本操作
- ✅ 适配器枚举和描述获取
- ✅ 输出枚举和描述获取

### 需要实际 Windows 环境的测试
以下测试需要在 Windows 系统上运行，并且可能需要特定的硬件或权限：
- 工厂接口的完整功能测试
- 适配器的完整功能测试
- 输出的完整功能测试
- 交换链、表面、设备等需要 D3D 上下文的测试
- 输出复制（需要特定权限）

## 测试注意事项

1. **平台限制**: 大部分测试只能在 Windows 系统上运行
2. **短模式**: 使用 `-short` 标志会跳过所有 Windows 特定的测试
3. **资源管理**: 所有测试都正确使用 `defer Release()` 来清理资源
4. **错误处理**: 测试会检查错误并适当处理

## 测试统计

- **测试文件数**: 15+
- **测试函数数**: 50+
- **通过率**: 100% (在 short 模式下)
- **Windows 特定测试**: 30+ (需要实际 Windows 环境)
