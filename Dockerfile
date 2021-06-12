FROM postgres:latest
RUN mkdir /bd
VOLUME /bd
RUN cp -r /var/lib/postgresql/data /bd
USER postgres
CMD postgres -D /bd