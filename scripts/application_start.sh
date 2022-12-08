#!/bin/bash

# shellcheck disable=SC2164
cd /home/ec2-user/movie-server-pipeline
docker-compose -f docker-compose_env.yaml down
docker-compose -f docker-compose_env.yaml up
