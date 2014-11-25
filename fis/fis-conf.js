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
        reg: /^\/coffee\/(.*)/i,
        release: '/static/js/$1'
    },
    {
        reg: /^\/octicons\/(.*)/i,
        release: '/static/octicons/$1'
    }
]);

fis.config.set('project.exclude', [
    /^\/(.*)\.md/i,
]);

fis.config.set('pack', {
    'style.css': [
        'style.scss',
    ],
    'main.js':[
    	'main.coffee'
    ],
	'octicons.css':[
    	'octicons.less'
    ]
});

fis.config.merge({
    deploy : {
		local : {
            to : '../'
        },
    }
});