### run rabbitmq docker container
sudo docker run -it --rm --name rabbitmq_service -p 5672:5672 -p 15672:15672 rabbitmq:3-management 

### steps to follow
- Run rabbitmq docker container
- Open http://0.0.0.0:15672 in browser
- Use login creds as username:password, guest:guest
- In terminal excute `go run sender.go`
- In another terminal excute `go run receiver.go`