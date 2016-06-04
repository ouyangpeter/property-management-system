FROM mysql:5.6

COPY cnf/my.cnf /etc/mysql/my.cnf
RUN /etc/init.d/mysql restart &&\
    sleep 10s &&\
    echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION" | mysql
#RUN chown -R mysql:mysql /var/lib/mysql
