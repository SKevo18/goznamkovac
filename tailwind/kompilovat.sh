#!/bin/bash
cd "$(dirname "$0")"

npx tailwindcss -i tailwind.css -o ../internal/sablonovac/sablony/staticke/tailwind.css
npx cleancss ../internal/sablonovac/sablony/staticke/tailwind.css -o ../internal/sablonovac/sablony/staticke/tailwind.min.css
rm ../internal/sablonovac/sablony/staticke/tailwind.css
