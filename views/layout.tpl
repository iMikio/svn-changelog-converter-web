<!doctype html>
<html lang="ja" class="h-100">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="/static/lib/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    {{ block "HtmlHead" . }}{{ end }}
  </head>
<body class="d-flex flex-column h-100">
  <nav class="navbar sticky-top navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Sticky top</a>
    </div>
  </nav>
  <main class="flex-shrink-0">
    <div class="container">
    {{ block "LayoutContent" . }}{{ end }}
    </div>
  </main>

  <footer class="footer mt-auto py-3 bg-light">
    <div class="container">
      <span class="text-muted">
        <div class="author">
          Official website:
          <a href="http://{{.Website}}">{{.Website}}</a> /
          Contact me:
          <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
        </div>
      </span>
    </div>
  </footer>
  <!-- Optional JavaScript; choose one of the two! -->
  <!-- Option 1: Bootstrap Bundle with Popper -->
  <script src="/static/lib/bootstrap/js/bootstrap.bundle.min.js"></script>

  <!-- Option 2: Separate Popper and Bootstrap JS -->
  <!--
  <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js" integrity="sha384-IQsoLXl5PILFhosVNubq5LC7Qb9DXgDA9i+tQ8Zj3iwWAwPtgFTxbJ8NT4GN1R8p" crossorigin="anonymous"></script>
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.min.js" integrity="sha384-cVKIPhGWiC2Al4u+LWgxfKTRIcfu0JTxR+EQDz/bgldoEyl4H0zUF0QKbrJ0EcQF" crossorigin="anonymous"></script>
  -->
  <script src="/static/js/reload.min.js"></script>
  {{ block "Scripts" . }}{{ end }}
</body>
</html>