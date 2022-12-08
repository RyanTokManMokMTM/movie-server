#!/bin/bash

echo 'run after_install.sh' >> /home/ec2-user/movie-server/logs/deploy.log
echo 'cd to movie-server direction' >> /home/ec2-user/movie-server/logs/deploy.log

cd /home/ec2-user/movie-server >> /home/ec2-user/movie-server/logs/deploy.log
echo 'go mod tidy' >> /home/ec2-user/movie-server/logs/deploy.log
go mod tidy