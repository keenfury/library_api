#!/bin/bash

# please run from api/scripts directory

if [ "$#" -ne 2 ]; then
    echo "Need new repo name and package folder";
    exit 1;
fi

repo_name=$1
path=$2

# copy files
mkdir ../../$1
cp -r ../* ../../$1
cp -r ../.vscode ../../$1

# rename cmd/api
mv ../../$1/cmd/api ../../$1/cmd/$1

encoded=$(echo $path | sed 's;/;\\/;g')
encoded=$(echo $encoded | sed 's;\.;\\.;g')
cd ../../$repo_name

# mac
find . -type f -print0 -exec sed -i '' "s/bitbucket\.org\/keenfury\/api/$encoded/g" {} +
# linux or gnu sed
# find . -type f -print0 -exec sed -i "s/bitbucket\.org\/keenfury\/api/$encoded/g" {} +
