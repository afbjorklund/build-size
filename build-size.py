#!/usr/bin/python

import glob
import os

work = os.environ["WORK"]
link = glob.glob(work + "/*/importcfg.link")[0]

sizes = {}
for line in open(link, 'r'):
    line = line.replace("packagefile ", "")
    module, library = line.strip().split('=')
    sizes[module] = os.path.getsize(library)

for item in sorted(sizes, key=sizes.get, reverse=True):
    print "%d\t%s" % (sizes[item] / 1024, item)
