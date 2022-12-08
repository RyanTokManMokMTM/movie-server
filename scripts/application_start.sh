#!/bin/bash
echo 'run application_start.sh' >> /home/ec2-user/movie-server/logs/deploy.log
fuser -k 8080/tcp >> /home/ec2-user/movie-server/logs/deploy.log
echo 'nohup go run movieservice.go' >> /home/ec2-user/movie-server/logs/deploy.log
nohup go run movieservice.go