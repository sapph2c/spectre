FROM scratch
COPY ./spectre /
ENTRYPOINT ["/spectre"]
