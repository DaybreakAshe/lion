const router = require('koa-router')()

router.get('/', async (ctx, next) => {
  await ctx.render('index', {
    title: 'Hello Koa 3333,new code!'
  })
})

router.get('/string', async (ctx, next) => {
  ctx.body = 'koa2 string'
})

router.get('/json', async (ctx, next) => {
  ctx.body = {
    title: 'koa2 json'
  }
})

router.get('/articleList', async (ctx, next) => {
  ctx.body = [
    {
      title: "test",
      content: "Death investigated at Burning Man while 70,000 festival attendees remain stuck in Nevada desert after rain",
      image: 'xxx',
      date: "2021-10-10",
    }
  ]
})

module.exports = router
