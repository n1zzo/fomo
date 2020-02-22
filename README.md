# fomo
## A Modular Event Aggregator

Fomo aims to collect events from a collection of sources and propose them
through different channels.

## Introduction

The event aggregator should be divided into a frontend, used to access the crawled information and a backed, used to retrieve information.

The backends should be modular, some of the possible information sources could be:
- Facebook pages
- Website crawling
- Instagram accounts

The frontend should expose the retrieved information into different formats, either configurable or non-configurable, ideas for the frontends could be:
- Twitter account
- RSS feed
- Instagram page
- Website

Probably the most appropriate language is Python3, the most stable way
to build a parser is to use xpath.

A side goal could be to crawl events from inappropriate places
like Facebook.

## Building

```
go get github.com/n1zzo/fomo
```

## Credits

Work Sans was created by Wei Huang and is licensed under the SIL Open Font
License v1.1 (http://scripts.sil.org/OFL).
