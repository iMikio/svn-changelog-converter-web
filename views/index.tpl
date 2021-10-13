{{ template "layout.tpl" . }}

{{ define "HtmlHead" }}

<title>SVN ChangeLog Converter</title>
{{ end }}
<!DOCTYPE html>

{{ define "LayoutContent" }}
<header>
  <h1 class="logo">SVN ChangeLog Converter</h1>
  <form action="/" method="post">
    <div class="mb-3">
      <div class="btn-group" role="group" aria-label="Basic radio toggle button group">
        <input type="radio" class="btn-check" name="type" id="type1" value="csv" autocomplete="off" checked>
        <label class="btn btn-outline-primary" for="type1">CSV</label>
      
        <input type="radio" class="btn-check" name="type" id="type2" value="tsv" autocomplete="off">
        <label class="btn btn-outline-primary" for="type2">TSV</label>
      </div>
    </div>
    <div class="mb-3">
      <label for="changeLogTextarea" class="form-label">Please enter changelog.</label>
      <textarea class="form-control" id="changeLogTextarea" name="text" rows="15"></textarea>
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
  </form>
</header>
{{ end }}

{{ define "Scripts" }}
{{ end }}
</html>
