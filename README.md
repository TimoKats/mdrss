# 📝MDRSS

MDRSS is a markdown to RSS converter written in GO. With this tool, you can write articles in a local folder and have them automatically formatted to an RSS compliant XML file. As a result, you don't have to write articles on your website first and have them be read by an RSS reader. Moreover, MDRSS automatically takes care of publication dates, categories (next update), formatting, etc.

### Getting started
After downloading and compiling the source code, you can add your configuration in `~/.mdrss/config`. This is a JSON file with a number of settings. You can use this template to setup your configuration

```JSON
  {
    "Author": "Timo Kats",
    "Description": "Timo's weblog",
    "InputFolder": "/path/to/articles/", # This will contain your markdown files
    "OutputFile": "/path/to/webserver/index.xml",
    "Link": "index.xml" # can be anything, might be deprecated
  }
```

In your input folder, you can add an article by creating a `.md` file. **Note, if your filename is prefixed with `draft-` it will not be included in the RSS file**. Finally, you can type `mdrss ls` to view the articles ready for publishing and `mdrss update` to create the RSS feed.

### Publishing the RSS file
You can submit your RSS file to a public website, or put it somewhere easily available on a public drive, apache server, etc. Note, there is a chance I will make a function for this on my website so that people who can't host RSS feeds themselves can still publish their feeds. Please let me know if you're interested in such a service.

### Why RSS?
 * No ads/trackers.
 * No doomscrolling algorithm/dopamine cycles.
 * No distractions when reading (it's just text).
 * Many great RSS readers available (newsboat!) that can create custom feeds.
 * Doesn't require a strong internet connection.
 * Amazing scene of personal blogs!
