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
    }
]);

fis.config.set('project.exclude', [
    /^\/(.*)\.md/i,
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