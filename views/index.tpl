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
      <div>
        <label for="fileExtensionButton" class="form-label">Choose file extension.</label>
      </div>
      <div class="btn-group" role="group" id="fileExtensionButton" aria-label="file extension type button group">
        <input type="radio" class="btn-check" name="type" id="type1" value="csv" autocomplete="off" checked>
        <label class="btn btn-outline-primary" for="type1">CSV</label>
      
        <input type="radio" class="btn-check" name="type" id="type2" value="tsv" autocomplete="off">
        <label class="btn btn-outline-primary" for="type2">TSV</label>
      </div>
    </div>
    <div class="mb-3">
      <div>
        <label for="characterCodeButton" class="form-label">Choose character code.</label>
      </div>
      <div class="btn-group" role="group" id="characterCodeButton" aria-label="character code button group">
        <input type="radio" class="btn-check" name="code" id="code1" value="utf-8" autocomplete="off" checked>
        <label class="btn btn-outline-primary" for="code1">UTF-8</label>
      
        <input type="radio" class="btn-check" name="code" id="code2" value="shift-jis" autocomplete="off">
        <label class="btn btn-outline-primary" for="code2">Shift-JIS</label>
      </div>
    </div>
    <div class="mb-3">
      <label for="changeLogTextarea" class="form-label">Enter changelog.</label>
      <textarea class="form-control" id="changeLogTextarea" name="text" rows="15"></textarea>
    </div>
    <button type="submit" class="btn btn-primary">Convert</button>
  </form>
</header>
{{ end }}

{{ define "Scripts" }}
{{ end }}
</html>
