### run rabbitmq docker container
sudo docker run -it --rm --name rabbitmq_service -p 5672:5672 -p 15672:15672 rabbitmq:3-management 

- http://0.0.0.0:15672
- run sender.go
- run receiver.go