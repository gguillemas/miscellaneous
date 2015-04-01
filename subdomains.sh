#!/bin/bash

ITERATIONS=10

function search() {
	curl -s -L "http://www.google.com/search?q=$1&num=100" -A 'Mozilla/4.0' |\
	grep -a -oP '(?<=<font color="green">)(.*?)(?=</font>)' |\
	sed 's/\(https\?:\/\/\)\?\([^\/]*\).*/\2/g' |\
	sort -u
}

query="site:$1"
for i in $(seq 1 $ITERATIONS); do
	for domain in $(search "$query"); do
		if [[ $query != *":$domain"* ]]; then
			echo $domain
			query="${query}%20-site:$domain"
		fi
	done
done
