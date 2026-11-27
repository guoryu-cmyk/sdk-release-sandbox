# sdk-release-sandbox

用于联调 SDK 发布控制台、Codebase CI 和 GitHub Actions 的 Go SDK 测试工程。外发目标固定为 `guoryu-cmyk/sdk-release-sandbox`。

The sandbox also supports a manual release smoke test to verify the complete release workflow.

| Codebase 流水线 | 职责 |
| --- | --- |
| `update_version_and_changelog` | 从平台固定的 `master` SHA 生成版本号和 Changelog。测试和构建通过后推送 `version_*`。控制台根据返回的分支 SHA 创建版本 MR。 |
| `sync_master_to_mainrelease` | 版本 MR 合入后过滤内网文件，对待发布的代码重新测试和构建。通过后同步到 Codebase 的 `mainrelease`，记录验证通过的提交。 |
| `publish_mainrelease_to_github` | 到发布窗口后，把 `mainrelease` 上已验证的提交推到 GitHub 发布分支，确认发布工作流启用后，再给同一个提交推 tag。 |

三条流水线支持两种模式：

- **控制台触发**：传入控制台签发的 `operation_id`，其他手动参数不参与执行。在 Codebase Pipeline Variables 配置 `SDK_RELEASE_CONSOLE_URL` 和 `SDK_RELEASE_CI_TOKEN`。CI 用 Bearer 凭据向控制台核实当前操作，写分支或 tag 前再次核实，执行后回传结果。无效操作 ID 不会自动切换独立模式。
- **Codebase 独立手动触发**：`operation_id` 留空，不请求发布控制台，也不回传控制台任务状态。发布审批和 MR 合入由操作者处理，仍执行代码校验、测试构建、文件过滤和 Git 引用冲突检查。

| 流水线 | 独立运行时填写 |
| --- | --- |
| `update_version_and_changelog` | `bump`：patch / minor / major；可选 `base_sha`（默认当前 master）、`baseline_sha`（自动探测不适用时）、`release_date`（默认北京时间当天）。根据日志中 `Standalone result` 的版本分支创建并合入版本 MR。 |
| `sync_master_to_mainrelease` | `merged_sha`：已合入 master 的版本 MR 完整提交 SHA，从该提交读取版本。记录日志输出的 `release_sha`。 |
| `publish_mainrelease_to_github` | `release_sha`：上一步输出的完整 SHA（须仍为 mainrelease 头部），从该提交读取版本和 tag。 |

GitHub 发布任务在两种模式下都需要 `GH_PUSH_TOKEN`。重试时使用相同参数；跨日期重试版本准备时显式填写原来的 `release_date`。

内网测试的 `SDK_RELEASE_CONSOLE_URL` 可使用控制台独立回调监听的 HTTP 根地址。HTTP 只允许直连内网 IP 或回环地址，其他目标使用 HTTPS。

GitHub Actions 在 tag 推送后校验对应提交，运行测试和构建，然后通过 GitHub API 创建测试 Release。手动触发仅校验已有 tag。CI 使用 `GH_PUSH_TOKEN` 查询 `release.yml` 并推送 Git 引用。凭据需要能读取 Actions 工作流，以及更新目标仓库和工作流文件。工作流等待超时后可以沿用同一发布提交重试。公开 Git 历史必须能安全承接当前 `mainrelease`，分支或 tag 冲突时流水线会停止。 [GitHub 工作流查询接口](https://docs.github.com/en/rest/actions/workflows#get-a-workflow) · [GitHub 触发规则](https://docs.github.com/en/actions/how-tos/write-workflows/choose-when-workflows-run/trigger-a-workflow)

本地联调使用临时 Git 仓库和模拟控制台，不需要配置平台地址或凭据。

```bash
PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover -s .codebase/release -v
go test ./...
go build ./...
```
