{{{template "inc/header.tpl" .}}}
{{{asset "sass/index.scss"}}}
        <div class="index-info-box">
          <div class="msg">
            <h3>项目目前处于开发阶段，希望跟进或交流的同学欢迎加入QQ群：235149571</h3>
            <p>Goj - Go Online Judge 是使用Golang开发的Online Judge System，系统分为Ojsite和Judger两部分，Ojsite是Goj系统的前端web实现，Judger是Goj系统后台判题服务器的实现，两者可以分离到不同的主机，通过网络进行通讯。</p>
            <p>Ojsite项目地址 <a target="_blank" href="https://github.com/duguying/ojsite">https://github.com/duguying/ojsite</a><br>Judger项目地址 <a target="_blank" href="https://github.com/duguying/judger">https://github.com/duguying/judger</a></p>
            <p><h3>judger客户端首个版本SDK发布，你可以借助SDK开发自己的Online Judge</h3>
            	<ul class="release">
            		<li><i class="octicon octicon-tag"></i><a target="_blank" href="https://github.com/gojudge/gojapi-java/releases/tag/v0.0.1">java sdk</a></li>
            		<li><i class="octicon octicon-tag"></i><a target="_blank" href="https://github.com/gojudge/gojapi-csharp/releases/tag/v0.0.1">C# sdk</a></li>
                <li><i class="octicon octicon-tag"></i><a target="_blank" href="https://github.com/gojudge/gojapi-php/releases">php sdk</a></li>
            	</ul>
            </p>
          </div>

          <div class="tile">
            <div class="sps">
              <h4><i class="octicon octicon-octoface"></i>Ojsite</h4>
              <p>
                <iframe src="http://ghbtns.com/github-btn.html?user=duguying&amp;repo=ojsite&amp;type=watch&amp;count=true" allowtransparency="true" frameborder="0" scrolling="0" width="90" height="20"></iframe>
                <iframe src="http://ghbtns.com/github-btn.html?user=duguying&amp;repo=ojsite&amp;type=fork&amp;count=true" allowtransparency="true" frameborder="0" scrolling="0" width="90" height="20"></iframe><br>
                Goj系统的前端web实现，go语言基于beego框架开发。web前端的开发和部署基于fis解决方案。界面按照简约无图的宗旨进行设计。
              </p>
            </div>

            <div class="sps">
              <h4><i class="octicon octicon-octoface"></i>Judger</h4>
              <p>
                <iframe src="http://ghbtns.com/github-btn.html?user=duguying&amp;repo=judger&amp;type=watch&amp;count=true" allowtransparency="true" frameborder="0" scrolling="0" width="90" height="20"></iframe>
                <iframe src="http://ghbtns.com/github-btn.html?user=duguying&amp;repo=judger&amp;type=fork&amp;count=true" allowtransparency="true" frameborder="0" scrolling="0" width="90" height="20"></iframe><br>
                Goj系统的后端判题服务器，服务基于go语言开发，C/C++沙箱/执行器基于C语言开发。Linux平台C/C++源码支持沙箱运行；Windows平台C/C++源码由执行器执行、不支持沙箱。
              </p>
            </div>

          </div>
        </div>
{{{template "inc/footer.tpl" .}}}
