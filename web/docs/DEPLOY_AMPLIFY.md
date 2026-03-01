# 将 Web 前端部署到 AWS Amplify

本文说明如何把本项目的 **web** 前端部署到 AWS Amplify Hosting。

## 一、前置条件

- 代码在 Git 仓库中（GitHub、GitLab、Bitbucket、AWS CodeCommit 等），Amplify 需连接该仓库
- 生产环境后端 API 已部署并可访问（用于配置 `VITE_SERVICE_BASE_URL`）

## 二、在 Amplify 控制台创建应用

1. 打开 [AWS Amplify 控制台](https://console.aws.amazon.com/amplify/)
2. 点击 **Create new app** → **Host web app**
3. 选择你的 Git 提供商和仓库，选择分支（如 `main`）
4. **必须设置根目录为 `web`**：在 “Build settings” / “Monorepo” 中，将 **Root directory**（或 **App root**）设为 **`web`**。这样构建时工作目录就是前端项目本身，`amplify.yml` 中的命令才能正确执行（无需再 `cd web`）。

## 三、构建配置

项目根目录的 **`amplify.yml`** 会在 **Root directory = web** 的前提下执行：

- **preBuild**：启用 pnpm（`corepack enable`）并安装依赖（`pnpm install`）
- **build**：执行 `pnpm run build`（即 `vite build --mode prod`）
- **artifacts**：发布目录为 `dist`（相对于 `web`）
- **cache**：缓存 `node_modules` 以加速后续构建

如需在控制台覆盖或查看，可在 Amplify 应用 → **Build settings** 中编辑，但建议以仓库中的 `amplify.yml` 为准。

## 四、环境变量（必配）

在 Amplify 应用 → **Environment variables** 中为**生产环境**添加：

| 变量名 | 说明 | 示例 |
|--------|------|------|
| `VITE_SERVICE_BASE_URL` | 生产环境后端 API 根地址（前端请求会发到这里） | `https://api.yourdomain.com` |

构建时 Vite 会把以 `VITE_` 开头的变量内联到前端代码，因此**修改后需要重新构建**才会生效。

其他可选变量（若未设置则使用 `web/.env.prod` 或 `web/.env` 中的默认值）：

- `VITE_APP_TITLE`、`VITE_APP_DESC` 等可根据需要配置

## 五、Node 版本

- 仓库中已包含 **`web/.nvmrc`**（内容为 `18`），若 Amplify 使用 `.nvmrc`，将自动使用 Node 18
- 若 Amplify 未自动识别，请在 **Build settings** → **Build image settings** 中选择 **Node 18** 或更高版本（与 `package.json` 中 `engines.node` 一致）

## 六、SPA 路由（History 模式）

前端使用 Vue Router **history** 模式，所有路径应由同一份 `index.html` 处理。在 Amplify 中需配置“重写”规则，把除静态资源外的请求都指向 `index.html`：

1. Amplify 应用 → **Hosting** → **Redirects and rewrites**
2. 添加一条 **Rewrite** 规则：
   - **Source address**:  
     `</^[^.]+$|\.(?!(css|gif|ico|jpg|js|png|txt|svg|woff|woff2|ttf|map|json)$)([^.]+$)/>`
   - **Target address**: `/index.html`
   - **Type**: **200 (Rewrite)**

这样刷新任意前端路由或直接访问子路径时，都会返回 `index.html`，由前端路由接管。

## 七、部署流程

1. 保存上述环境变量和重写规则
2. 在 Amplify 中 **Save** 并 **Redeploy this version**，或推送代码触发自动构建
3. 构建成功后，Amplify 会使用 `dist` 的内容提供静态站点，并给出默认域名（如 `https://main.xxx.amplifyapp.com`）
4. 如需自定义域名，在 **Hosting** → **Domain management** 中绑定并按提示完成 DNS/SSL 配置

## 八、常见问题

- **`cd: web: No such file or directory`**  
  说明构建时工作目录里没有 `web` 子目录。请在 Amplify 的 **Build settings** 中将 **Root directory** 设为 **`web`**，这样构建会在 `web` 目录下执行，无需再 `cd web`。

- **构建失败：pnpm 未找到**  
  确保 `amplify.yml` 中 preBuild 包含 `corepack enable`（当前配置已包含）。

- **页面空白或接口 404**  
  检查 `VITE_SERVICE_BASE_URL` 是否与真实后端地址一致，且后端 CORS 允许该 Amplify 域名。

- **刷新子路径 404**  
  说明 SPA 重写未生效，请按第六步检查并添加 Rewrite 规则。

- **需要区分多环境**  
  可在 Amplify 为不同分支配置不同环境变量（如 `main` 用生产 API，`develop` 用测试 API），并确保构建命令使用对应 mode（当前为 `prod`）。
