FROM denoland/deno:latest as development
WORKDIR /app
RUN deno install --unstable -A -f -n aleph https://deno.land/x/aleph@v0.3.0-alpha.33/cli.ts
RUN deno cache https://deno.land/x/aleph@v0.3.0-alpha.33/cli/dev.ts
RUN deno cache https://deno.land/x/aleph@v0.3.0-alpha.33/cli/build.ts
RUN deno cache https://deno.land/x/aleph@v0.3.0-alpha.33/cli/start.ts
RUN aleph analyze
COPY . .

FROM development as build
RUN aleph build

FROM denoland/deno:distroless as production
WORKDIR /app
COPY --from=build /deno-dir /deno-dir
COPY --from=build /usr/local/bin/aleph /usr/local/bin/aleph
EXPOSE 80
CMD ["run", "--allow-read", "--allow-write", "--allow-net","--allow-env", "--allow-run", "--allow-plugin", "--allow-hrtime", "--unstable", "https://deno.land/x/aleph@v0.3.0-alpha.33/cli.ts", "start", "--host", "0.0.0.0", "--port", "80"]
