# URL shortener
A url shortener service written in golang. This service uses base62 encoding to generate short strings associated with a URL.

## Web API
1. POST `/api/url/short` - reads url from payload and returns short url
2. GET `/{path}` - this API is invoked when a short url is entered in a browser
