#!/bin/sh
grep 'packagefile' $WORK/*/importcfg.link | sed -e 's/packagefile //' | tr '=' ' ' |
while read module library; do printf "$module\t"; du -ks "$library"; done |
sort -k2nr | cut -f1-2 | awk '{print $2"\t"$1; }'
