# 📝MDRSS
[![Go Report](https://goreportcard.com/badge/github.com/TimoKats/mdrss)](https://goreportcard.com/badge/github.com/TimoKats/mdrss)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![GitHub tag](https://img.shields.io/github/tag/TimoKats/mdrss?include_prereleases=&sort=semver&color=blue)](https://github.com/TimoKats/mdrss/releases/)


MDRSS is a markdown to RSS converter written in GO. With this tool, you can write articles in a local folder and have them automatically formatted to an RSS compliamdrss XML file. Moreover, MDRSS automatically takes care of publication dates, categories, formatting, etc.

## Getting started
You can install the binary using `go install github.com/TimoKats/mdrss@latest`. After this you can add your configuration in `~/.mdrss` or you can specify your own config path like so: `mdrss --config other/path/to/config update`.

```ini
Author=Timo
Description=Timo weblog
InputFolder=/path/to/articles/ ; This will comdrssain your markdown files
OutputFile=/path/to/webserver/index.xml
Link=index.xml ; can be anything, might be deprecated
```

## Publication
In your input folder, you can add an article by creating a .md file. Note, if your filename is prefixed with `draft-` it will not be included in the RSS file.  

Finally, you can type `mdrss ls` to view the articles ready for publishing and `mdrss update` to create the RSS feed. Note, the title of the RSS articles are based on markdown headers. E.g. format the first line of the markdown articles as so: `# this is a title`.


## Docs

Note, topics are added based on subfolders in your input folder. For example, if your input folder is `articles/`, then `articles/tech` will contain the markdown files for tech related articles.

<table>
  <thead>
    <tr>
      <th width="500px">Field</th>
      <th width="500px">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr width="600px">
      <td>Author</td>
      <td>Name of the RSS feed author.</td>
    </tr>
    <tr width="600px">
      <td>Description</td>
      <td>Description of the RSS feed.</td>
    </tr>
    <tr width="600px">
      <td>InputFolder</td>
      <td>Folder containing the markdown files. Add subfolders to enable the usage of topics.</td>
    </tr>
    <tr width="600px">
      <td>OutputFile</td>
      <td>Set this to write all articles to *one* XML file.</td>
    </tr>
    <tr width="600px">
      <td>OutputFolder</td>
      <td>Set this to write RSS topics to seperate XML files in one folder. File names will be based on topics. </td>
    </tr>
    <tr width="600px">
      <td>Link</td>
      <td>A bit redundant, but you can set this to the page hosting the RSS feed. </td>
    </tr>
  </tbody>
</table>
