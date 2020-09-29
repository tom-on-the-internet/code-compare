#! /usr/bin/env python

import sys

text: str = " ".join([arg.strip() for arg in sys.argv[1:]])

if not text:
    print("USAGE: sarc The text you want to be sarcastic")
    sys.exit(1)

should_upper_case = True
output = ""

for character in text:
    if character.isalpha():
        should_upper_case = not should_upper_case

    output = (
        output + character.upper() if should_upper_case else output + character.lower()
    )

print(output)
