export default {
    npmClient: 'yarn',
    plugins: ['@umijs/plugins/dist/antd'],
    antd: {},
    routes: [
        {path: '/', redirect: '/login'},
        {path: '/login', component: 'login'},
        {path: '/register', component: 'register'},
        {path: '/', component: 'layout',
            routes: [
                {path: '/account', component: 'account'}
            ]
        },
    ],
    "theme": {
        "@primary-color": "#FA541C",
    },
    proxy: {
        '/api': {
            target: 'http://127.0.0.1:8989',
            changeOrigin: true,
            pathRewrite: {'^/': ''},
        },
    }
};
