# TinyGo Web Site

This is the web site for TinyGo.

Built using Hugo:

http://gohugo.io/

and the Hugo "Learn" theme:

https://github.com/matcornic/hugo-theme-learn

## Installation

Install Hugo command line tool for your operating system.

Clone this repo using:

    git clone git@github.com:tinygo-org/tinygo-site.git

Change directories into the tinygo-site directory:

    cd tinygo-site

Install the "Learn" theme:

    git submodule update --init

Now you should be able to run the site locally:

    hugo serve

Once the site code is running locally, you can navigate to it by going to http://localhost:1313

## Deploy to GH Pages

Perform the one time setup of getting the `gh-pages` branch into your local tree:

    git checkout gh-pages
    git checkout master

Now, assuming you have the correct permissions, you can deploy the site code.

Commit all changes to master.

Run `publish.sh` script.

That's it.

## TODO:

- Generate custom godocs for all tinygo build tags, add to "Documentation" section, and point Godocs links there.
- Activate search.
- Custom domain for all packages
- ?
