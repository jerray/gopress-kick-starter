# gopress-kick-starter

gopress kick starter

## Web项目目录结构

以下是使用该框架时推荐的项目目录结构。

```
root
  |- models
       |- user.go
       |- post.go
  |- controllers
       |- users.go
       |- posts.go
  |- middlewares
       |- auth.go
  |- services
       |- database.go
  |- views
       |- _partials
            |- header.handlebars
            |- footer.handlebars
            |- users
                 |- avatar.handlebars
       |- users
            |- login.handlebars
            |- profile.handlebars
  |- main.go
```

以上目录结构除`views`目录必须包含`_partials`外，其他路径均无严格要求。
Handlebars模板是根据使用情况动态加载到内存中的，但全局Partials是在程序启动时加载的。
所以partials文件必须放在模板路径下的 `_partials` 路径内。

如果项目中希望包含多个command，建议使用 [cobra](https://github.com/spf13/cobra) 代替main.go做为启动器。
同时使用 [viper](https://github.com/spf13/cobra) 来加载配置文件。
