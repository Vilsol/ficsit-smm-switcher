FROM scratch
COPY ficsit-smm-switcher /
EXPOSE 8080
ENTRYPOINT ["/ficsit-smm-switcher"]