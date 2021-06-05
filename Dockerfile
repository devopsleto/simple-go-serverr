FROM postgres:latest

RUN mkdir /bd
WORKDIR /bd

CMD postgres -D /bd