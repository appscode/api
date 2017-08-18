#!/usr/bin/env python

# Needed for antipackage with python 2
from __future__ import absolute_import

import re
import sys


def to_camel_case(snake_str):
    if snake_str.startswith('"') and snake_str.endswith('"'):
        return snake_str
    components = snake_str.split('_')
    # We capitalize the first letter of each component except the first one
    # with the 'title' method and join them together.
    return components[0] + ''.join((x[0].upper() + x[1:]) for x in components[1:])


def convert(fname):
    with open(fname, 'r+') as f:
        content = f.readlines()
        out = []
        for line in content:
            words = re.split(r'(\s+)', line)
            for idx, w in enumerate(words):
                words[idx] = to_camel_case(w)
            out.append(''.join(words))

        f.seek(0)
        f.writelines(out)
        f.truncate()


if __name__ == "__main__":
    convert(sys.argv[1])
