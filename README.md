# Backend with Go

This template wants to solve the problems that a dev might encounter when trying to test a backend API in a server.

- ### CORS protocol  
It's a mechanism that uses HTTP headers for allowing a browser to access resources in a server, without configure this the browser will not allow any connection to the API we made.  

[github.com/gorilla/mux](https://github.com/gorilla/mux)  
[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
  
- ### HTTPS connection  
Browsers don't allow responses from servers that don't have a SSL certificate. One way to solve this problem is to automate SSL certficates using a reverse proxy with Traefik and "let's encrypt".  

[Let's encrypt](https://letsencrypt.org/getting-started/)  
[Reverse proxy](https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/)  
[Traefik](https://doc.traefik.io/traefik/)  

- ### Deploy with Docker-compose  
Docker offers multiple advantages that simplify the proccess of deploy a service, this implementation is very basic and includes the reverse proxy and the backend API.  

[Docker](https://docs.docker.com/get-started/)
[Docker compose](https://docs.docker.com/compose/)
