FROM debian:buster-slim
# Official Debian and Ubuntu images automatically run apt-get clean, so explicit invocation is not required.
RUN set -xe && apt-get update && apt-get install -y curl
RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime
RUN echo "Europe/Moscow" > /etc/timezone
COPY /deployments/wait-for-elastic.sh /wait-for-elastic.sh
RUN chmod +x /wait-for-elastic.sh
COPY build/app_microservice /usr/local/bin/app_microservice
RUN chmod +x /usr/local/bin/app_microservice
CMD ["/usr/local/bin/app_microservice"]
EXPOSE 8080