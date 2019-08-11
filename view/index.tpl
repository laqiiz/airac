<!DOCTYPE html>
<title>example-golang-oauth</title>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
      integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<link rel="stylesheet" href="http://netdna.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

<div class="container">
    {{.Message}}
    <div class="well">OAuth Sample Top</div>
    <br/>
    <div class="panel panel-default">
        <div class="panel-heading">
            <i class="fa fa-google" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;Googleログイン
        </div>
        <div class="panel-body">
            <a href="/google/oauth2">OAuth2認証</a>
        </div>
    </div>
    <br/>
    <div class="panel panel-default">
        <div class="panel-heading">
            <i class="fa fa-twitter" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;Twitterログイン
        </div>
        <div class="panel-body">
            <a href="/twitter/oauth">OAuth認証</a>
        </div>
    </div>
    <br/>
    <div class="panel panel-default">
        <div class="panel-heading">
            <i class="fa fa-facebook" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;Facebookログイン
        </div>
        <div class="panel-body">
            <a href="/facebook/oauth2">OAuth2認証</a>
        </div>
    </div>
    <br/>
    <div class="panel panel-default">
        <div class="panel-heading">
            <i class="fa fa-github" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;GitHubログイン
        </div>
        <div class="panel-body">
            <a href="/github/oauth2">OAuth2認証</a>
        </div>
    </div>
    <br/>
</div>
