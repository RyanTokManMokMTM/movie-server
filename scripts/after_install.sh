#!/bin/bash
# shellcheck disable=SC2164
#cd /home/ec2-user/movie-server-pipeline
#docker-compose -f docker-compose_env.yaml down

cd /home/ec2-user/movie-server-pipeline
docker build . -t movie-server