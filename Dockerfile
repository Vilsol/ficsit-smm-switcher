FROM scratch
COPY ficsit-switcher /
EXPOSE 8080
ENTRYPOINT ["/ficsit-switcher"]