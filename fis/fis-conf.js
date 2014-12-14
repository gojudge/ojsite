fis.config.set('modules.parser.scss', 'sass');
fis.config.set('roadmap.ext.scss', 'css');

fis.config.set('modules.parser.coffee', 'coffee-script');
fis.config.set('roadmap.ext.coffee', 'js');

fis.config.set('modules.parser.less', 'less');
fis.config.set('roadmap.ext.less', 'css');

fis.config.set('roadmap.path',[
    {
        reg: /^\/sass\/(.*)/i,
        release: '/static/css/$1'
    },
    {
        reg: /^\/js\/(.*)/i,
        release: '/static/js/$1'
    },
    {
        reg: /^\/octicons\/(.*)/i,
        release: '/static/octicons/$1'
    },
    {
        reg: /^\/fontello\/(.*)/i,
        release: '/static/fontello/$1'
    },
    {
        reg: /^\/codemirror\/(.*)/i,
        release: '/static/codemirror/$1'
    },
    {
        reg: /^\/img\/(.*)/i,
        release: '/static/img/$1'
    },
    {
        reg: /^\/ng\/(.*)/i,
        release: '/static/ng/$1'
    },
    {
        reg: /^\/ueditor\/(.*)/i,
        release: '/static/ueditor/$1'
    },
    {
        reg: /^\/epiceditor\/(.*)/i,
        release: '/static/epiceditor/$1'
    },
    {
        reg: /^\/js\/angular.min.js/i,
        release: '/static/js/angular.min.js'
    }
]);

fis.config.set('project.exclude', [
    /^\/(.*)\.md/i,
    /^\/(.*)\.zip/i,
    /^\/sass\/_(.*)\.scss/i,
    /^\/js\/user\/(.*)/i,
    /^\/js\/student\/(.*)/i,
    /^\/js\/taecher\/(.*)/i,
    /^\/js\/admin\/(.*)/i,
]);

fis.config.merge({
    deploy : {
		local : {
            to : '../'
        },
    }
});