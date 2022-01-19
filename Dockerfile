FROM scratch
COPY ficsit-smm-launcher /
EXPOSE 8080
ENTRYPOINT ["/ficsit-smm-launcher"]