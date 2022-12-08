#!/bin/bash

# shellcheck disable=SC2164
cd /home/ec2-user/movie-server-pipeline
/usr/local/bin/docker-compose -f docker-compose_env.yaml down -d
/usr/local/bin/docker-compose -f docker-compose_env.yaml up -d
