FROM node:latest as development
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
CMD ["yarn", "dev", "--hostname", "0.0.0.0", "--port", "3000"]

FROM node:latest as build
WORKDIR /work
COPY . .
RUN yarn install --frozen-lockfile
RUN yarn build

FROM node:latest as production
WORKDIR /app
COPY --from=build /work/.next /app/.next
COPY --from=build /work/package.json /app/package.json
COPY --from=build /work/yarn.lock /app/yarn.lock
COPY --from=build /work/next.config.js /app/next.config.js
RUN yarn install --frozen-lockfile

CMD ["yarn", "start", "--hostname", "0.0.0.0", "--port", "3000"]
