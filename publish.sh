#!/bin/sh

echo "Generating site"
hugo

echo "Updating gh-pages branch..."
cd public && git add --all && git commit -m "Publishing to gh-pages"

echo "Publishing to site"
git push -f origin gh-pages
cd ..
