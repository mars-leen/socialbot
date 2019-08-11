#!/bin/bash
a='media'
b='front'
c='admin'
if [[ $1 == ${a} ]]
then
    echo "upload bin program"
    if [[ ! -d /www/new.century.wedding ]]; then
      mkdir -p /www/new.century.wedding
    fi
    if [[ -e www/new.century.wedding/webwedding.bk ]]; then
       rm -rf /www/new.century.wedding/webwedding.bk
    fi
    if [[ -e /www/new.century.wedding/webwedding ]]; then
        mv /www/new.century.wedding/webwedding /www/new.century.wedding/webwedding.bk
    fi
    rz -E
    chmod 777 web
    mv web /www/new.century.wedding/webwedding
    pkill webwedding
    nohup /www/new.century.wedding/webwedding >> m.log 2>&1 &
fi

if [[ $1 == ${b} ]]
then
   echo "upload front view"
   if [[ ! -d /www/new.century.wedding/views/front ]]; then
      mkdir -p /www/new.century.wedding/views/front
    fi
    rz -E
    tar -xvf dist.tar -C /www/new.century.wedding/views/front/
    mv /www/new.century.wedding/views/front/dist /www/new.century.wedding/views/front/socialbot-face
    rm -rf dist.tar
fi

if [[ $1 == ${c} ]]
then
    echo "upload front view"
    if [[ ! -d /www/new.century.wedding/views/admin ]]; then
      mkdir -p /www/new.century.wedding/views/admin
    fi
    rz -E
    tar -xvf dist.tar -C /www/new.century.wedding/views/admin/
    mv /www/new.century.wedding/views/admin/dist /www/new.century.wedding/views/admin/socialbot-brain
    rm -rf dist.tar
fi
