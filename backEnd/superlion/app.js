const Koa = require('koa');
const app = new Koa();

// 定义一个路由
app.use(async (ctx) => {
    ctx.body = 'Hello, Koa!';
});

// 启动服务器
const port = process.env.PORT || 3000;
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});