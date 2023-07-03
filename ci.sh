#!/usr/bin/env bash
set -ex

go generate -x ./...

# only commits if there are changes to commit
if [ -z "$(git status --porcelain)" ];
then
  echo "nothing to commit" 
else
  git config --global user.name 'lingonbot'
  git config --global user.email 'lingonbot@users.noreply.github.com'
  git add .
  git commit -m "automated update of providers"
  git push
fi
