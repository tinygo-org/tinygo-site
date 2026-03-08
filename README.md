# TinyGo Web Site

[![Netlify Status](https://api.netlify.com/api/v1/badges/83fc0c21-220b-4d35-ad59-3d48f31bb4b6/deploy-status)](https://app.netlify.com/sites/tinygo/deploys)

This is the web site for TinyGo.

Built using Hugo:

http://gohugo.io/

and the Hugo "Docsy" theme:

https://github.com/google/docsy

## Prerequisites

Install [Hugo](https://gohugo.io/installation/) command line tool for your
operating system with the "extended". Note that the most recent compatible
version of Hugo is `v0.139.5`.

    CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@v0.139.5

Hugo will automatically download the [docsy](https://www.docsy.dev/docs/get-started/docsy-as-module/)
NodeJS modules needed for processing the site CSS.

## Installation

Clone this repo using:

    git clone git@github.com:tinygo-org/tinygo-site.git

Change directories into the tinygo-site directory:

    cd tinygo-site

You are now able to run the site locally, "Docsy" theme will be automatically downloaded as module:

    hugo serve

Once the site code is running locally, you can navigate to it by going to http://localhost:1313

## Deploy to Netlify

Pushing to the `release` branch automatically deploys the latest site to Netlify.

That's it.

## TODO:

- ?
