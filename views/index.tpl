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
  <form action="/" method="post">
    <div class="mb-3">
      <label for="changeLogTextarea" class="form-label">Example textarea</label>
      <textarea class="form-control" id="changeLogTextarea" name="text" rows="3"></textarea>
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
  </form>
</header>
{{ end }}

{{ define "Scripts" }}
{{ end }}

</html>
