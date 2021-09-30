FROM ubuntu:latest
MAINTAINER steve.g.szabo

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y curl wget python3 python3-pip gunicorn && \
    apt-get clean && \
    pip3 install flask

ENV WEBAPP_VERSION=1.0.3

COPY root/app /app
RUN chown -R www-data:www-data /app
EXPOSE 8080
USER www-data
ENTRYPOINT ["/app/entrypoint.sh"]
