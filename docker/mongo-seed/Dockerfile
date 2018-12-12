FROM mongo

RUN mkdir -p /app \
 && echo "{}" > /app/data.json

COPY entrypoint.sh /app

WORKDIR /app

CMD /app/entrypoint.sh
