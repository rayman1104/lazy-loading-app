FROM node:10

WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 4200
CMD npx ng build --prod && npx ng serve --prod --host 0.0.0.0
