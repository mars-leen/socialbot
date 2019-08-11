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
   if [[ ! -d /www/new.century.wedding/views/front/socialbot-face ]]; then
      mkdir -p /www/new.century.wedding/views/front/socialbot-face
    fi
    rz -E
    tar -xvf dist.tar -C /www/new.century.wedding/views/front/socialbot-face
    rm -rf dist.tar
fi

if [[ $1 == ${c} ]]
then
    echo "upload front view"
    if [[ ! -d /www/new.century.wedding/views/admin/socialbot-brain ]]; then
      mkdir -p /www/new.century.wedding/views/admin/socialbot-brain
    fi
    rz -E
    tar -xvf dist.tar -C /www/new.century.wedding/views/admin/socialbot-brain
    rm -rf dist.tar
fi
