[Foundations](https://lanej.io/charmed/foundations.html)Text

# [Text](https://lanej.io/charmed/foundations-text.html#level-foundations-text)

1 component · [All foundations](https://lanej.io/charmed/foundations.html)

- [Permissions](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions)

[Foundations](https://lanej.io/charmed/foundations.html) / Text

## [Permissions](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions)

Colored permission combinations, exact rwx details, unknown data and explicit special bits.

### [Native preview](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions-preview)

Saved terminal output · truecolor reference

Output of the Go usage example below

```
▛ ▌ ▌  rwxr-xr-x
? ? ?  ?????????
```

Demo · truecolor · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
▘  r--  read                                                                                        
▀  rw-  read + write                                                                                
▌  r-x  read + execute                                                                              
▛  rwx  read + write + execute                                                                      
▝  -w-  write                                                                                       
▖  --x  execute                                                                                     
▞  -wx  write + execute                                                                             
∅  ---  none                                                                                        
                                                                                                    
Owner / group / other    exact mode                                                                 
▛ ▌ ▌  docs/      rwxr-xr-x                                                                         
▀ ▘ ▘  app.ts     rw-r--r--                                                                         
▛ ▌ ▌  deploy.sh  rwxr-xr-x                                                                         
▀ ∅ ∅  .env       rw-------                                                                         
▘ ▘ ▘  report.pdf r--r--r--                                                                         
? ? ?  uninspected ?????????                                                                        
                                                                                                    
Special bits / explicit modifiers                                                                   
▛ ▌ ▌ setuid  rwsr-xr-x                                                                             
▀ ▘ ▘ setgid  rw-r-Sr--                                                                             
▛ ▛ ▛ sticky  rwxrwxrwt                                                                             
                                                                                                    
Mode bits are not effective access.
```

Narrow · truecolor · 40 × 30 terminal cells

```
Permission combinations / one group     
Read: top left. Write: top right.       
Execute: bottom left.                   
                                        
▘  r--  read                            
▀  rw-  read + write                    
▌  r-x  read + execute                  
▛  rwx  read + write + execute          
▝  -w-  write                           
▖  --x  execute                         
▞  -wx  write + execute                 
∅  ---  none                            
                                        
Owner / group / other    exact mode     
▛ ▌ ▌  docs/      rwxr-xr-x             
▀ ▘ ▘  app.ts     rw-r--r--             
▛ ▌ ▌  deploy.sh  rwxr-xr-x             
▀ ∅ ∅  .env       rw-------             
▘ ▘ ▘  report.pdf r--r--r--             
? ? ?  uninspected ?????????            
                                        
Special bits / explicit modifiers       
▛ ▌ ▌ setuid  rwsr-xr-x                 
▀ ▘ ▘ setgid  rw-r-Sr--                 
▛ ▛ ▛ sticky  rwxrwxrwt                 
                                        
Mode bits are not effective access.
```

256 colors · ansi256 · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
▘  r--  read                                                                                        
▀  rw-  read + write                                                                                
▌  r-x  read + execute                                                                              
▛  rwx  read + write + execute                                                                      
▝  -w-  write                                                                                       
▖  --x  execute                                                                                     
▞  -wx  write + execute                                                                             
∅  ---  none                                                                                        
                                                                                                    
Owner / group / other    exact mode                                                                 
▛ ▌ ▌  docs/      rwxr-xr-x                                                                         
▀ ▘ ▘  app.ts     rw-r--r--                                                                         
▛ ▌ ▌  deploy.sh  rwxr-xr-x                                                                         
▀ ∅ ∅  .env       rw-------                                                                         
▘ ▘ ▘  report.pdf r--r--r--                                                                         
? ? ?  uninspected ?????????                                                                        
                                                                                                    
Special bits / explicit modifiers                                                                   
▛ ▌ ▌ setuid  rwsr-xr-x                                                                             
▀ ▘ ▘ setgid  rw-r-Sr--                                                                             
▛ ▛ ▛ sticky  rwxrwxrwt                                                                             
                                                                                                    
Mode bits are not effective access.
```

16 colors · ansi16 · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
▘  r--  read                                                                                        
▀  rw-  read + write                                                                                
▌  r-x  read + execute                                                                              
▛  rwx  read + write + execute                                                                      
▝  -w-  write                                                                                       
▖  --x  execute                                                                                     
▞  -wx  write + execute                                                                             
∅  ---  none                                                                                        
                                                                                                    
Owner / group / other    exact mode                                                                 
▛ ▌ ▌  docs/      rwxr-xr-x                                                                         
▀ ▘ ▘  app.ts     rw-r--r--                                                                         
▛ ▌ ▌  deploy.sh  rwxr-xr-x                                                                         
▀ ∅ ∅  .env       rw-------                                                                         
▘ ▘ ▘  report.pdf r--r--r--                                                                         
? ? ?  uninspected ?????????                                                                        
                                                                                                    
Special bits / explicit modifiers                                                                   
▛ ▌ ▌ setuid  rwsr-xr-x                                                                             
▀ ▘ ▘ setgid  rw-r-Sr--                                                                             
▛ ▛ ▛ sticky  rwxrwxrwt                                                                             
                                                                                                    
Mode bits are not effective access.
```

Monochrome · none · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
r--  r--  read                                                                                      
rw-  rw-  read + write                                                                              
r-x  r-x  read + execute                                                                            
rwx  rwx  read + write + execute                                                                    
-w-  -w-  write                                                                                     
--x  --x  execute                                                                                   
-wx  -wx  write + execute                                                                           
---  ---  none                                                                                      
                                                                                                    
Owner / group / other    exact mode                                                                 
rwx r-x r-x  docs/      rwxr-xr-x                                                                   
rw- r-- r--  app.ts     rw-r--r--                                                                   
rwx r-x r-x  deploy.sh  rwxr-xr-x                                                                   
rw- --- ---  .env       rw-------                                                                   
r-- r-- r--  report.pdf r--r--r--                                                                   
??? ??? ???  uninspected ?????????                                                                  
                                                                                                    
Special bits / explicit modifiers                                                                   
rws r-x r-x setuid  rwsr-xr-x                                                                       
rw- r-S r-- setgid  rw-r-Sr--                                                                       
rwx rwx rwt sticky  rwxrwxrwt                                                                       
                                                                                                    
Mode bits are not effective access.
```

light · truecolor · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
▘  r--  read                                                                                        
▀  rw-  read + write                                                                                
▌  r-x  read + execute                                                                              
▛  rwx  read + write + execute                                                                      
▝  -w-  write                                                                                       
▖  --x  execute                                                                                     
▞  -wx  write + execute                                                                             
∅  ---  none                                                                                        
                                                                                                    
Owner / group / other    exact mode                                                                 
▛ ▌ ▌  docs/      rwxr-xr-x                                                                         
▀ ▘ ▘  app.ts     rw-r--r--                                                                         
▛ ▌ ▌  deploy.sh  rwxr-xr-x                                                                         
▀ ∅ ∅  .env       rw-------                                                                         
▘ ▘ ▘  report.pdf r--r--r--                                                                         
? ? ?  uninspected ?????????                                                                        
                                                                                                    
Special bits / explicit modifiers                                                                   
▛ ▌ ▌ setuid  rwsr-xr-x                                                                             
▀ ▘ ▘ setgid  rw-r-Sr--                                                                             
▛ ▛ ▛ sticky  rwxrwxrwt                                                                             
                                                                                                    
Mode bits are not effective access.
```

contrast · truecolor · 100 × 30 terminal cells

```
Permission combinations / one group                                                                 
Read: top left. Write: top right.                                                                   
Execute: bottom left.                                                                               
                                                                                                    
▘  r--  read                                                                                        
▀  rw-  read + write                                                                                
▌  r-x  read + execute                                                                              
▛  rwx  read + write + execute                                                                      
▝  -w-  write                                                                                       
▖  --x  execute                                                                                     
▞  -wx  write + execute                                                                             
∅  ---  none                                                                                        
                                                                                                    
Owner / group / other    exact mode                                                                 
▛ ▌ ▌  docs/      rwxr-xr-x                                                                         
▀ ▘ ▘  app.ts     rw-r--r--                                                                         
▛ ▌ ▌  deploy.sh  rwxr-xr-x                                                                         
▀ ∅ ∅  .env       rw-------                                                                         
▘ ▘ ▘  report.pdf r--r--r--                                                                         
? ? ?  uninspected ?????????                                                                        
                                                                                                    
Special bits / explicit modifiers                                                                   
▛ ▌ ▌ setuid  rwsr-xr-x                                                                             
▀ ▘ ▘ setgid  rw-r-Sr--                                                                             
▛ ▛ ▛ sticky  rwxrwxrwt                                                                             
                                                                                                    
Mode bits are not effective access.
```

### [Go usage](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions-code)

Imports: [`layout`](https://pkg.go.dev/github.com/lanej/charmed/layout "github.com/lanej/charmed/layout") · [`text`](https://pkg.go.dev/github.com/lanej/charmed/text "github.com/lanej/charmed/text") · [`theme`](https://pkg.go.dev/github.com/lanej/charmed/theme "github.com/lanej/charmed/theme")

```go
palette := theme.New(theme.Dark)
mode := text.Permissions{Mode: 0o755, Available: true}
unknown := text.Permissions{}

return layout.Column(
	mode.Compact(palette)+"  "+mode.Label(),
	unknown.Compact(palette)+"  "+unknown.Label(),
)
```

### [Updates, focus & sizing](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions-integration)

Render from application-owned values and pass one shared theme to related elements. These display helpers have no event loop or keyboard focus; the parent supplies updated values and recomputes the available terminal-cell allocation when the screen changes size.

### [Run this example](https://lanej.io/charmed/foundations-text.html#foundations%2fpermissions-run)

```sh
go run ./examples/catalog foundations/permissions
```

[Example source](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/docs/snippets/foundations_permissions.go) · [Implementation](https://github.com/lanej/charmed/blob/b12c045a6fa7af4e77a5defd71387d493b0381c3/story/permissions.go) · [Set up a checkout](https://lanej.io/charmed/getting-started.html)

Dependency map: not audited.
