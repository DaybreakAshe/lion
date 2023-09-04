var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express11' });
});

router.get('/test', function (req, res, next) {
  res.render('index', { title: 'test data###' });
});

module.exports = router;
