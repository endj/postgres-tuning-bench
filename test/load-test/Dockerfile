FROM python:3.7-slim

COPY loadtest.py .

RUN apt-get update && \
    apt-get install gcc -y && \
    pip install locust && \
    apt-get remove -y gcc && \
    apt-get -y autoremove

CMD [ "locust","--headless","--host","https://gateway","-f", "./loadtest.py" ]
