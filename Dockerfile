FROM debian:stable

RUN apt update
RUN apt install -y ca-certificates

ADD notely /usr/bin/notely

CMD ["notely"]
