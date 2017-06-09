# go_glossary_tutorial
A tutorial on how I created Golossary.  Check out the live streams [here](https://www.liveedu.tv/vohumana/lmnXo-golossary/)

## Overview
This tutorial will show you how to create a simple HTTP API in Go.  The API will accept text as part of the request and return JSON of the words and definitions.  Behind the scenes, we will be making a call to a free dictionary API.  I will guide you from "Hello World" all the way to a command line app you can use to generate a glossary.

**What are the requirements?**
- A basic understanding of Go.
- A properly configured install of Go.
- A basic understanding of how the web works.

**What is the target audience?**
- Someone wanting to learn how to build and consume API's in Go

**When are the streaming sessions (streaming schedule)?**
Every Monday at 1900 (7:00PM) PST until the project is complete.

## Session breakdown
**Session 1:** Setting up main and calling the dictionary API
- Setting up the project
- Create Hello World
- Making a simple call to the dictionary API

**Result:** A program that can get JSON from the dictionary API and print it to the screen.

**Session 2:** Parsing the response and abstracting our dictionary
- Use the Go standard library to parse JSON
- Create structs to parse the Pearson API
- Use a tool I made to simplify consuming API's (with explanation how it works)
- Creating an abstraction

**Result:** A program that can parse JSON from the dictionary and print only the definition to the screen.

**Session 3:** Parse a text file, call the API, and clean up code.
- Load a text file and parse the text from it
- Call the API for each unique word we find
- Add command line flags

**Result:** A program that can parse the text from a file and define all of the unique words in that file.

## Resources 
- [Pearson Dictionaries](http://developer.pearson.com/apis/dictionaries#/)
- [Golang](https://golang.org/)
- [Great API for finding documentation](http://devdocs.io/)