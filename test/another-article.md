# This article has a title.


# Added public PGP key email to website

Hi all. I've added my public PGP key to the files section of this website. It's linked to my main e-mail address (hello[at]timokats[dot]xyz).  

[audio/mpeg](https://www.native-instruments.com/fileadmin/ni_media/producer/koresoundpack/fm8transientattacks/audio/1_FM8ROXXS.mp3)

Note, for me it's not a hard requirement to use it when sending me an email. In fact, it's mainly there for those that prefer/require encrypted correspondence when contacting me. 

Timo

And here goes some text...

**bold** text, __however you like it__.
*italic* too? _Of course_!
Sometimes ***strong emphasis*** is needed to get across a ___point___.
~~You can always striketrough a bad idea~~,
Make something unique with a ~subscript~,
Or power up with a ^superscript^!

A [link](https://timokats.xyz)

And this is a list:
- hello
  - some indendation
- hello again
* other things

And this is another (ordered) list:
1. howdy
2. 1 + 1 = ...
4. out of order?
5. doesn't matter!

Codeblocks:
```go
// The regex that captures this codeblock:
fencedCodeBlock := regexp.MustCompile("^```")

	// Code blocks keep original formatting, whitespace is important!
	  // Show off your tabwidth in style!
```

Can I write text in between the code blocks?

```C++
int main() {
    int n, t1 = 0, t2 = 1, nextTerm = 0;

    cout << "Enter the number of terms: ";
    cin >> n;

    cout << "Fibonacci Series: ";

    for (int i = 1; i <= n; ++i) {
        // Prints the first two terms.
        if(i == 1) {
            cout << t1 << ", ";
            continue;
        }
        if(i == 2) {
            cout << t2 << ", ";
            continue;
        }
        nextTerm = t1 + t2;
        t1 = t2;
        t2 = nextTerm;
        
        cout << nextTerm << ", ";
    }
    return 0;
```


This feature also works `inline` as well!

And back to text again
