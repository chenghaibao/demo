FROM nginx

ENV VERSION long

ARG TESTARG=111

LABEL author=haibao

WORKDIR /usr/src/test

COPY . .

VOLUME ["/data1","/data2"]

EXPOSE 89

USER root

HEALTHCHECK --interval=3s --timeout=2s  CMD curl -f http://localhost/ || exit 1

RUN echo "这是$TESTARG一个$VERSION本地构建的nginx镜像" > /usr/share/nginx/html/index.html \ 
 && mkdir index \ 
 && touch aa.txt 

