# API 

 API responsible for processing for working with fake new

### Run 
```go
go build .
go run main.go server 
```

### Docker 
```
docker build -t gcloud/fakefinder-crawler:1.0.0 .
docker run --name  app -d --network fakefinderstack_elk -p 8080:8080 --env PLN_URL={{ENV}} -t gcloud/fakefinder-crawler:1.0.0
docker logs bot --follow

```

### Exemplo
```
request:
curl -H "Content-Type: application/json" -d '{"description":"Anvisa recebe informações oficiais para avaliar retomada de testes da vacina de Oxford para Covid-19"}' http://localhost:8080/search/news

response: 
{"description":"Anvisa recebe informações oficiais para avaliar retomada de testes da vacina de Oxford para Covid-19",
"text":[{"date":"2020-09-12T20:38:51Z",
"title":"Anvisa recebe informações oficiais para avaliar retomada de testes da vacina de Oxford para Covid-19",
"similarity":"0.08620689655172414",
"link":"https://g1.globo.com/economia/noticia/2020/09/12/candidato-de-trump-e-eleito-para-comandar-bid.ghtml"}]}

```

### Export
```
export PLN_URL={{something}}

```

### Template Application 
![Tamplete](https://drstearns.github.io/tutorials/gomiddleware/img/flow.png)

### References
[Middleware](https://drstearns.github.io/tutorials/gomiddleware/) - Middleware Patterns in Go.  
[Route](https://thenewstack.io/make-a-restful-json-api-go/) - Making a RESTful JSON API in Go.  
[serve](https://aaf.engineering/go-web-application-structure-part-2/) - Go Web Application Structure - Part 2 - Routing/Serving.  
[Handler](https://www.alexedwards.net/blog/making-and-using-middleware) - Making and Using HTTP Middleware.  
[HandlerFunc](https://www.alexedwards.net/blog/a-recap-of-request-handling) - A Recap of Request Handling in Go.  
[ChainMiddleware](https://kenyaappexperts.com/blog/using-middleware-in-go/) - Using Middleware In Go The Right Way.  
[ServeMux](https://gist.github.com/delsner/64e79da93a77aa364e79013d3baeaa3e#file-address-go-L10) - Git repository.
[Elastic](//https://olivere.github.io/elastic/)
[Elastic1](https://www.freecodecamp.org/news/go-elasticsearch/)
[Elastic2](https://olivere.github.io/elastic/)