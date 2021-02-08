FROM mysql:5.7

ADD ./init /docker-entrypoint-initdb.d
EXPOSE 3306
CMD ["mysqld"]
