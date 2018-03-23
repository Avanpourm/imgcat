#!/usr/bin/python

import os
import base64
import codecs
import sys

TERM = os.environ.get('TERM')
TERM_IS_SCREEN = TERM and TERM.startswith('screen')

_IMG_PRE = '\033Ptmux;\033\033]' if TERM_IS_SCREEN else '\033]'
_IMG_POST = '\a\033\\' if TERM_IS_SCREEN else '\a'


def _read_as_base64(path):
    with codecs.open(path, mode='rb') as fh:
        return base64.b64encode(fh.read())


def imgcat(path, inline=1, preserve_aspect_ratio=0, **kwargs):
    return '\n%s1337;File=inline=%d;preserveAspectRatio=%d:%s%s' % (
        _IMG_PRE, inline, preserve_aspect_ratio,
        _read_as_base64(path), _IMG_POST)

if __name__ == '__main__':
    if len(sys.argv)<2:
        exit()
    print imgcat(sys.argv[1])
