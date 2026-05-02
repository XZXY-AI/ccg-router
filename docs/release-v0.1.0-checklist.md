# v0.1.0 发布清单

目标：发布第一个可安装、可验证、可传播的 public release。

## 当前状态

- 主仓库：`XZXY-AI/ccg-router`
- Homebrew tap：`XZXY-AI/homebrew-tap`
- 默认分支：`main`
- 发布流水线：`.github/workflows/release.yml`
- Release 产物：Linux/macOS，amd64/arm64，tar.gz，checksums，cosign keyless signature

## 发布前必须完成

1. 确认 `main` 最新 CI 通过。
2. 设置 `HOMEBREW_TAP_TOKEN`：
   - 创建 fine-grained GitHub token。
   - 只授予 `XZXY-AI/homebrew-tap` 的 Contents read/write 权限。
   - 写入主仓库 secret：

   ```bash
   gh secret set HOMEBREW_TAP_TOKEN --repo XZXY-AI/ccg-router
   ```

3. 本地验证：

   ```bash
   make lint
   make test
   make build
   make vuln
   ```

4. 检查公开文档没有旧组织名、真实密钥或发布占位。下面的正则用字符类避开清单文本自我命中：

   ```bash
   rg -n --hidden 'c[c]g-labs|PUBLISH[-]BEFORE|github[_]pat_|g[h]p_|s[k]-' .
   ```

## 发布命令

```bash
git fetch origin
git checkout main
git pull --ff-only
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

Tag push 会触发 GitHub Actions release workflow。

## 发布后验证

1. Release workflow 成功。
2. GitHub Release 出现四个平台 archive、`checksums.txt` 和 `.sig` 文件。
3. Homebrew tap 出现 `Formula/ccg-router.rb`。
4. 安装验证：

   ```bash
   go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
   curl -fsSL https://raw.githubusercontent.com/XZXY-AI/ccg-router/main/scripts/install.sh | bash
   brew install XZXY-AI/tap/ccg-router
   ccg-router doctor
   ```

5. README Quickstart 可把 release binary、shell installer、Homebrew 恢复为主要安装方式。

## 回滚

如果 release workflow 失败：

1. 保留失败日志。
2. 删除失败 release 草稿或 release 页面。
3. 如 tag 已发布且产物不可用，删除远端 tag：

   ```bash
   git push origin :refs/tags/v0.1.0
   git tag -d v0.1.0
   ```

4. 修复后重新打同名 tag。
