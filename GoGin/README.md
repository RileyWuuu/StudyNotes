# GoGin study notes

## Why use gin?
* Has wrappers around route registration (meaning less boilerplate)
  * net/http   
  ```
  E.g.
  func main(){
    mux := http.NewServeMux()
	  mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
      switch request.Method {
        case "GET":
          writer.WriteHeader(http.StatusOK)
          fmt.Fprint(writer, "pong")
        default:
          http.NotFound(writer, request)
      }
	  })
	  s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
    }
    log.Fatal(s.ListenAndServe())
  }
  ```
  * gin
  ```
  E.g.
  func main(){
    server := gin.New()
    s.Get("ping",func(c *gin.Context){
      c.String(http.StatusOK,"pong")
    })
    log.Fatal(http.ListenAndServe(":8080",s))
  }
  ```
* Offers a few more convenience wrappers around deserialisation (meaning no worries about json.Unmarshal function)
* uses httprouter = faster

## Recap
### Websocket flow chart
![0](https://user-images.githubusercontent.com/71340325/182541888-9fb9ef0b-a66c-4ba3-ae87-7fcaf29ddb8f.jpg)
