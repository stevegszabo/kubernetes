### STAGE 1: Build ###

# We label our stage as ‘builder’
FROM node:12-alpine as builder

COPY package.json package-lock.json ./

## Storing node modules on a separate layer will prevent unnecessary npm installs at each build
RUN npm set strict-ssl false && npm i && mkdir /ng-app && cp -R ./node_modules ./ng-app

WORKDIR /ng-app

COPY . .

## Build the angular app in production mode and store the artifacts in dist folder
RUN $(npm bin)/ng build --prod


### STAGE 2: Setup ###

FROM nginx:1.14-alpine

ENV APPLICATION_BACK=application:5000

## Copy our default nginx config
COPY nginx/default.conf /etc/nginx/conf.d/
COPY entrypoint.sh /

## Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

## From ‘builder’ stage copy over the artifacts in dist folder to default nginx public folder
COPY --from=builder /ng-app/dist/hashi-demo-app /usr/share/nginx/html

CMD ["/entrypoint.sh"]
