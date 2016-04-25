#!/bin/bash
docker run -d -e "MYSQL_ROOT_PASSWORD=sdf723usdf7weri" --name "docker_mysql" -p "3706:3306" -v "/Users/sep/opt/var/property-management-system/mysql:/var/lib/mysql" mysql_my:v1.0
