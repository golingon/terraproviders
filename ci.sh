#!/bin/usr/env bash

go generate .

# only commits if there are changes to commit
[ -z "$(git status --porcelain)" ]  && echo "nothing to commit" || {
  set -ex
  git config --global user.name 'lingonbot'
  git config --global user.email 'lingonbot@users.noreply.github.com'
  git commit -am "automated update of providers"
  git push
}
