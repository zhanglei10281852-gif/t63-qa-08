# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

就绪检查的上游请求已经取消，但依赖检查仍持续到自身的长超时，服务关闭因此被拖延。请先不要修改代码，定位父级取消信号在哪里丢失，并提供可重复证据。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/t63-qa-08
- 仓库地址：https://github.com/zhanglei10281852-gif/t63-qa-08.git
- parent SHA：924cbbba73d2ace34a675f655a4e4dbf5ded02fc

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/t63-qa-08.git bug-repro
cd bug-repro
git checkout --detach 924cbbba73d2ace34a675f655a4e4dbf5ded02fc
go test ./internal/health -run TestCheckerStopsWhenParentRequestIsCancelled -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/health -run TestCheckerStopsWhenParentRequestIsCancelled -count=1
--- FAIL: TestCheckerStopsWhenParentRequestIsCancelled (0.50s)
    parent_cancellation_test.go:26: health checker ignored parent cancellation
FAIL
FAIL	sanitation-operations/internal/health	0.503s
FAIL

```

stderr：

```text
warning: internal/health/parent_cancellation_test.go has type 100755, expected 100644
warning: internal/health/parent_cancellation_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/health -run TestCheckerStopsWhenParentRequestIsCancelled -count=1
--- FAIL: TestCheckerStopsWhenParentRequestIsCancelled (0.52s)
    parent_cancellation_test.go:26: health checker ignored parent cancellation
FAIL
FAIL	sanitation-operations/internal/health	0.643s
FAIL

```

stderr：

```text
warning: internal/health/parent_cancellation_test.go has type 100755, expected 100644
warning: internal/health/parent_cancellation_test.go has type 100755, expected 100644

```

## 通过条件

在触发条件下，定向测试 TestCheckerStopsWhenParentRequestIsCancelled 应通过，相关包、全量测试、竞态测试和构建检查均通过；回退 gold 唯一修复后定向测试重新失败。
