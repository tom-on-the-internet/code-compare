#! /usr/bin/env node

const [, , ...args] = process.argv;
const text = args.map((arg) => arg.trim()).join(" ");

if (!text) {
  console.log("USAGE: sarc The text you want to be sarcastic");
  process.exit(1);
}

let shouldUpperCase = true;
let output = "";

text.split("").forEach((character) => {
  const isAlpha = character.toLowerCase() != character.toUpperCase();
  if (isAlpha) {
    shouldUpperCase = !shouldUpperCase;
  }

  output += shouldUpperCase ? character.toUpperCase() : character.toLowerCase();
});

console.log(output);
