FROM debian:stretch

RUN apt-get update && apt-get -y install curl fortune-mod fortunes-bofh-excuses fortunes-off  && apt-get clean

RUN curl -sL -o /tmp/fortunes-spam.deb http://mirrors.cat.pdx.edu/ubuntu/pool/universe/f/fortunes-spam/fortunes-spam_1.8-0ubuntu1_all.deb && dpkg -i /tmp/fortunes-spam.deb

COPY hello-http /

ENTRYPOINT "/hello-http"
