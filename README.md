# golang-qiniu

golang for qiniu SDK demo

go语言 七牛云SDK云储存demo

- upload_Token 上传凭证

- download_Token 下载凭证

- bucket_manager 资源管理对象相关

- access_token 账户管理凭证和密钥

- bucket 新建空间名

- space 获取标准存储的当前存储量

- blob_id 获取外网流出流量统计和 GET 请求次数

- rs_put 获取 PUT 请求次数

> 调用次demo 需要你的拥有七牛云平台的 账号、密码、accessKey、secretKey、bucket(空间名)

> 另外 app 文件夹是我自己config配置，你可以用自己的封装的config来读取配置，也可以直接写固定内容。


每个文件名对应一个demo，每个demo都有单元测试，你可以直接运行测试（前提是你得自己换成自己的相关信息）。