package sitemap

// from www.sitemaps.org/index.html
const testPage = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
    <title>sitemaps.org - Home</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <script type="text/javascript">  var appInsights=window.appInsights||function(config){function i(config){t[config]=function(){var i=arguments;t.queue.push(function(){t[config].apply(t,i)})}}var t={config:config},u=document,e=window,o="script",s="AuthenticatedUserContext",h="start",c="stop",l="Track",a=l+"Event",v=l+"Page",y=u.createElement(o),r,f;y.src=config.url||"https://az416426.vo.msecnd.net/scripts/a/ai.0.js";u.getElementsByTagName(o)[0].parentNode.appendChild(y);try{t.cookie=u.cookie}catch(p){}for(t.queue=[],t.version="1.0",r=["Event","Exception","Metric","PageView","Trace","Dependency"];r.length;)i("track"+r.pop());return i("set"+s),i("clear"+s),i(h+a),i(c+a),i(h+v),i(c+v),i("flush"),config.disableExceptionTracking||(r="onerror",i("_"+r),f=e[r],e[r]=function(config,i,u,e,o){var s=f&&f(config,i,u,e,o);return s!==!0&&t["_"+r](config,i,u,e,o),s}),t}({instrumentationKey:"7e3b1cc3-2b8a-4ed1-a44d-254e44c50719"});window.appInsights=appInsights;appInsights.trackPageView();</script>
    <meta name="description" content="The Sitemaps protocol enables webmasters to information earch engine about pages on their site that are available for crawling." />
        <link rel="stylesheet" type="text/css" href="/sitemaps.css" media="screen, projection">
<script type="text/javascript" src="/lang.js"></script>
    <style type="text/css">
        .style1
        {
            color: #FF0000;
        }
    </style>
</head>
<body>
    <div id="container">
        <div id="intro">
            <div id="pageHeader">
                <h1>
                    sitemaps.org</h1>
            </div>
            <div id="selectionbar">
                <ul>
                    <li><a href="faq.php">FAQ</a></li>
                    <li><a href="protocol.php">Protocol</a></li>
                    <li class="activelink"><a href="#">Home</a></li>
                </ul>
            </div>
            <!-- end selectionbar -->
        </div>
        <!-- end intro -->
        <div style="padding: 14px; float: right;" id="languagebox">
            Language: <span id="langContainer"></span>
        </div>
        <div id="mainContent">
            <h1>
                What are Sitemaps?</h1>
            <p>
                Sitemaps are an easy way for webmasters to inform search engines about pages on
                their sites that are available for crawling. In its simplest form, a Sitemap is
                an XML file that lists URLs for a site along with additional metadata about each
                URL (when it was last updated, how often it usually changes, and how important it
                is, relative to other URLs in the site) so that search engines can more intelligently
                crawl the site.</p>
            <p>
                Web crawlers usually discover pages from links within the site and from other sites.
                Sitemaps supplement this data to allow crawlers that support Sitemaps to pick up
                all URLs in the Sitemap and learn about those URLs using the associated metadata.
                Using the Sitemap <a href="protocol.php">protocol</a> does not guarantee that web
                pages are included in search engines, but provides hints for web crawlers to do
                a better job of crawling your site.</p>
            <p>
                Sitemap 0.90 is offered under the terms of the <a href="http://creativecommons.org/licenses/by-sa/2.5/">
                    Attribution-ShareAlike Creative Commons License</a> and has wide adoption, including
                support from Google, Yahoo!, and Microsoft.</p>
            <p class="date">
                Last Updated: 27 February 2008
            </p>
        </div>
        <!-- end maincontent -->
    </div>
    <!-- closes #container -->
    <div id="footer">
        <p>
            <a href="terms.php">Terms and conditions</a></p>
    </div>
</body>
</html>
`

const testPageWithBaseTag = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
	<base href="http://example.com"> 
</head>
<body>

       <ul>
           <li><a href="http://example.com/testpage.php">test</a></li>
           <li><a href="http://externallink.com/testpage.php">test</a></li>
           <li><a href="faq.php">FAQ</a></li>
           <li><a href="/protocol.php">Protocol</a></li>
           <li><a href="//example.com/doubleslash">doubleslash</a></li>
           <li class="activelink"><a href="#">Home</a></li>
       </ul>
</body>
</html>
`
