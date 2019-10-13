# remote pbcopy for iTerm2

`pbcopy` is a well-known macOS tool that copies data to the clipboard.
It's very useful, but available only in your local machine, not in remote machines.

Fortunately, with OSC52 escape sequence,
we can access the local machine clipboard via a remote machine.

I prepared a simple tool that is `pbcopy` for remote machines.

## Install

1. First, make sure you use iTerm2 version 3.0.0 or later.
2. Copy a preferred `pbcopy` to a directory where `$PATH` is set.

       [local]  $ ssh remote

       # If you prefer a self-contained binary, then
       [remote] $ wget -O pbcopy-linux-amd64.tar.gz https://github.com/skaji/remote-pbcopy-iterm2/releases/latest/download/pbcopy-linux-amd64.tar.gz
       [remote] $ tar xf pbcopy-linux-amd64.tar.gz
       [remote] $ mv pbcopy /path/to/bin/

       # If you prefer a perl script, then
       [remote] $ wget https://raw.githubusercontent.com/skaji/remote-pbcopy-iterm2/master/pbcopy
       [remote] $ chmod +x pbcopy
       [remote] $ mv pbcopy /path/to/bin/

3. Check "Applications in terminal may access clipboard" in iTerm2 Preferences:

    ![preferences.png](https://raw.githubusercontent.com/skaji/remote-pbcopy-iterm2/master/misc/preferences.png)

## Usage

Just like the normal `pbcopy`:

    [local]  $ ssh remote
    [remote] $ date | pbcopy
    [remote] $ exit
    [local]  $ pbpaste
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
* https://chromium.googlesource.com/apps/libapps/+/HEAD/hterm/etc/osc52.vim
* https://chromium.googlesource.com/apps/libapps/+/HEAD/hterm/etc/osc52.el
* https://chromium.googlesource.com/apps/libapps/+/HEAD/hterm/etc/osc52.sh

## Author

Shoichi Kaji

## License

MIT
