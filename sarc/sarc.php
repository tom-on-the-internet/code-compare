<?php

array_shift($argv);

$text = implode(" ", array_map(fn ($arg) => trim($arg), $argv));

if (!$text) {
  print("USAGE: sarc The text you want to be sarcastic" . PHP_EOL);
  exit(1);
}

$shouldUpperCase = true;
$output = "";

for ($i = 0; $i < strlen($text); $i++) {
  if (ctype_alpha($text[$i])) {
    $shouldUpperCase = !$shouldUpperCase;
  }

  $output .= $shouldUpperCase ? strtoupper($text[$i]) : strtolower($text[$i]);
}

print($output . PHP_EOL);
