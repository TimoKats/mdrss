# 📝MDRSS

MDRSS is a markdown to RSS converter written in GO. With this tool, you can write articles in a local folder and have them automatically formatted to an RSS compliant XML file. As a result, you don't have to write articles on your website first and have them be read by an RSS reader. Moreover, MDRSS automatically takes care of publication dates, categories (next update), formatting, etc.

### Getting started
You can install the binary using `go install github.com/TimoKats/mdrss@latest` (make sure your GOPATH is set correctly!). After this you can add your feeduration in `~/.mdrss/feed.json`. This is a JSON file with a number of settings. You can use this template to setup your feeduration.

```JSON
  {
    "Author": "Timo Kats",
    "Description": "Timo's weblog",
    "InputFolder": "/path/to/articles/", # This will contain your markdown files
    "OutputFile": "/path/to/webserver/index.xml",
    "Link": "index.xml" # can be anything, might be deprecated
  }
```

Note, since the last update (credit to [@varphi-online](https://github.com/varphi-online)) you can also run `mdrss init` to enter an interactive prompt that will fill in this feed for you. Finally, you can also specify your own feed path like so.

```shell
mdrss --feed other/path/to/feed update
```

In your input folder, you can add an article by creating a `.md` file. **Note, if your filename is prefixed with `draft-` it will not be included in the RSS file**. Finally, you can type `mdrss ls` to view the articles ready for publishing and `mdrss update` to create the RSS feed. Note, the title of the RSS articles are based on markdown headers. E.g. format the first line of the markdown articles as so: `# this is a title`. 

### Publishing the RSS file
You can submit your RSS file to a public website, or put it somewhere easily available on a public drive, apache server, etc. Note, there is a chance I will make a function for this on my website so that people who can't host RSS feeds themselves can still publish their feeds. Please let me know if you're interested in such a service.

### Why RSS?
 * No ads/trackers/cookies.
 * No distractions when reading (it's just text).
 * Many great RSS readers available (newsboat!) that can create custom feeds.
 * Doesn't require a strong internet connection.
 * Amazing scene of personal blogs!

### Support/requesting features
You can always request fixes and features by opening an issue. If you want to support my work you can also sponsor this project through the button on the top right. Thank you!
