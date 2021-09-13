FROM node:16

WORKDIR /workdir

COPY yarn.lock tsconfig.json package.json /workdir/

RUN yarn global add serve

RUN yarn install --no-progress

COPY public /workdir/public
COPY src /workdir/src

RUN yarn build

ENTRYPOINT [ "serve", "-s", "build" ]