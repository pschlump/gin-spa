#!/bin/bash

mkdir -p x
cd x

wget -o t1.err -O t1.out 'http://127.0.0.1:9094/index.html'
wget -o t2.err -O t2.out 'http://127.0.0.1:9094/login'
wget -o t3.err -O t3.out 'http://127.0.0.1:9094/logout'

if diff t1.out t2.out >/dev/null ; then
	:
else
	echo "FAILED"
fi

if diff t3.out t1.out >/dev/null ; then
	:
else
	echo "FAILED"
fi
