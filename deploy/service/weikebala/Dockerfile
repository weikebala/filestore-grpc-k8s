FROM centos:7

ADD bin/weikebala /

RUN chmod 777 /weikebala

ENV PARAMS=""

ENTRYPOINT ["sh","-c","/weikebala $PARAMS"]