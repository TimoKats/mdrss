# 📝MDRSS

MDRSS is a markdown to RSS converter written in GO. With this tool, you can write articles in a local folder and have them automatically formatted to an RSS compliant XML file. Moreover, MDRSS automatically takes care of publication dates, categories (next update), formatting, etc.

### Configuration
You can install the binary using `go install github.com/TimoKats/mdrss@latest` (make sure your GOPATH is set correctly!). After this you can add your feeduration in `~/.mdrss/feed.json`. This is a file with a number of settings. You can use this template.

```ini
Author=Timo
Description=Timo's weblog
InputFolder=/path/to/articles/", # This will contain your markdown files
OutputFile=/path/to/webserver/index.xml",
Link=index.xml # can be anything, might be deprecated
```

Note, you can also specify your own feed path like so: `mdrss --config other/path/to/config update`

### Publication
In your input folder, you can add an article by creating a `.md` file. **Note, if your filename is prefixed with `draft-` it will not be included in the RSS file**. Finally, you can type `mdrss ls` to view the articles ready for publishing and `mdrss update` to create the RSS feed. Note, the title of the RSS articles are based on markdown headers. E.g. format the first line of the markdown articles as so: `# this is a title`.

