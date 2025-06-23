FROM node:20.10.0-bookworm

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci --force
RUN npm install -g @angular/cli@^14.0.0 http-server
EXPOSE 4200

ENTRYPOINT ["npm", "run"]
CMD ["start:dev"]
