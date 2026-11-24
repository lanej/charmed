[Atoms](https://lanej.io/charmed/atoms.html)HTTP

# [HTTP](https://lanej.io/charmed/atoms-http.html#level-atoms-http)

3 components · [All atoms](https://lanej.io/charmed/atoms.html)

- [URL](https://lanej.io/charmed/atoms-http.html#text%2furl)
- [Method](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method)
- [Status code](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status)

[Atoms](https://lanej.io/charmed/atoms.html) / HTTP

## [URL](https://lanej.io/charmed/atoms-http.html#text%2furl)

Path, path with query, host with path, or exact captured URL display; no network requests or mutation.

### [Native preview](https://lanej.io/charmed/atoms-http.html#text%2furl-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
/v1/labels                                                 
/v1/labels?expand=rate&locale=en_US                        
https://api.example.test/v1/labels?expand=rate&locale=en_US
```

Demo · truecolor · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

Narrow · truecolor · 40 × 16 terminal cells

```
URL display / captured source unchanged 
                                        
path                                    
/v1/labels/p_42                         
                                        
path-query                              
/v1/labels/p_42?limit=25&tag=a%2Fb      
                                        
host-path                               
api.example.test:8443/v1/labels/p_42?lim
it=25&tag=a%2Fb                         
                                        
full                                    
https://api.example.test:8443/v1/labels/
p_42?limit=25&tag=a%2Fb#detail
```

256 colors · ansi256 · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

16 colors · ansi16 · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

Monochrome · none · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

light · truecolor · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

contrast · truecolor · 100 × 30 terminal cells

```
URL display / captured source unchanged                                                             
                                                                                                    
path                                                                                                
/v1/labels/p_42                                                                                     
                                                                                                    
path-query                                                                                          
/v1/labels/p_42?limit=25&tag=a%2Fb                                                                  
                                                                                                    
host-path                                                                                           
api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb                                             
                                                                                                    
full                                                                                                
https://api.example.test:8443/v1/labels/p_42?limit=25&tag=a%2Fb#detail
```

### [Go usage](https://lanej.io/charmed/atoms-http.html#text%2furl-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text")

```go
url := "https://api.example.test/v1/labels?expand=rate&locale=en_US"

return layout.Column(
	text.FormatURL(url, text.PathOnly),
	text.FormatURL(url, text.PathAndQuery),
	text.FormatURL(url, text.FullURL),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-http.html#text%2furl-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-http.html#text%2furl-run)

```sh
go run ./examples/catalog text/url
```

[Example source](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/docs/snippets/text_url.go) · [Implementation](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/text/url.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / HTTP

## [Method](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method)

Categorical HTTP verbs with explicit extension-method fallback.

### [Native preview](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 GET    
 POST   
 PATCH  
 DELETE
```

Demo · truecolor · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

Narrow · truecolor · 40 × 16 terminal cells

```
HTTP methods / categorical colors       
                                        
 GET                                    
 HEAD                                   
 POST                                   
 PUT                                    
 PATCH                                  
 DELETE                                 
 OPTIONS                                
 CONNECT                                
 TRACE                                  
 PROPFIND                               
 QUERY                                  
 ?
```

256 colors · ansi256 · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

16 colors · ansi16 · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

Monochrome · none · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

light · truecolor · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

contrast · truecolor · 100 × 30 terminal cells

```
HTTP methods / categorical colors                                                                   
                                                                                                    
 GET                                                                                                
 HEAD                                                                                               
 POST                                                                                               
 PUT                                                                                                
 PATCH                                                                                              
 DELETE                                                                                             
 OPTIONS                                                                                            
 CONNECT                                                                                            
 TRACE                                                                                              
 PROPFIND                                                                                           
 QUERY                                                                                              
 ?
```

### [Go usage](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.HTTPMethod("GET", palette), text.HTTPMethod("POST", palette),
	text.HTTPMethod("PATCH", palette), text.HTTPMethod("DELETE", palette),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-http.html#text%2fhttp-method-run)

```sh
go run ./examples/catalog text/http-method
```

[Example source](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/docs/snippets/text_http_method.go) · [Implementation](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/text/http.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.

[Atoms](https://lanej.io/charmed/atoms.html) / HTTP

## [Status code](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status)

Response classes retain exact codes, including unknown and invalid observations.

### [Native preview](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
 200  OK         
 404  Not Found  
 502  Bad Gateway
```

Demo · truecolor · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

Narrow · truecolor · 40 × 16 terminal cells

```
HTTP response codes / exact observations
                                        
 100  Continue                          
 200  OK                                
 201  Created                           
 204  No Content                        
 301  Moved Permanently                 
 304  Not Modified                      
 400  Bad Request                       
 401  Unauthorized                      
 404  Not Found                         
 429  Too Many Requests                 
 500  Internal Server Error             
 502  Bad Gateway                       
 503  Service Unavailable               
 ?
```

256 colors · ansi256 · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

16 colors · ansi16 · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

Monochrome · none · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

light · truecolor · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

contrast · truecolor · 100 × 30 terminal cells

```
HTTP response codes / exact observations                                                            
                                                                                                    
 100  Continue                                                                                      
 200  OK                                                                                            
 201  Created                                                                                       
 204  No Content                                                                                    
 301  Moved Permanently                                                                             
 304  Not Modified                                                                                  
 400  Bad Request                                                                                   
 401  Unauthorized                                                                                  
 404  Not Found                                                                                     
 429  Too Many Requests                                                                             
 500  Internal Server Error                                                                         
 502  Bad Gateway                                                                                   
 503  Service Unavailable                                                                           
 ?                                                                                                  
 799?
```

### [Go usage](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)

return layout.Column(
	text.HTTPStatusCode(200, palette)+" OK",
	text.HTTPStatusCode(404, palette)+" Not Found",
	text.HTTPStatusCode(502, palette)+" Bad Gateway",
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/atoms-http.html#text%2fhttp-status-run)

```sh
go run ./examples/catalog text/http-status
```

[Example source](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/docs/snippets/text_http_status.go) · [Implementation](https://github.com/lanej/charmed/blob/27e05a1997640606888402c2223835e8d5e8e451/text/http.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
