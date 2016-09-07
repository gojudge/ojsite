<!-- <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="renderer" content="webkit|ie-stand|ie-comp" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <meta name="description" content="">
    <meta name="author" content="">
    <title>{{{.title}}}</title>
    {{{asset "sass/style.scss"}}}
    {{{asset "octicons/octicons.scss"}}}
    {{{asset "fontello/css/fontello.css"}}}
    <script src="//cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
    {{{asset "js/global.js"}}}
    <link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="bookmark" type="image/x-icon" />
    <link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="icon" type="image/x-icon" />
    <link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="shortcut icon" type="image/x-icon" />
  </head>
  <body>
    <div class="wrapper">

      <noscript>Please enable JavaScript in your browser!</noscript>

      <div class="header">
      	<ul>
      		<li><a href="/"><i class="logo icon-goj"></i></a></li>
      		<li><a href="/problems">{{{i18n "problems"}}}</a></li>
          <li><a href="/teach">{{{i18n "teach"}}}</a></li>
          <li><a href="">{{{i18n "discuss"}}}</a></li>
          <li>
            <form action="/search" method="get" target="_blank">
              <input type="text" name="keyword" id="" placeholder="Search">
            </form>
          </li>

{{{if eq .userIs "guest"}}}
          <li class="right"><a href="/register">{{{i18n "register"}}}</a></li>
          <li class="right"><a href="/login"><span title="sign in" class="octicon octicon-sign-in"></span> {{{i18n "login"}}}</a></li>
{{{else}}}
          <li class="right"><a href="/logout"><span title="sign out" class="octicon octicon-sign-out"></span></a></li>
  {{{if eq .userIs "admin"}}}
          <li class="right"><a href=""><span title="administration" class="octicon octicon-circuit-board"></span></a></li>
  {{{end}}}
  {{{if eq .userIs "teacher"}}}
          <li class="right"><a href="/teacher"><span title="teacher" class="octicon octicon-person"></span></a></li>
  {{{end}}}
  {{{if eq .userIs "student"}}}
          <li class="right"><a href=""><span title="class" class="octicon octicon-organization"></span></a></li>
  {{{end}}}
          <li class="right"><a href="/user/setting/profile"><span title="user setting" class="octicon octicon-settings"></span></a></li>
          <li class="right">
            <a href="/{{{.username}}}"><img src="http://gravatar.duoshuo.com/avatar/{{{.email_md5}}}" alt="">{{{.nickname}}}</a>
          </li>
{{{end}}}
      	</ul>
      </div>

      <div class="main">
      -->
      aaaa
