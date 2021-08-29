#!/usr/sbin/env python3

import readline

sysapiendpoints = [
    "GET/v3/brands/:brand_name/clusters/:cluster_name",
    "GET/v3/brands/:brand_name/clusters/:cluster_name/suspensions",
    "PATCH/v3/brands/:brand_name/clusters/:cluster_name"
]

brands = [
    "mybrand.com",
    "otherbrand.com"
]

tree = {}

def generate(key, l):
    """ Generate tree """
    keyword = l[0]

    if keyword not in key["list"]:
        key["list"].append(keyword)

    if keyword not in key["keys"]:
        key["keys"][keyword] = {}
        key["keys"][keyword]["list"] = []
        key["keys"][keyword]["keys"] = {}

    if len(l) > 1:
        generate(key["keys"][keyword], l[1:])

def prompt(key, path):
    readline.parse_and_bind("tab: complete")

    def complete(text,state):
        volcab = key["list"]
        results = [x for x in volcab if x.startswith(text)]
        return results[state]

    readline.set_completer(complete)

    line = input(f'{path}> ')
    if line in key["keys"]:

        app = line
        if line.startswith(":"):
            while True:
                var = input(f"set variable {line}? ")
                if var:
                    app = var
                    break

        path = path + "/" + app

        prompt(key["keys"][line], path)



if __name__ == "__main__":

    tree["start"] = {}
    tree["start"]["list"] = []
    tree["start"]["keys"] = {}

    paths = []
    for path in sysapiendpoints:
        if ":brand_name" in path:
            for brand in brands:
                tmp = path.replace(":brand_name", brand)
                paths.append(tmp)
        else:
            paths.append(path)

    for path in paths:
        tmp = path.split("/")
        generate(tree["start"], tmp)

    prompt(tree["start"], "")
