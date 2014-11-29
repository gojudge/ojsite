<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="renderer" content="webkit|ie-stand|ie-comp" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <meta name="description" content="">
    <meta name="author" content="">
    <title>{{.title}}</title>
    {{asset "sass/style.scss"}}
    {{asset "octicons/octicons.scss"}}
    <script src="//cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
  </head>
  <body>
    <div class="wrapper">

      <noscript>Please enable JavaScript in your browser!</noscript>

      <div class="header">
      	<ul>
      		<li><span class="logo octicon octicon-mortar-board" title="Goj"></span></li>
      		<li><a href="/problems">{{i18n "problems"}}</a></li>
          <li><a href="/student/verify">{{i18n "teach"}}</a></li>
          <li><a href="">{{i18n "discuss"}}</a></li>
          <li>
            <form action="" method="get">
              <input type="text" name="keyword" id="" placeholder="Search">
            </form>
          </li>

{{if eq .userIs "guest"}}
          <li class="right"><a href="/register">{{i18n "register"}}</a></li>
          <li class="right"><a href="/login"><span title="sign in" class="octicon octicon-sign-in"></span> {{i18n "login"}}</a></li>
{{else}}
          <li class="right"><a href=""><span title="sign out" class="octicon octicon-sign-out"></span></a></li>
  {{if eq .userIs "admin"}}
          <li class="right"><a href=""><span title="administration" class="octicon octicon-circuit-board"></span></a></li>
  {{end}}
  {{if eq .userIs "teacher"}}
          <li class="right"><span title="teacher" class="octicon octicon-person"></span></li>
  {{end}}
  {{if eq .userIs "student"}}
          <li class="right"><a href=""><span title="class" class="octicon octicon-organization"></span></a></li>
  {{end}}
          <li class="right">
            <a href=""><img src="http://gravatar.duoshuo.com/avatar/5fedd018b3227bc4043934309da8c290" alt="">独孤影</a>
          </li>
{{end}}
      	</ul>
      </div>

      <div class="main">