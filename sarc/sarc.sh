#!/usr/bin/env bash

input="${*,,}"
input=$(echo "$input" | xargs -0) #trim with space
character_index=0
should_uppercase="false"

if [ -z "$input" ]; then echo "USAGE: sarc The text you want to be sarcastic"; fi

while [ "$character_index" -lt "${#input}" ]
do
  char="${input:$character_index:1}"

  # only toggle (or not) the character if it's an alphabetical character
  if [[ $char =~ [[:alpha:]] ]]
  then
    if $should_uppercase
    then
      char="$(echo "$char" | tr '[:lower:]' '[:upper:]')"
      should_uppercase="false"
    else
      should_uppercase="true"
    fi
  fi

  output="${output}$char"
  character_index=$(( character_index + 1 ))
done

echo "$output"
