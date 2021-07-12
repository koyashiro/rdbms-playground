FROM node:latest as development
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
CMD ["yarn", "dev", "--host", "0.0.0.0", "--port", "80"]

FROM development as build
RUN yarn build

FROM gcr.io/distroless/nodejs:latest as production
WORKDIR /root
COPY --from=build /app/node_modules /app/node_modules
COPY --from=build /app/dist /app/dist
EXPOSE 80
CMD ["/app/node_modules/.bin/vite", "preview", "--host", "0.0.0.0","--port", "80"]
