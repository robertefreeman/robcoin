FROM alpine:3.7

RUN apk --no-cache add  --update py-pip libpq python-dev curl

RUN rm -f /usr/lib/libdb-5.3.so

RUN pip install flask==0.10.1 && pip install -U flask-cors

ADD / /app

WORKDIR /app

CMD python app.py
