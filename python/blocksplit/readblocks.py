#!/sbin/env python3

import time


def readblock(filename):
    block = []
    with open(filename, "r") as f:
        for line in f:
            line = line.strip()
            if line == '':
                if len(block) > 0:
                    yield block
                block = []
            else:
                block.append(line)
    yield block


def timer(func):
    def wrapper(*args, **kwargs):
        start = time.time()
        value = func(*args, **kwargs)
        print(f'Took: {time.time() - start }')
        return value
    return wrapper


@timer
def readsplit(filename):
    with open(filename, "r") as f:
        prg = [x.strip().split(" ") for x in f]
    print(prg)


for b in readblock('blocks.txt'):
    print(f'--- {b} ---')

readsplit('splits.txt')
