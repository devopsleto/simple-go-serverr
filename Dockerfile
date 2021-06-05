FROM alpine

RUN mkdir /app 

COPY simple-go-serverr /app/simple-go-serverr

CMD cd /app && ./simple-go-serverr