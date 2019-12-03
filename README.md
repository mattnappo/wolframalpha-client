# The WolframAlpha Desktop Client Backend

[![Go Report Card](https://goreportcard.com/badge/github.com/xoreo/wolframalpha-client)](https://goreportcard.com/report/github.com/xoreo/wolframalpha-client)
[![Build Status](https://travis-ci.com/xoreo/wolframalpha-client.svg?branch=master)](https://travis-ci.com/xoreo/wolframalpha-client)
[![Code Quality](https://api.codacy.com/project/badge/Grade/02a0881e998c4f86bdfe877c6dee220a?isInternal=true)](https://travis-ci.com/xoreo/wolframalpha-client://app.codacy.com/manual/xoreo/wolframalpha-client/dashboard)

The WolframAlpha Desktop Client Backend

## The API
The backend launches a Selenium session (in the `core` package), and then scrapes the Wolfram Alpha DOM (in the `engine`) package to obtain the results from Wolfram Alpha. Then, an API server is initialized.
<br>These are the currently-supported routes in that API:

**search** - makes a search on Wolfram Alpha and returns the response images/LaTeX.

| Attribute | Value                                   |
|-----------|-----------------------------------------|
| Method    | `POST`                                  |
| Endpoint  | `/api/search/`                          |
| Request   | ```{search_text: "the search query"}``` |
