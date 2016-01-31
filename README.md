# remote pbcopy for iTerm2

`pbcopy` is a well-known MacOSX tool that copy data to the clipboard.
It's very useful, but available only in your local machine, not in remote machines.

Fortunately, with OSC52 escape sequence,
we can access the local machine clipboard via a remote machine.

I prepared a simple perl script that is "pbcopy" for remote machines.

## Install

First install the iTerm2 which
[removes the limit for OSC 52](https://github.com/gnachman/iTerm2/commit/9afa15e199b53d06ffaf18fc28b56d45ff722d19) from:

https://www.iterm2.com/downloads.html (maybe Test Releases or Nightly Builds)

Then just copy `pbcopy` to a directory where `$PATH` is set.

    [remote] $ wget https://raw.githubusercontent.com/shoichikaji/remote-pbcopy-iterm2/master/pbcopy
    [remote] $ chmod +x pbcopy
    [remote] $ mv pbcopy /path/to/bin/

And make sure to check "Allow clipboard access to terminal apps" in iTerm2 Preferences:

![preferences.png](https://raw.githubusercontent.com/shoichikaji/remote-pbcopy-iterm2/master/misc/preferences.png)

## Usage

Just like the normal `pbcopy`:

    [remote] $ data | pbcopy

    [local] $ pbpaste
    Sun Jan 18 20:28:03 JST 2015

## How about `pbpaste`?

Currently iTerm2 does not allow OSC 52 read access for security reasons.
But we can just use command+V key for paste content from clipboard.

If you want to save the content of clipboard to a remote file, try this:

    [remote] cat > out.txt
    # press command+V to paste content of clipboard,
    # and press contorol+D which indicats EOF

## See also

For OSC52

* http://doda.b.sourceforge.jp/2011/12/15/tmux-set-clipboard/
* http://qiita.com/kefir_/items/1f635fe66b778932e278
* http://qiita.com/kefir_/items/515ed5264fce40dec522

For iTerm2's "The issue with OSC 52 only reading the first 1024 chars has been fixed":

* https://code.google.com/p/iterm2/issues/detail?id=1428

## Author

Shoichi Kaji

# License

This software is licensed under the same terms as Perl.
