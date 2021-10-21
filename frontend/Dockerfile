FROM node:17.0.1 as development
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
CMD ["yarn", "dev", "--hostname", "0.0.0.0", "--port", "3000"]

FROM development as build
RUN yarn build
RUN yarn install --production --frozen-lockfile

FROM node:17.0.1-slim as production
WORKDIR /app
COPY --from=build /app/package.json /app/package.json
COPY --from=build /app/yarn.lock /app/yarn.lock
COPY --from=build /app/node_modules /app/node_modules
COPY --from=build /app/.next /app/.next
COPY --from=build /app/next.config.js /app/next.config.js

CMD ["yarn", "start", "--hostname", "0.0.0.0", "--port", "3000"]
