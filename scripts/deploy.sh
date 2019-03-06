#!/bin/bash

openssl aes-256-cbc -K $encrypted_0c35eebf403c_key -iv $encrypted_0c35eebf403c_iv -in id_ed25519_seiyu5off.enc -out ~/.ssh/id_ed25519_seiyu5off -d
chmod 600 ~/.ssh/id_ed25519_seiyu5off
echo -e "
Host github.com
StrictHostKeyChecking no
IdentityFile ~/.ssh/id_ed25519_seiyu5off
" >> ~/.ssh/config

git config --global user.name "Travis CI"
git config --global user.email "travis.logo.bee.doo19@gmail.com"
git clone git@github.com:shooontan/seiyu5off.git

cd seiyu5off
mkdir -p www/api
cp -ar ../www/api/* ./www/api

git add .
if [[ $(git diff --cached) ]]; then
  git commit -m ":calendar: by Travis CI"
  git push origin master
else
  echo "nothing to commit, working tree clean"
fi
