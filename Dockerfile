FROM debian:10

COPY ./dumbstored /bin/dumbstored

VOLUME /var/dumbstored

CMD ["/bin/dumbstored"]
