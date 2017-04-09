# remote pbcopy for iTerm2

`pbcopy` is a well-known MacOSX tool that copy data to the clipboard.
It's very useful, but available only in your local machine, not in remote machines.

Fortunately, with OSC52 escape sequence,
we can access the local machine clipboard via a remote machine.

I prepared a simple perl script that is `pbcopy` for remote machines.

## Install

1. First, make sure you use iTerm2 version 3.0.0 or later.
2. Copy `pbcopy` to a directory where `$PATH` is set.

       [remote] $ wget https://raw.githubusercontent.com/skaji/remote-pbcopy-iterm2/master/pbcopy
       [remote] $ chmod +x pbcopy
       [remote] $ mv pbcopy /path/to/bin/

3. Check "Allow clipboard access to terminal apps" in iTerm2 Preferences:

    ![preferences.png](https://raw.githubusercontent.com/skaji/remote-pbcopy-iterm2/master/misc/preferences.png)

## Usage

Just like the normal `pbcopy`:

    [remote] $ data | pbcopy

    [local] $ pbpaste
    Sun Jan 18 20:28:03 JST 2015

## How about `pbpaste`?

Currently iTerm2 does not allow OSC 52 read access for security reasons.
But we can just use command+V key to paste content from clipboard.

If you want to save the content of clipboard to a remote file, try this:

    [remote] cat > out.txt
    # press command+V to paste content of clipboard,
    # and press control+D which indicats EOF

## See also

For OSC52

* http://doda.b.sourceforge.jp/2011/12/15/tmux-set-clipboard/
* http://qiita.com/kefir_/items/1f635fe66b778932e278
* http://qiita.com/kefir_/items/515ed5264fce40dec522

## Author

Shoichi Kaji

## License

This software is licensed under the same terms as Perl.
