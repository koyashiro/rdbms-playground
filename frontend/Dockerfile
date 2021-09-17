FROM node:latest as development
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
CMD ["yarn", "dev", "--hostname", "0.0.0.0", "--port", "3000"]
