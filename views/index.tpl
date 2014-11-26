
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="renderer" content="webkit|ie-stand|ie-comp" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <meta name="description" content="">
    <meta name="author" content="">
    <title>Online Judge System</title>
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

          <li class="right"><a href=""><span title="sign out" class="octicon octicon-sign-out"></span></a></li>
          <li class="right"><a href=""><span title="administration" class="octicon octicon-circuit-board"></span></a></li>
          <li class="right"><a href=""><span title="class" class="octicon octicon-organization"></span></a></li>
          <li class="right">
            <a href=""><img src="http://gravatar.duoshuo.com/avatar/5fedd018b3227bc4043934309da8c290" alt="">独孤影</a>
          </li>
      	</ul>
      </div>

      <div class="main">
        <div class="index-info-box">
          
        </div>
      </div>

    </div>

    <div class="footer">
      <p class="left">©2015 Goj · 版本:{{ver "app"}} · 加载时间</p>

      <div class="right">
        <a href="https://github.com/duguying/ojsite"><span class="octicon octicon-mark-github"></span></a>
        <div>语言选项</div>
        <a href="http://oj.duguying.net">官方网站</a>
        <span>{{ver "golang"}}</span>
      </div>
        
    </div>

  </body>
  <script src="http://cdn.bootcss.com/jquery/1.11.1/jquery.min.js"></script>
</html>
