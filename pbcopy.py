#! /usr/bin/env python3

import base64
import os
import subprocess
import sys
from typing import Callable


def normal_esc(b64: str) -> str:
    return "\x1B]52;;" + b64 + "\x1B\x5C"


def tmux_esc(b64: str) -> str:
    return "\x1BPtmux;\x1B\x1B]52;;" + b64 + "\x1B\x1B\x5C\x5C\x1B\x5C"


def screen_esc(b64: str) -> str:
    out = []
    for i in range(sys.maxsize):
        begin, end = i * 76, min((i + 1) * 76, len(b64))
        out.append(("\x1BP\x1B]52;;" if begin == 0 else "\x1B\x5C\x1BP") + b64[begin:end])
        if end == len(b64):
            break
    out.append("\x07\x1B\x5C")

    return "".join(out)


def is_tmux_cc(pid: str) -> bool:
    try:
        out = subprocess.check_output(["ps", "-p", pid, "-o", "command="])
        out = out.rstrip(b"\n\r")
        for arg in out.split(b" "):
            if arg == b"-CC":
                return True
        return False
    except subprocess.CalledProcessError:
        return False


def choose_esc() -> Callable[[str], str]:
    env = os.getenv("TMUX")
    if env:
        envs = env.split(",")
        if len(envs) > 1 and is_tmux_cc(envs[1]):
            return normal_esc
        else:
            return tmux_esc

    env = os.getenv("TERM")
    if env and env.startswith("screen"):
        return screen_esc

    return normal_esc


def run():
    if len(sys.argv) == 1:
        b = sys.stdin.buffer.read()
    else:
        if sys.argv[1] == "-h" or sys.argv[1] == "--help":
            print("Usage:\n  pbcopy FILE\n  some-command | pbcopy\n")
            sys.exit(1)
        b = open(sys.argv[1], "rb").read()

    b = b.rstrip(b"\n\r")
    if b:
        b64 = base64.b64encode(b).decode(encoding="UTF-8")
        print(choose_esc()(b64))


if __name__ == "__main__":
    run()
