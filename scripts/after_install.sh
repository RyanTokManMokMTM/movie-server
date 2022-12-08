#!/bin/bash

echo "stop and restart docker-compose" >> /home/ec2-user/movie-server/logs/deploy.log
docker-compose -f docker-compose_env.yaml down
