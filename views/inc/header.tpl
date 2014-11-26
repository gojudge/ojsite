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
  </head>
  <body>
    <div class="wrapper">

      <noscript>Please enable JavaScript in your browser!</noscript>

      <div class="header">
      	<ul>
      		<li><span class="logo octicon octicon-mortar-board" title="Goj"></span></li>
      		<li><a href="">题库</a></li>
          <li><a href="">教学</a></li>
          <li><a href="">讨论</a></li>
          <li>
            <form action="" method="get">
              <input type="text" name="keyword" id="" placeholder="Search">
            </form>
          </li>

{{if eq .userIs "guest"}}
          <li class="right"><a href=""><span title="sign out" class="octicon octicon-sign-in"></span> 登录</a></li>
{{else}}
          <li class="right"><a href=""><span title="sign out" class="octicon octicon-sign-out"></span></a></li>
  {{if eq .userIs "admin"}}
          <li class="right"><a href=""><span title="administration" class="octicon octicon-circuit-board"></span></a></li>
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