<!DOCTYPE html>
<title>example-golang-oauth</title>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
      integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

<div class="container">
    <br/>
    <div class="alert alert-success" role="alert">Facebook Account Login Success！</div>
    <div class="panel panel-success">
        <div class="panel-heading">情報</div>
        <div class="panel-body">
            <br/>
            <ul class="list-group">
                <li class="list-group-item">ID：{{.ID}}
                <li class="list-group-item">Name：{{.Name}}
                <li class="list-group-item">Email：{{.Email}}
            </ul>
        </div>
    </div>
    <a href="/">戻る</a>
</div>
