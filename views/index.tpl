{{ template "layout.tpl" . }}

{{ define "HtmlHead" }}

<title>Beego</title>
{{ end }}
<!DOCTYPE html>

{{ define "LayoutContent" }}
<header>
  <h1 class="logo">Welcome to Beego</h1>
  <div class="description">
    Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
  </div>
</header>
<footer>
  <div class="author">
    Official website:
    <a href="http://{{.Website}}">{{.Website}}</a> /
    Contact me:
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
  </div>
</footer>
<div class="backdrop"></div>
{{ end }}

{{ define "Scripts" }}
{{ end }}

</html>
