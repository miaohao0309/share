#!/usr/bin/python

from sys import argv
from docker import Client
import re

def get_image(img_hash,imgs):
    for i in imgs:
        if img_hash in i['Id']:
            img = i
            return img
    raise ImageNotFound("Image not found..\n")

def insert_step(step):
    if "#(nop)" in step:
        to_add = step.split("#(nop) ")[1]
    else:
        to_add = ("RUN{}".format(step))
        to_add = to_add.replace("&&", "\\\n    &&")
        reg = re.findall(r'RUN(.*-c)',to_add)
        to_add = to_add.replace(''.join(reg) , "")
    return to_add.strip(' ')

def parse_history(hist):
    cmds = []
    rec = False
    first_tag = False
    actual_tag = False
    for i in hist:
        if i['Tags']:
            actual_tag = i['Tags'][0]
            if first_tag and not rec:
                break
            first_tag = True
        cmd = insert_step(i['CreatedBy'])
        cmds.append(cmd)
    if not rec:
        cmds.append("FROM {}".format(actual_tag))
    return cmds

def main():
    cmds = []
    cli = Client(base_url='unix://var/run/docker.sock')
    img = get_image(argv[-1],cli.images())
    hist = cli.history(img['RepoTags'][0])
    cmds = parse_history(hist)
    for i in reversed(cmds):
        print(i)

if __name__ == '__main__':
    main() 

