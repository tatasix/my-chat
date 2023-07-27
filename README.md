# my-wechat

此项目分成几个大的功能服务模块
- **chat服务**
使用[chatgpt](https://github.com/chy4pro/chatgpt-wechat)项目，做了很多个性化开发
可在微信 **安全使用（通过企业微信中转到微信，无封号风险）** 的 ChatGPT 个人助手应用,
- **cron服务**
定时任务服务
- **server服务**
管理后台服务 
使用[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)项目 
- **script服务**
脚本执行服务

## 主要能力
### chat
- 微信可用：基于企业微信中转，可在微信中安全使用
- 客服消息：[支持多渠道客服消息接入](./doc/custom_support_service.md)
- 代理支持： `http/sock5` 代理 && 反向域名代理支持, 除了 `openai` 也兼容了 `azure-openai`
- 余额查询： `openai` 余额查询
- 会话：
  - 场景模式：支持动态修改 `prompt`，预定义了上百种 `prompt` 角色模板
  - 连续对话：自适应的上下文设计，让 LLM 🧠拥有更长时间的短期记忆，避免手动清理上下文
  - 会话切换：多会话支持，可切换聊天场景,不丢失对话内容
  - 极速响应：支持基于流式接口的分段消息响应
- 消息：
  - 语音消息：支持多国语音消息，以及多语言文本输入
  - 图片消息：支持图片消息（答题场景）
- 绘画：
  - 支持`stable diffusion 1.5`作图
- 私有数据：`milvus` 私有化向量知识库支持
- 插件机制：
  - 目前已支持 日期、`search`、`wikipedia`
  - 各位也可按照规则自行开发接入其他能力（summary, 天气 ...）

### server
- go的插件中心，现已支持 ：微信支付、登录等，K8s相关操作 ，第三方登录 等等插件
- 权限管理：基于jwt和casbin实现的权限管理
- 文件上传下载：实现基于七牛云的文件上传操作（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改（测试环境不开放此功能）。
- 富文本编辑器：MarkDown编辑器功能嵌入。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。

### 修改 `chat-api.yaml` 配置后如何让其生效
<details>
<summary></summary>

- 你可以通过 `docker-compose restart web` 重启 web 服务
- 或者 `docker-compose build && docker-compose up -d` 重启整个服务
</details>

### 如何修改 redis 密码? 
<details>
<summary></summary>

- 首先修改 `chat/service/chat/api/etc/chat-api.yaml` 
```yaml
RedisCache:
    Pass: "xxxxxx"
```
- 再修改 `chat/build/redis/redis.conf`
```ini
requirepass "xxxxx"
```
- 最后 `docker-compose down && docker-compose up -d` 重启整个服务
</details>
