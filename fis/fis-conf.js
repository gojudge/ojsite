fis.config.set('modules.parser.scss', 'sass');
fis.config.set('roadmap.ext.scss', 'css');

fis.config.set('modules.parser.coffee', 'coffee-script');
fis.config.set('roadmap.ext.coffee', 'js');

fis.config.set('roadmap.path',[
    {
        reg: /^\/sass\/(.*)/i,
        release: '/css/$1'
    },
    {
        reg: /^\/coffee\/(.*)/i,
        release: '/js/$1'
    }
]);

fis.config.set('pack', {
    'style.css': [
        'style.scss',
    ],
    'main.js':[
    	'main.coffee'
    ]
});

fis.config.merge({
    deploy : {
		local : {
            to : '../static'
        },
    }
});