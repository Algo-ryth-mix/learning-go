# go url-shortener

simple http url-shortener

- server.go contains the runtime component that starts the server
- shorten.go contains the endpoints for shortening a url and redirecting to the short url
- global.go contains the kv-store for storing the urls 
- (*) show.go contains a small extra function to show all urls that are stored in the kv-store
- (*) backup.go contains extra functionality to keep a persistent image of the stored urls, so that the server can "remember" saved urls
- letterInt.go contains a small snippet to convert numbers to a 5-letter code
- static/index.html contains the home-page of the url-shortener
