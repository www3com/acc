export default {
  npmClient: 'yarn',
  publicPath: '/web/',
  hash: true,
  routes: [
    {path: '/', redirect: '/login'},
    {path: '/login', component: 'login'},
    {path: '/register', component: 'register'},
    {path: '/', component: 'layout',
      routes: [
          {path: '/total', component: 'account'},
          {path: '/bill', component: 'bill'},
          {path: '/account', component: 'account'},
          {path: '/report', component: 'account'},
          {path: '/settings', component: 'settings'}
      ]
    },
  ],
  proxy: {
    '/api': {
      target: 'http://127.0.0.1:8989',
      changeOrigin: true,
      pathRewrite: { '^/': '' },
    },
  },
};
